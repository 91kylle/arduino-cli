// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package commands

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

type Result struct {
	Failed               bool     `protobuf:"varint,1,opt,name=failed,proto3" json:"failed,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dff099eb2e3dfdb, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "arduino.commands.Result")
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_0dff099eb2e3dfdb) }

var fileDescriptor_0dff099eb2e3dfdb = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0x2c, 0x4a, 0x29,
	0xcd, 0xcc, 0xcb, 0xd7, 0x83, 0x89, 0x2b, 0xf9, 0x71, 0xb1, 0x05, 0xa5, 0x16, 0x97, 0xe6, 0x94,
	0x08, 0x89, 0x71, 0xb1, 0xa5, 0x25, 0x66, 0xe6, 0xa4, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x04, 0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0xac, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04,
	0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xeb, 0xa4, 0x19, 0xa5, 0x9e, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0x04, 0xb2, 0x42, 0x1f, 0x6a, 0x1d, 0x8c, 0xd6, 0x4d, 0xce, 0xc9, 0xd4, 0x87, 0x59, 0x9d,
	0xc4, 0x06, 0x76, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x21, 0x74, 0x8e, 0x9e, 0xa5, 0x00,
	0x00, 0x00,
}
