// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testopesto/testopesto.proto

package testopesto

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

type GetTimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b770b6ace7b7118, []int{0}
}

func (m *GetTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeRequest.Unmarshal(m, b)
}
func (m *GetTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeRequest.Merge(m, src)
}
func (m *GetTimeRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeRequest.Size(m)
}
func (m *GetTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeRequest proto.InternalMessageInfo

type GetTimeReply struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeReply) Reset()         { *m = GetTimeReply{} }
func (m *GetTimeReply) String() string { return proto.CompactTextString(m) }
func (*GetTimeReply) ProtoMessage()    {}
func (*GetTimeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b770b6ace7b7118, []int{1}
}

func (m *GetTimeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeReply.Unmarshal(m, b)
}
func (m *GetTimeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeReply.Marshal(b, m, deterministic)
}
func (m *GetTimeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeReply.Merge(m, src)
}
func (m *GetTimeReply) XXX_Size() int {
	return xxx_messageInfo_GetTimeReply.Size(m)
}
func (m *GetTimeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeReply proto.InternalMessageInfo

func (m *GetTimeReply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTimeRequest)(nil), "testopesto.GetTimeRequest")
	proto.RegisterType((*GetTimeReply)(nil), "testopesto.GetTimeReply")
}

func init() {
	proto.RegisterFile("testopesto/testopesto.proto", fileDescriptor_0b770b6ace7b7118)
}

var fileDescriptor_0b770b6ace7b7118 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x49, 0x2d, 0x2e,
	0xc9, 0x2f, 0x00, 0x11, 0xfa, 0x08, 0xa6, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x10, 0x17, 0x42,
	0x44, 0x49, 0x80, 0x8b, 0xcf, 0x3d, 0xb5, 0x24, 0x24, 0x33, 0x37, 0x35, 0x28, 0xb5, 0xb0, 0x34,
	0xb5, 0xb8, 0x44, 0x49, 0x85, 0x8b, 0x07, 0x2e, 0x52, 0x90, 0x53, 0x29, 0x24, 0xc2, 0xc5, 0x5a,
	0x04, 0x62, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x46, 0xc1, 0x5c, 0x7c, 0xe9,
	0xa9, 0x25, 0xce, 0xa5, 0x45, 0x45, 0xa9, 0x79, 0x60, 0xc5, 0x42, 0x8e, 0x5c, 0xec, 0xe9, 0x10,
	0x7d, 0x42, 0x52, 0x7a, 0x48, 0x76, 0xa2, 0x1a, 0x2f, 0x25, 0x81, 0x55, 0xae, 0x20, 0xa7, 0x52,
	0x89, 0xc1, 0xc9, 0x34, 0xca, 0x38, 0xa3, 0xa4, 0xa4, 0xa0, 0xd8, 0x4a, 0x5f, 0x3f, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xd8, 0x25, 0xd5, 0x27, 0x31, 0x2c, 0x33,
	0x25, 0x51, 0x3f, 0x25, 0xb3, 0xb8, 0xb2, 0x58, 0xb7, 0x3c, 0x35, 0x35, 0xdb, 0xd4, 0x1a, 0x61,
	0x4a, 0x12, 0x1b, 0xd8, 0x5b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xbd, 0x09, 0xca,
	0xf5, 0x00, 0x00, 0x00,
}
