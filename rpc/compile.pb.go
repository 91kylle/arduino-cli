// Code generated by protoc-gen-go. DO NOT EDIT.
// source: compile.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

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

type CompileReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	SketchPath           string    `protobuf:"bytes,3,opt,name=sketchPath,proto3" json:"sketchPath,omitempty"`
	ShowProperties       bool      `protobuf:"varint,4,opt,name=showProperties,proto3" json:"showProperties,omitempty"`
	Preprocess           bool      `protobuf:"varint,5,opt,name=preprocess,proto3" json:"preprocess,omitempty"`
	BuildCachePath       string    `protobuf:"bytes,6,opt,name=buildCachePath,proto3" json:"buildCachePath,omitempty"`
	BuildPath            string    `protobuf:"bytes,7,opt,name=buildPath,proto3" json:"buildPath,omitempty"`
	BuildProperties      []string  `protobuf:"bytes,8,rep,name=buildProperties,proto3" json:"buildProperties,omitempty"`
	Warnings             string    `protobuf:"bytes,9,opt,name=warnings,proto3" json:"warnings,omitempty"`
	Verbose              bool      `protobuf:"varint,10,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Quiet                bool      `protobuf:"varint,11,opt,name=quiet,proto3" json:"quiet,omitempty"`
	VidPid               string    `protobuf:"bytes,12,opt,name=vidPid,proto3" json:"vidPid,omitempty"`
	ExportFile           string    `protobuf:"bytes,13,opt,name=exportFile,proto3" json:"exportFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CompileReq) Reset()         { *m = CompileReq{} }
func (m *CompileReq) String() string { return proto.CompactTextString(m) }
func (*CompileReq) ProtoMessage()    {}
func (*CompileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_5a0f6dfd9804c4ee, []int{0}
}
func (m *CompileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileReq.Unmarshal(m, b)
}
func (m *CompileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileReq.Marshal(b, m, deterministic)
}
func (dst *CompileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileReq.Merge(dst, src)
}
func (m *CompileReq) XXX_Size() int {
	return xxx_messageInfo_CompileReq.Size(m)
}
func (m *CompileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileReq.DiscardUnknown(m)
}

var xxx_messageInfo_CompileReq proto.InternalMessageInfo

func (m *CompileReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CompileReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *CompileReq) GetSketchPath() string {
	if m != nil {
		return m.SketchPath
	}
	return ""
}

func (m *CompileReq) GetShowProperties() bool {
	if m != nil {
		return m.ShowProperties
	}
	return false
}

func (m *CompileReq) GetPreprocess() bool {
	if m != nil {
		return m.Preprocess
	}
	return false
}

func (m *CompileReq) GetBuildCachePath() string {
	if m != nil {
		return m.BuildCachePath
	}
	return ""
}

func (m *CompileReq) GetBuildPath() string {
	if m != nil {
		return m.BuildPath
	}
	return ""
}

func (m *CompileReq) GetBuildProperties() []string {
	if m != nil {
		return m.BuildProperties
	}
	return nil
}

func (m *CompileReq) GetWarnings() string {
	if m != nil {
		return m.Warnings
	}
	return ""
}

func (m *CompileReq) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func (m *CompileReq) GetQuiet() bool {
	if m != nil {
		return m.Quiet
	}
	return false
}

func (m *CompileReq) GetVidPid() string {
	if m != nil {
		return m.VidPid
	}
	return ""
}

func (m *CompileReq) GetExportFile() string {
	if m != nil {
		return m.ExportFile
	}
	return ""
}

type CompileResp struct {
	Output               []byte        `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	TaskProgress         *TaskProgress `protobuf:"bytes,2,opt,name=task_progress,json=taskProgress,proto3" json:"task_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CompileResp) Reset()         { *m = CompileResp{} }
func (m *CompileResp) String() string { return proto.CompactTextString(m) }
func (*CompileResp) ProtoMessage()    {}
func (*CompileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_compile_5a0f6dfd9804c4ee, []int{1}
}
func (m *CompileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompileResp.Unmarshal(m, b)
}
func (m *CompileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompileResp.Marshal(b, m, deterministic)
}
func (dst *CompileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompileResp.Merge(dst, src)
}
func (m *CompileResp) XXX_Size() int {
	return xxx_messageInfo_CompileResp.Size(m)
}
func (m *CompileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CompileResp.DiscardUnknown(m)
}

var xxx_messageInfo_CompileResp proto.InternalMessageInfo

func (m *CompileResp) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *CompileResp) GetTaskProgress() *TaskProgress {
	if m != nil {
		return m.TaskProgress
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileReq)(nil), "arduino.CompileReq")
	proto.RegisterType((*CompileResp)(nil), "arduino.CompileResp")
}

func init() { proto.RegisterFile("compile.proto", fileDescriptor_compile_5a0f6dfd9804c4ee) }

