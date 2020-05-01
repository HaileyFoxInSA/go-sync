// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model_type_store_schema_descriptor.proto

package sync_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ModelTypeStoreSchemaDescriptor struct {
	VersionNumber        *int64   `protobuf:"varint,1,opt,name=version_number,json=versionNumber" json:"version_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelTypeStoreSchemaDescriptor) Reset()         { *m = ModelTypeStoreSchemaDescriptor{} }
func (m *ModelTypeStoreSchemaDescriptor) String() string { return proto.CompactTextString(m) }
func (*ModelTypeStoreSchemaDescriptor) ProtoMessage()    {}
func (*ModelTypeStoreSchemaDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7900f39ccf06c2c, []int{0}
}

func (m *ModelTypeStoreSchemaDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelTypeStoreSchemaDescriptor.Unmarshal(m, b)
}
func (m *ModelTypeStoreSchemaDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelTypeStoreSchemaDescriptor.Marshal(b, m, deterministic)
}
func (m *ModelTypeStoreSchemaDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelTypeStoreSchemaDescriptor.Merge(m, src)
}
func (m *ModelTypeStoreSchemaDescriptor) XXX_Size() int {
	return xxx_messageInfo_ModelTypeStoreSchemaDescriptor.Size(m)
}
func (m *ModelTypeStoreSchemaDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelTypeStoreSchemaDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ModelTypeStoreSchemaDescriptor proto.InternalMessageInfo

func (m *ModelTypeStoreSchemaDescriptor) GetVersionNumber() int64 {
	if m != nil && m.VersionNumber != nil {
		return *m.VersionNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*ModelTypeStoreSchemaDescriptor)(nil), "sync_pb.ModelTypeStoreSchemaDescriptor")
}

func init() {
	proto.RegisterFile("model_type_store_schema_descriptor.proto", fileDescriptor_c7900f39ccf06c2c)
}

var fileDescriptor_c7900f39ccf06c2c = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcd, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0x80, 0x61, 0xc4, 0x43, 0xb0, 0x50, 0x07, 0x4f, 0x9d, 0x22, 0x02, 0x41, 0x08, 0xf6, 0x21,
	0x22, 0xa8, 0x4b, 0x11, 0xd9, 0x7d, 0xc9, 0x75, 0x48, 0xc1, 0xd9, 0x19, 0x66, 0xd6, 0xc0, 0xb7,
	0x8f, 0xb5, 0xe8, 0xfa, 0xc3, 0xc7, 0x6f, 0x2a, 0xa4, 0x16, 0x06, 0x17, 0x27, 0x06, 0xa7, 0x91,
	0x04, 0x9c, 0xfa, 0x0e, 0xf0, 0xe9, 0x5a, 0x50, 0x2f, 0x3d, 0x47, 0x12, 0xcb, 0x42, 0x91, 0x8a,
	0x85, 0x4e, 0xc1, 0x3b, 0x6e, 0x76, 0x27, 0xb3, 0xb9, 0x24, 0xf4, 0x98, 0x18, 0xea, 0x44, 0xea,
	0x59, 0x1c, 0xff, 0xa0, 0x28, 0xcd, 0xea, 0x0d, 0xa2, 0x3d, 0x05, 0x17, 0x46, 0x6c, 0x40, 0xd6,
	0xd9, 0x36, 0xab, 0xf2, 0xfb, 0xf2, 0x57, 0xaf, 0x73, 0x3c, 0xec, 0x4d, 0x49, 0xf2, 0xb2, 0xbe,
	0x13, 0xc2, 0x7e, 0x44, 0xeb, 0x09, 0x99, 0x02, 0x84, 0xa8, 0x36, 0xbd, 0xbe, 0x5f, 0x4f, 0xc3,
	0x39, 0xbf, 0x65, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xc5, 0x17, 0xae, 0xa9, 0x00, 0x00,
	0x00,
}
