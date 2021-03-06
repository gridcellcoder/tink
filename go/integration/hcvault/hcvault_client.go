// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Package hcvault provides integration with the HashiCorp Vault (https://www.vaultproject.io/).
// Below there is an example of how the integration code can be used:

// package main
//
// import (
// 	"fmt"
// 	"log"
//
//  "github.com/google/tink/go/aead"
//  "github.com/google/tink/go/core/registry"
//  "github.com/google/tink/go/integration/hcvault"
//  "github.com/google/tink/go/keyset"
// )
//
// const (
// 	keyURI = "hcvault://hcvault.corp.com:8200/transit/keys/key-1"
// )
//
// func main() {
//  tlsConf := getTLSConfig()
//  token := getVaultToken()
//  vaultClient, err := hcvault.NewHCVaultClient(keyURI, tlsConf, token)
// 	if err != nil {
//    // handle error
// 	}
// 	registry.RegisterKMSClient(vaultClient)
//
// 	dek := aead.AES128CTRHMACSHA256KeyTemplate()
// 	kh, err := keyset.NewHandle(aead.KMSEnvelopeAEADKeyTemplate(keyURI, dek))
// 	if err != nil {
//    // handle error
// 	}
// 	a, err := aead.New(kh)
// 	if err != nil {
//    // handle error
// 	}
//
// 	msg := "secret message"
// 	ct, err := a.Encrypt([]byte(msg), nil)
// 	if err != nil {
//    // handle error
// 	}
//
// 	pt, err := a.Decrypt(ct, nil)
// 	if err != nil {
//    // handle error
// 	}
// }

package hcvault

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/google/tink/go/core/registry"
	"github.com/google/tink/go/tink"
)

const (
	vaultPrefix = "hcvault://"
)

// vaultClient represents a client that connects to the HashiCorp Vault backend.
type vaultClient struct {
	keyURIPrefix string
	client       *api.Logical
}

var _ registry.KMSClient = (*vaultClient)(nil)

// NewHCVaultClient returns a new client to HashiCorp Vault.
// uriPrefix parameter is a valid URI which must have "hcvault" scheme and
// vault server address and port. Specific key URIs will be matched against this
// prefix to determine if the client supports the key or not.
// tlsCfg represents tls.Config which will be used to communicate with Vault
// server via HTTPS protocol. If not specified a default tls.Config{} will be
// used.
func NewHCVaultClient(uriPrefix string, tlsCfg *tls.Config, token string) (registry.KMSClient, error) {
	if !strings.HasPrefix(strings.ToLower(uriPrefix), vaultPrefix) {
		return nil, fmt.Errorf("key URI must start with %s", vaultPrefix)
	}

	httpClient := api.DefaultConfig().HttpClient
	transport := httpClient.Transport.(*http.Transport)
	if tlsCfg == nil {
		tlsCfg = &tls.Config{}
	} else {
		tlsCfg = tlsCfg.Clone()
	}
	transport.TLSClientConfig = tlsCfg

	vurl, err := url.Parse(uriPrefix)
	if err != nil {
		return nil, err
	}

	cfg := &api.Config{
		Address:    "https://" + vurl.Host,
		HttpClient: httpClient,
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	client.SetToken(token)
	return &vaultClient{
		keyURIPrefix: uriPrefix,
		client:       client.Logical(),
	}, nil

}

// Supported returns true if this client does support keyURI.
func (c *vaultClient) Supported(keyURI string) bool {
	return strings.HasPrefix(keyURI, c.keyURIPrefix)
}

// GetAEAD gets an AEAD backend by keyURI.
func (c *vaultClient) GetAEAD(keyURI string) (tink.AEAD, error) {
	if !strings.HasPrefix(keyURI, c.keyURIPrefix) {
		return nil, fmt.Errorf("this client is bound to %s prefix, cannot load keys bound to %s", c.keyURIPrefix, keyURI)
	}
	return NewHCVaultAEAD(keyURI, c.client), nil
}
