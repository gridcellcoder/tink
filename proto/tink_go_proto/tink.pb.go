// Copyright 2017 Google Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tink.proto

package tink_go_proto // import "github.com/google/tink/proto/tink_go_proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyStatusType int32

const (
	KeyStatusType_UNKNOWN_STATUS KeyStatusType = 0
	KeyStatusType_ENABLED        KeyStatusType = 1
	KeyStatusType_DISABLED       KeyStatusType = 2
	KeyStatusType_DESTROYED      KeyStatusType = 3
)

var KeyStatusType_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "ENABLED",
	2: "DISABLED",
	3: "DESTROYED",
}

var KeyStatusType_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"ENABLED":        1,
	"DISABLED":       2,
	"DESTROYED":      3,
}

func (x KeyStatusType) String() string {
	return proto.EnumName(KeyStatusType_name, int32(x))
}

func (KeyStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{0}
}

// Tink produces and accepts ciphertexts or signatures that consist
// of a prefix and a payload. The payload and its format is determined
// entirely by the primitive, but the prefix has to be one of the following
// 4 types:
//   - Legacy: prefix is 5 bytes, starts with \x00 and followed by a 4-byte
//             key id that is computed from the key material.
//   - Crunchy: prefix is 5 bytes, starts with \x00 and followed by a 4-byte
//             key id that is generated randomly.
//   - Tink  : prefix is 5 bytes, starts with \x01 and followed by 4-byte
//             key id that is generated randomly.
//   - Raw   : prefix is 0 byte, i.e., empty.
type OutputPrefixType int32

const (
	OutputPrefixType_UNKNOWN_PREFIX OutputPrefixType = 0
	OutputPrefixType_TINK           OutputPrefixType = 1
	OutputPrefixType_LEGACY         OutputPrefixType = 2
	OutputPrefixType_RAW            OutputPrefixType = 3
	// CRUNCHY is like LEGACY, but with two differences:
	//   - Its key id is generated randomly (like TINK)
	//   - Its signature schemes don't append zero to sign messages
	OutputPrefixType_CRUNCHY OutputPrefixType = 4
)

var OutputPrefixType_name = map[int32]string{
	0: "UNKNOWN_PREFIX",
	1: "TINK",
	2: "LEGACY",
	3: "RAW",
	4: "CRUNCHY",
}

var OutputPrefixType_value = map[string]int32{
	"UNKNOWN_PREFIX": 0,
	"TINK":           1,
	"LEGACY":         2,
	"RAW":            3,
	"CRUNCHY":        4,
}

func (x OutputPrefixType) String() string {
	return proto.EnumName(OutputPrefixType_name, int32(x))
}

func (OutputPrefixType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{1}
}

type KeyData_KeyMaterialType int32

const (
	KeyData_UNKNOWN_KEYMATERIAL KeyData_KeyMaterialType = 0
	KeyData_SYMMETRIC           KeyData_KeyMaterialType = 1
	KeyData_ASYMMETRIC_PRIVATE  KeyData_KeyMaterialType = 2
	KeyData_ASYMMETRIC_PUBLIC   KeyData_KeyMaterialType = 3
	KeyData_REMOTE              KeyData_KeyMaterialType = 4
)

var KeyData_KeyMaterialType_name = map[int32]string{
	0: "UNKNOWN_KEYMATERIAL",
	1: "SYMMETRIC",
	2: "ASYMMETRIC_PRIVATE",
	3: "ASYMMETRIC_PUBLIC",
	4: "REMOTE",
}

var KeyData_KeyMaterialType_value = map[string]int32{
	"UNKNOWN_KEYMATERIAL": 0,
	"SYMMETRIC":           1,
	"ASYMMETRIC_PRIVATE":  2,
	"ASYMMETRIC_PUBLIC":   3,
	"REMOTE":              4,
}

func (x KeyData_KeyMaterialType) String() string {
	return proto.EnumName(KeyData_KeyMaterialType_name, int32(x))
}

func (KeyData_KeyMaterialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{1, 0}
}

