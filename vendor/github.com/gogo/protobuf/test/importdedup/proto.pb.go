// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto.proto

package importdedup

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import subpkg "github.com/gogo/protobuf/test/importdedup/subpkg"

import github_com_gogo_protobuf_test_importdedup_subpkg "github.com/gogo/protobuf/test/importdedup/subpkg"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Object struct {
	CustomField          *github_com_gogo_protobuf_test_importdedup_subpkg.CustomType `protobuf:"bytes,1,opt,name=CustomField,customtype=github.com/gogo/protobuf/test/importdedup/subpkg.CustomType" json:"CustomField,omitempty"`
	SubObject            *subpkg.SubObject                                            `protobuf:"bytes,2,opt,name=SubObject" json:"SubObject,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                     `json:"-"`
	XXX_unrecognized     []byte                                                       `json:"-"`
	XXX_sizecache        int32                                                        `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto_38d4f6a4f3773b6e, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (dst *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(dst, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetSubObject() *subpkg.SubObject {
	if m != nil {
		return m.SubObject
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "importdedup.Object")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_proto_38d4f6a4f3773b6e) }

var fileDescriptor_proto_38d4f6a4f3773b6e = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xdc, 0x99, 0xb9, 0x05, 0xf9, 0x45, 0x25, 0x29, 0xa9, 0x29, 0xa5,
	0x05, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9,
	0xe9, 0xf9, 0xfa, 0x60, 0x35, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xf4, 0x4a,
	0xd9, 0xe3, 0x54, 0x5e, 0x92, 0x5a, 0x5c, 0xa2, 0x8f, 0x64, 0xb2, 0x7e, 0x71, 0x69, 0x52, 0x41,
	0x76, 0x3a, 0x98, 0x42, 0x58, 0xae, 0x34, 0x87, 0x91, 0x8b, 0xcd, 0x3f, 0x29, 0x2b, 0x35, 0xb9,
	0x44, 0x28, 0x91, 0x8b, 0xdb, 0xb9, 0xb4, 0xb8, 0x24, 0x3f, 0xd7, 0x2d, 0x33, 0x35, 0x27, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xc7, 0xc9, 0xfe, 0xd6, 0x3d, 0x79, 0x6b, 0x52, 0x2d, 0xd1, 0x83,
	0x98, 0x13, 0x52, 0x59, 0x90, 0x1a, 0x84, 0x6c, 0xa6, 0x90, 0x3e, 0x17, 0x67, 0x70, 0x69, 0x12,
	0xc4, 0x3e, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x41, 0x3d, 0xa8, 0x1e, 0xb8, 0x44, 0x10,
	0x42, 0x0d, 0x20, 0x00, 0x00, 0xff, 0xff, 0x21, 0x11, 0x7d, 0xc2, 0x29, 0x01, 0x00, 0x00,
}