var fileDescriptor_compile_5a0f6dfd9804c4ee = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x6f, 0xe2, 0x30,
	0x10, 0xc5, 0xc5, 0xff, 0x30, 0x84, 0x5d, 0xad, 0xb5, 0xbb, 0xb2, 0x50, 0x55, 0x21, 0x54, 0x55,
	0x5c, 0x08, 0x52, 0x7b, 0xeb, 0xb1, 0x48, 0x95, 0x7a, 0x8b, 0xa2, 0x9e, 0x7a, 0xa9, 0xf2, 0xc7,
	0x25, 0x16, 0x60, 0x1b, 0xdb, 0x81, 0x7e, 0xa7, 0x7e, 0xc9, 0x3a, 0x93, 0x10, 0x10, 0x27, 0x78,
	0xbf, 0x79, 0x1e, 0x4f, 0xe6, 0x19, 0xc6, 0xa9, 0xdc, 0x29, 0xbe, 0x65, 0x81, 0xd2, 0xd2, 0x4a,
	0x32, 0x88, 0x75, 0x56, 0x70, 0x21, 0x27, 0xbe, 0xe3, 0x3b, 0x29, 0x2a, 0x3c, 0xfb, 0xee, 0x00,
	0xac, 0x2a, 0x63, 0xc4, 0xf6, 0x64, 0x01, 0x1e, 0x17, 0xc6, 0xc6, 0x22, 0x65, 0xb4, 0x35, 0x6d,
	0xcd, 0x47, 0x0f, 0x7f, 0x82, 0xfa, 0x60, 0xf0, 0x5a, 0x17, 0xa2, 0xc6, 0x42, 0x08, 0x74, 0x3f,
	0xf7, 0x89, 0xa0, 0x6d, 0x67, 0x1d, 0x46, 0xf8, 0x9f, 0xdc, 0x02, 0x98, 0x0d, 0xb3, 0x69, 0x1e,
	0xc6, 0x36, 0xa7, 0x1d, 0xac, 0x5c, 0x10, 0x72, 0x0f, 0xbf, 0x4c, 0x2e, 0x8f, 0xa1, 0x96, 0x8a,
	0x69, 0xcb, 0x99, 0xa1, 0x5d, 0xe7, 0xf1, 0xa2, 0x2b, 0x5a, 0xf6, 0x51, 0x9a, 0xb9, 0x29, 0x53,
	0x66, 0x0c, 0xed, 0xa1, 0xe7, 0x82, 0x94, 0x7d, 0x92, 0x82, 0x6f, 0xb3, 0x55, 0x9c, 0xe6, 0x0c,
	0xef, 0xea, 0xe3, 0x5d, 0x57, 0x94, 0xdc, 0xc0, 0x10, 0x09, 0x5a, 0x06, 0x68, 0x39, 0x03, 0x32,
	0x87, 0xdf, 0x95, 0x38, 0x8f, 0xe3, 0x4d, 0x3b, 0xce, 0x73, 0x8d, 0xc9, 0x04, 0xbc, 0x63, 0xac,
	0x05, 0x17, 0x6b, 0x43, 0x87, 0xd8, 0xa6, 0xd1, 0x84, 0xc2, 0xe0, 0xc0, 0x74, 0x22, 0x0d, 0xa3,
	0x80, 0x83, 0x9e, 0x24, 0xf9, 0x0b, 0xbd, 0x7d, 0xc1, 0x99, 0xa5, 0x23, 0xe4, 0x95, 0x20, 0xff,
	0xa1, 0x7f, 0xe0, 0x59, 0xc8, 0x33, 0xea, 0x63, 0xa7, 0x5a, 0x95, 0xdf, 0xcc, 0xbe, 0x94, 0xd4,
	0xf6, 0xc5, 0xe5, 0x41, 0xc7, 0xd5, 0xee, 0xce, 0x64, 0x16, 0xc3, 0xa8, 0x09, 0xcb, 0xa8, 0xb2,
	0x8d, 0x2c, 0xac, 0x2a, 0x2c, 0x66, 0xe5, 0x47, 0xb5, 0x22, 0x4f, 0x30, 0xb6, 0xb1, 0xd9, 0x7c,
	0xb8, 0x55, 0xad, 0x75, 0xb9, 0xbd, 0x36, 0x46, 0xf9, 0xaf, 0x89, 0xf2, 0xcd, 0x55, 0xc3, 0xba,
	0x18, 0xf9, 0xf6, 0x42, 0x3d, 0xdf, 0xbd, 0xcf, 0xd6, 0xdc, 0xe6, 0x45, 0x12, 0xb8, 0x77, 0xb2,
	0xac, 0x0f, 0x9c, 0x7e, 0x17, 0xe9, 0x96, 0x2f, 0xb5, 0x4a, 0x93, 0x3e, 0xbe, 0x9e, 0xc7, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x6c, 0x56, 0xa9, 0x65, 0x02, 0x00, 0x00,
}