type KeyTemplate struct {
	// Required.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Optional.
	// If missing, it means the key type doesn't require a *KeyFormat proto.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Optional.
	// If missing, uses OutputPrefixType.TINK.
	OutputPrefixType     OutputPrefixType `protobuf:"varint,3,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KeyTemplate) Reset()         { *m = KeyTemplate{} }
func (m *KeyTemplate) String() string { return proto.CompactTextString(m) }
func (*KeyTemplate) ProtoMessage()    {}
func (*KeyTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{0}
}
func (m *KeyTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyTemplate.Unmarshal(m, b)
}
func (m *KeyTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyTemplate.Marshal(b, m, deterministic)
}
func (dst *KeyTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyTemplate.Merge(dst, src)
}
func (m *KeyTemplate) XXX_Size() int {
	return xxx_messageInfo_KeyTemplate.Size(m)
}
func (m *KeyTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_KeyTemplate proto.InternalMessageInfo

func (m *KeyTemplate) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *KeyTemplate) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyTemplate) GetOutputPrefixType() OutputPrefixType {
	if m != nil {
		return m.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

// The actual *Key-proto is wrapped in a KeyData message, which in addition
// to this serialized proto contains also type_url identifying the
// definition of *Key-proto (as in KeyFormat-message), and some extra metadata
// about the type key material.
type KeyData struct {
	// Required.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Required.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Required.
	KeyMaterialType      KeyData_KeyMaterialType `protobuf:"varint,3,opt,name=key_material_type,json=keyMaterialType,proto3,enum=google.crypto.tink.KeyData_KeyMaterialType" json:"key_material_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeyData) Reset()         { *m = KeyData{} }
func (m *KeyData) String() string { return proto.CompactTextString(m) }
func (*KeyData) ProtoMessage()    {}
func (*KeyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{1}
}
func (m *KeyData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyData.Unmarshal(m, b)
}
func (m *KeyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyData.Marshal(b, m, deterministic)
}
func (dst *KeyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyData.Merge(dst, src)
}
func (m *KeyData) XXX_Size() int {
	return xxx_messageInfo_KeyData.Size(m)
}
func (m *KeyData) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyData.DiscardUnknown(m)
}

var xxx_messageInfo_KeyData proto.InternalMessageInfo

func (m *KeyData) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *KeyData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyData) GetKeyMaterialType() KeyData_KeyMaterialType {
	if m != nil {
		return m.KeyMaterialType
	}
	return KeyData_UNKNOWN_KEYMATERIAL
}

