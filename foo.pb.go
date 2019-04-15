// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package main

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

type Foo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Bar                  *Foo_Bar `protobuf:"bytes,3,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Foo) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Foo) GetBar() *Foo_Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

type Foo_Bar struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo_Bar) Reset()         { *m = Foo_Bar{} }
func (m *Foo_Bar) String() string { return proto.CompactTextString(m) }
func (*Foo_Bar) ProtoMessage()    {}
func (*Foo_Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0, 0}
}

func (m *Foo_Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo_Bar.Unmarshal(m, b)
}
func (m *Foo_Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo_Bar.Marshal(b, m, deterministic)
}
func (m *Foo_Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo_Bar.Merge(m, src)
}
func (m *Foo_Bar) XXX_Size() int {
	return xxx_messageInfo_Foo_Bar.Size(m)
}
func (m *Foo_Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Foo_Bar proto.InternalMessageInfo

func (m *Foo_Bar) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Foo)(nil), "main.Foo")
	proto.RegisterType((*Foo_Bar)(nil), "main.Foo.Bar")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x4a, 0xe7, 0x62, 0x76,
	0xcb, 0xcf, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0xf3, 0x4a, 0x73, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58,
	0x83, 0x40, 0x4c, 0x21, 0x79, 0x2e, 0xe6, 0xa4, 0xc4, 0x22, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e,
	0x23, 0x5e, 0x3d, 0x90, 0x01, 0x7a, 0x6e, 0xf9, 0xf9, 0x7a, 0x4e, 0x89, 0x45, 0x41, 0x20, 0x19,
	0x29, 0x49, 0x2e, 0x66, 0xa7, 0xc4, 0x22, 0x6c, 0xa6, 0x25, 0xb1, 0x81, 0x6d, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xcd, 0xee, 0xe2, 0x82, 0x00, 0x00, 0x00,
}
