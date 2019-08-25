// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_helloworld_0196a2ef510324e3, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

var E_DefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "main.default_string",
	Tag:           "bytes,50000,opt,name=default_string,json=defaultString",
	Filename:      "helloworld.proto",
}

var E_DefaultInt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50001,
	Name:          "main.default_int",
	Tag:           "varint,50001,opt,name=default_int,json=defaultInt",
	Filename:      "helloworld.proto",
}

func init() {
	proto.RegisterType((*Message)(nil), "main.Message")
	proto.RegisterExtension(E_DefaultString)
	proto.RegisterExtension(E_DefaultInt)
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_helloworld_0196a2ef510324e3) }

var fileDescriptor_helloworld_0196a2ef510324e3 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d,
	0xcc, 0xcc, 0x93, 0x52, 0x48, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6,
	0xe9, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0x41, 0xd4, 0x29, 0x39, 0x72,
	0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0xc9, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x71, 0x35, 0x6d, 0x95, 0x60, 0x4b, 0xcf, 0x2f, 0xc8,
	0x48, 0x2d, 0x0a, 0x02, 0x8b, 0x0b, 0x89, 0x71, 0x31, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30,
	0x6a, 0xb0, 0x3a, 0xb1, 0x74, 0x6c, 0x95, 0xe0, 0x0a, 0x02, 0x09, 0x58, 0xb9, 0x71, 0xf1, 0xa5,
	0xa4, 0xa6, 0x25, 0x96, 0xe6, 0x94, 0xc4, 0x17, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x0b, 0xc9, 0xea,
	0x41, 0xec, 0xd5, 0x83, 0xd9, 0xab, 0xe7, 0x96, 0x99, 0x9a, 0x93, 0xe2, 0x5f, 0x50, 0x92, 0x99,
	0x9f, 0x57, 0x2c, 0x71, 0xa1, 0x8d, 0x19, 0x64, 0x45, 0x10, 0x2f, 0x54, 0x5b, 0x30, 0x58, 0x97,
	0x95, 0x03, 0x17, 0x37, 0xcc, 0x9c, 0xcc, 0xbc, 0x12, 0x42, 0x86, 0x5c, 0x04, 0x1b, 0xc2, 0x1a,
	0xc4, 0x05, 0xd5, 0xe3, 0x99, 0x57, 0x92, 0xc4, 0x06, 0x56, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x8c, 0x0b, 0x48, 0x0f, 0x01, 0x00, 0x00,
}