// A Tink user works usually not with single keys, but with keysets,
// to enable key rotation.  The keys in a keyset can belong to different
// implementations/key types, but must all implement the same primitive.
// Any given keyset (and any given key) can be used for one primitive only.
type Keyset struct {
	// Identifies key used to generate new crypto data (encrypt, sign).
	// Required.
	PrimaryKeyId uint32 `protobuf:"varint,1,opt,name=primary_key_id,json=primaryKeyId,proto3" json:"primary_key_id,omitempty"`
	// Actual keys in the Keyset.
	// Required.
	Key                  []*Keyset_Key `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Keyset) Reset()         { *m = Keyset{} }
func (m *Keyset) String() string { return proto.CompactTextString(m) }
func (*Keyset) ProtoMessage()    {}
func (*Keyset) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{2}
}
func (m *Keyset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyset.Unmarshal(m, b)
}
func (m *Keyset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyset.Marshal(b, m, deterministic)
}
func (dst *Keyset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyset.Merge(dst, src)
}
func (m *Keyset) XXX_Size() int {
	return xxx_messageInfo_Keyset.Size(m)
}
func (m *Keyset) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyset.DiscardUnknown(m)
}

var xxx_messageInfo_Keyset proto.InternalMessageInfo

func (m *Keyset) GetPrimaryKeyId() uint32 {
	if m != nil {
		return m.PrimaryKeyId
	}
	return 0
}

func (m *Keyset) GetKey() []*Keyset_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

type Keyset_Key struct {
	// Contains the actual, instantiation specific key proto.
	// By convention, each key proto contains a version field.
	KeyData *KeyData      `protobuf:"bytes,1,opt,name=key_data,json=keyData,proto3" json:"key_data,omitempty"`
	Status  KeyStatusType `protobuf:"varint,2,opt,name=status,proto3,enum=google.crypto.tink.KeyStatusType" json:"status,omitempty"`
	// Identifies a key within a keyset, is a part of metadata
	// of a ciphertext/signature.
	KeyId uint32 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Determines the prefix of the ciphertexts/signatures produced by this key.
	// This value is copied verbatim from the key template.
	OutputPrefixType     OutputPrefixType `protobuf:"varint,4,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Keyset_Key) Reset()         { *m = Keyset_Key{} }
func (m *Keyset_Key) String() string { return proto.CompactTextString(m) }
func (*Keyset_Key) ProtoMessage()    {}
func (*Keyset_Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{2, 0}
}
func (m *Keyset_Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyset_Key.Unmarshal(m, b)
}
func (m *Keyset_Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyset_Key.Marshal(b, m, deterministic)
}
func (dst *Keyset_Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyset_Key.Merge(dst, src)
}
func (m *Keyset_Key) XXX_Size() int {
	return xxx_messageInfo_Keyset_Key.Size(m)
}
func (m *Keyset_Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyset_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Keyset_Key proto.InternalMessageInfo

func (m *Keyset_Key) GetKeyData() *KeyData {
	if m != nil {
		return m.KeyData
	}
	return nil
}

func (m *Keyset_Key) GetStatus() KeyStatusType {
	if m != nil {
		return m.Status
	}
	return KeyStatusType_UNKNOWN_STATUS
}

func (m *Keyset_Key) GetKeyId() uint32 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *Keyset_Key) GetOutputPrefixType() OutputPrefixType {
	if m != nil {
		return m.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

// Represents a "safe" Keyset that doesn't contain any actual key material,
// thus can be used for logging or monitoring. Most fields are copied from
// Keyset.
type KeysetInfo struct {
	// See Keyset.primary_key_id.
	PrimaryKeyId uint32 `protobuf:"varint,1,opt,name=primary_key_id,json=primaryKeyId,proto3" json:"primary_key_id,omitempty"`
	// KeyInfos in the KeysetInfo.
	// Each KeyInfo is corresponding to a Key in the corresponding Keyset.
	KeyInfo              []*KeysetInfo_KeyInfo `protobuf:"bytes,2,rep,name=key_info,json=keyInfo,proto3" json:"key_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeysetInfo) Reset()         { *m = KeysetInfo{} }
func (m *KeysetInfo) String() string { return proto.CompactTextString(m) }
func (*KeysetInfo) ProtoMessage()    {}
func (*KeysetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{3}
}
func (m *KeysetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeysetInfo.Unmarshal(m, b)
}
func (m *KeysetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeysetInfo.Marshal(b, m, deterministic)
}
func (dst *KeysetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeysetInfo.Merge(dst, src)
}
func (m *KeysetInfo) XXX_Size() int {
	return xxx_messageInfo_KeysetInfo.Size(m)
}
func (m *KeysetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeysetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeysetInfo proto.InternalMessageInfo

func (m *KeysetInfo) GetPrimaryKeyId() uint32 {
	if m != nil {
		return m.PrimaryKeyId
	}
	return 0
}

func (m *KeysetInfo) GetKeyInfo() []*KeysetInfo_KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

type KeysetInfo_KeyInfo struct {
	// the type url of this key,
	// e.g., type.googleapis.com/google.crypto.tink.HmacKey.
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// See Keyset.Key.status.
	Status KeyStatusType `protobuf:"varint,2,opt,name=status,proto3,enum=google.crypto.tink.KeyStatusType" json:"status,omitempty"`
	// See Keyset.Key.key_id.
	KeyId uint32 `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// See Keyset.Key.output_prefix_type.
	OutputPrefixType     OutputPrefixType `protobuf:"varint,4,opt,name=output_prefix_type,json=outputPrefixType,proto3,enum=google.crypto.tink.OutputPrefixType" json:"output_prefix_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KeysetInfo_KeyInfo) Reset()         { *m = KeysetInfo_KeyInfo{} }
func (m *KeysetInfo_KeyInfo) String() string { return proto.CompactTextString(m) }
func (*KeysetInfo_KeyInfo) ProtoMessage()    {}
func (*KeysetInfo_KeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{3, 0}
}
func (m *KeysetInfo_KeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeysetInfo_KeyInfo.Unmarshal(m, b)
}
func (m *KeysetInfo_KeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeysetInfo_KeyInfo.Marshal(b, m, deterministic)
}
func (dst *KeysetInfo_KeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeysetInfo_KeyInfo.Merge(dst, src)
}
func (m *KeysetInfo_KeyInfo) XXX_Size() int {
	return xxx_messageInfo_KeysetInfo_KeyInfo.Size(m)
}
func (m *KeysetInfo_KeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeysetInfo_KeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeysetInfo_KeyInfo proto.InternalMessageInfo

func (m *KeysetInfo_KeyInfo) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *KeysetInfo_KeyInfo) GetStatus() KeyStatusType {
	if m != nil {
		return m.Status
	}
	return KeyStatusType_UNKNOWN_STATUS
}

func (m *KeysetInfo_KeyInfo) GetKeyId() uint32 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *KeysetInfo_KeyInfo) GetOutputPrefixType() OutputPrefixType {
	if m != nil {
		return m.OutputPrefixType
	}
	return OutputPrefixType_UNKNOWN_PREFIX
}

// Represents a keyset that is encrypted with a master key.
type EncryptedKeyset struct {
	// Required.
	EncryptedKeyset []byte `protobuf:"bytes,2,opt,name=encrypted_keyset,json=encryptedKeyset,proto3" json:"encrypted_keyset,omitempty"`
	// Optional.
	KeysetInfo           *KeysetInfo `protobuf:"bytes,3,opt,name=keyset_info,json=keysetInfo,proto3" json:"keyset_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EncryptedKeyset) Reset()         { *m = EncryptedKeyset{} }
func (m *EncryptedKeyset) String() string { return proto.CompactTextString(m) }
func (*EncryptedKeyset) ProtoMessage()    {}
func (*EncryptedKeyset) Descriptor() ([]byte, []int) {
	return fileDescriptor_25afefcb68a59a56, []int{4}
}
func (m *EncryptedKeyset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedKeyset.Unmarshal(m, b)
}
func (m *EncryptedKeyset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedKeyset.Marshal(b, m, deterministic)
}
func (dst *EncryptedKeyset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedKeyset.Merge(dst, src)
}
func (m *EncryptedKeyset) XXX_Size() int {
	return xxx_messageInfo_EncryptedKeyset.Size(m)
}
func (m *EncryptedKeyset) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedKeyset.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedKeyset proto.InternalMessageInfo

func (m *EncryptedKeyset) GetEncryptedKeyset() []byte {
	if m != nil {
		return m.EncryptedKeyset
	}
	return nil
}

func (m *EncryptedKeyset) GetKeysetInfo() *KeysetInfo {
	if m != nil {
		return m.KeysetInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyTemplate)(nil), "google.crypto.tink.KeyTemplate")
	proto.RegisterType((*KeyData)(nil), "google.crypto.tink.KeyData")
	proto.RegisterType((*Keyset)(nil), "google.crypto.tink.Keyset")
	proto.RegisterType((*Keyset_Key)(nil), "google.crypto.tink.Keyset.Key")
	proto.RegisterType((*KeysetInfo)(nil), "google.crypto.tink.KeysetInfo")
	proto.RegisterType((*KeysetInfo_KeyInfo)(nil), "google.crypto.tink.KeysetInfo.KeyInfo")
	proto.RegisterType((*EncryptedKeyset)(nil), "google.crypto.tink.EncryptedKeyset")
	proto.RegisterEnum("google.crypto.tink.KeyStatusType", KeyStatusType_name, KeyStatusType_value)
	proto.RegisterEnum("google.crypto.tink.OutputPrefixType", OutputPrefixType_name, OutputPrefixType_value)
	proto.RegisterEnum("google.crypto.tink.KeyData_KeyMaterialType", KeyData_KeyMaterialType_name, KeyData_KeyMaterialType_value)
}

func init() { proto.RegisterFile("proto/tink.proto", fileDescriptor_25afefcb68a59a56) }

var fileDescriptor_25afefcb68a59a56 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0xed, 0x34, 0x49, 0x27, 0x6d, 0xb2, 0x5d, 0x28, 0x84, 0x82, 0xaa, 0x10, 0x55, 0x28,
	0x14, 0x29, 0x45, 0x41, 0x42, 0xe2, 0x84, 0x9c, 0x64, 0x01, 0xcb, 0xf9, 0xd3, 0xc6, 0xa1, 0x84,
	0x8b, 0xe5, 0x36, 0xdb, 0x60, 0x39, 0x89, 0x2d, 0x67, 0x83, 0xf0, 0x81, 0x07, 0xe0, 0xca, 0x23,
	0xf0, 0x28, 0x1c, 0x38, 0xf0, 0x1a, 0xbc, 0x08, 0xda, 0xb5, 0x5b, 0xda, 0xd0, 0x46, 0x20, 0x4e,
	0x9c, 0x76, 0x67, 0xfc, 0xcd, 0xcf, 0xf7, 0x79, 0x66, 0x01, 0x05, 0xa1, 0xcf, 0xfd, 0x43, 0xee,
	0xce, 0xbc, 0xaa, 0xbc, 0x62, 0x3c, 0xf6, 0xfd, 0xf1, 0x84, 0x55, 0x4f, 0xc2, 0x28, 0xe0, 0x7e,
	0x55, 0x7c, 0x29, 0x7f, 0x56, 0x20, 0x67, 0xb2, 0xc8, 0x62, 0xd3, 0x60, 0xe2, 0x70, 0x86, 0xef,
	0x40, 0x96, 0x47, 0x01, 0xb3, 0x17, 0xe1, 0xa4, 0xa8, 0x94, 0x94, 0xca, 0x06, 0xcd, 0x08, 0x7b,
	0x10, 0x4e, 0xf0, 0x4d, 0x58, 0x7f, 0xef, 0x4c, 0x16, 0xac, 0xa8, 0x96, 0x94, 0xca, 0x26, 0x8d,
	0x0d, 0x4c, 0x01, 0xfb, 0x0b, 0x1e, 0x2c, 0xb8, 0x1d, 0x84, 0xec, 0xd4, 0xfd, 0x60, 0x0b, 0x78,
	0x51, 0x2b, 0x29, 0x95, 0x7c, 0x6d, 0xbf, 0xfa, 0x7b, 0xc5, 0x6a, 0x57, 0xa2, 0x7b, 0x12, 0x6c,
	0x45, 0x01, 0xa3, 0xc8, 0x5f, 0xf2, 0x94, 0x3f, 0xa9, 0x90, 0x31, 0x59, 0xd4, 0x74, 0xb8, 0xf3,
	0xf7, 0x0d, 0x1d, 0xc1, 0xb6, 0xc7, 0x22, 0x7b, 0xea, 0x70, 0x16, 0xba, 0xce, 0xe4, 0x62, 0x3f,
	0x8f, 0xae, 0xea, 0x27, 0x29, 0x24, 0xce, 0x76, 0x12, 0x23, 0xdb, 0x2a, 0x78, 0x97, 0x1d, 0x65,
	0x0e, 0x85, 0x25, 0x0c, 0xbe, 0x0d, 0x37, 0x06, 0x1d, 0xb3, 0xd3, 0x3d, 0xea, 0xd8, 0x26, 0x19,
	0xb6, 0x75, 0x8b, 0x50, 0x43, 0x6f, 0xa1, 0x35, 0xbc, 0x05, 0x1b, 0xfd, 0x61, 0xbb, 0x4d, 0x2c,
	0x6a, 0x34, 0x90, 0x82, 0x6f, 0x01, 0xd6, 0xcf, 0x6d, 0xbb, 0x47, 0x8d, 0xd7, 0xba, 0x45, 0x90,
	0x8a, 0x77, 0x60, 0xfb, 0xa2, 0x7f, 0x50, 0x6f, 0x19, 0x0d, 0xa4, 0x61, 0x80, 0x34, 0x25, 0xed,
	0xae, 0x45, 0x50, 0xaa, 0xfc, 0x4d, 0x85, 0xb4, 0xc9, 0xa2, 0x39, 0xe3, 0x78, 0x1f, 0xf2, 0x41,
	0xe8, 0x4e, 0x9d, 0x30, 0xb2, 0x05, 0x43, 0x77, 0x24, 0x05, 0xd9, 0xa2, 0x9b, 0x89, 0xd7, 0x64,
	0x91, 0x31, 0xc2, 0x8f, 0x41, 0xf3, 0x58, 0x54, 0x54, 0x4b, 0x5a, 0x25, 0x57, 0xdb, 0xbb, 0x86,
	0xf1, 0x9c, 0x71, 0x71, 0x50, 0x01, 0xdd, 0xfd, 0xa1, 0x80, 0x66, 0xb2, 0x08, 0x3f, 0x85, 0xac,
	0xc8, 0x3b, 0x72, 0xb8, 0x23, 0x33, 0xe7, 0x6a, 0x77, 0x57, 0x08, 0x46, 0x33, 0x5e, 0xf2, 0x8b,
	0x9e, 0x41, 0x7a, 0xce, 0x1d, 0xbe, 0x98, 0xcb, 0x1f, 0x91, 0xaf, 0xdd, 0xbf, 0x26, 0xaa, 0x2f,
	0x41, 0x52, 0xdc, 0x24, 0x00, 0xef, 0x40, 0x3a, 0xa1, 0xa2, 0x49, 0x2a, 0xeb, 0x9e, 0xe4, 0x70,
	0xf5, 0x50, 0xa5, 0xfe, 0x69, 0xa8, 0xbe, 0xaa, 0x00, 0x31, 0x73, 0x63, 0x76, 0xea, 0xff, 0xa1,
	0x98, 0x7a, 0x2c, 0x89, 0x3b, 0x3b, 0xf5, 0x13, 0x45, 0x1f, 0x5c, 0xaf, 0xa8, 0xc8, 0x2b, 0xae,
	0xe2, 0x94, 0xea, 0x88, 0xcb, 0xee, 0x77, 0x45, 0x0e, 0xb3, 0x2c, 0xba, 0x62, 0x98, 0xff, 0x0f,
	0x11, 0x3f, 0x42, 0x81, 0xcc, 0x64, 0x0c, 0x1b, 0x25, 0x53, 0xf9, 0x10, 0x10, 0x3b, 0x73, 0x09,
	0x29, 0xe7, 0x8c, 0x27, 0x0b, 0x59, 0x60, 0x4b, 0xd0, 0xe7, 0x90, 0x8b, 0x01, 0xb1, 0xa0, 0x9a,
	0x9c, 0xb1, 0xbd, 0xd5, 0x82, 0x52, 0xf0, 0xce, 0xef, 0x07, 0x6d, 0xd8, 0xba, 0x24, 0x01, 0xc6,
	0x90, 0x3f, 0x5b, 0xc0, 0xbe, 0xa5, 0x5b, 0x83, 0x3e, 0x5a, 0xc3, 0x39, 0xc8, 0x90, 0x8e, 0x5e,
	0x6f, 0x91, 0x26, 0x52, 0xf0, 0x26, 0x64, 0x9b, 0x46, 0x3f, 0xb6, 0x54, 0xb1, 0x96, 0x4d, 0xd2,
	0xb7, 0x68, 0x77, 0x48, 0x9a, 0x48, 0x3b, 0xa0, 0x80, 0x96, 0x39, 0x5f, 0xcc, 0xd8, 0xa3, 0xe4,
	0x85, 0xf1, 0x06, 0xad, 0xe1, 0x2c, 0xa4, 0x2c, 0xa3, 0x63, 0x22, 0x45, 0x6c, 0x66, 0x8b, 0xbc,
	0xd4, 0x1b, 0x43, 0xa4, 0xe2, 0x0c, 0x68, 0x54, 0x3f, 0x42, 0x9a, 0x28, 0xd8, 0xa0, 0x83, 0x4e,
	0xe3, 0xd5, 0x10, 0xa5, 0xea, 0x03, 0xb8, 0x77, 0xe2, 0x4f, 0xaf, 0xe2, 0x24, 0x1f, 0xe1, 0x9e,
	0xf2, 0xf6, 0x60, 0xec, 0xf2, 0x77, 0x8b, 0xe3, 0xea, 0x89, 0x3f, 0x3d, 0x8c, 0x61, 0xf2, 0x91,
	0x3e, 0xfc, 0xf5, 0x5e, 0xdb, 0x63, 0xdf, 0x96, 0xd6, 0x17, 0x35, 0x2d, 0x0a, 0xf7, 0xea, 0xc7,
	0x69, 0x69, 0x3f, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xec, 0xe9, 0x2d, 0xe6, 0xd6, 0x05, 0x00,
	0x00,
}