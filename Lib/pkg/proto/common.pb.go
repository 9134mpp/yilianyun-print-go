// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type ErrorDescription struct {
	Success              string   `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	Body                 *Body    `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorDescription) Reset()         { *m = ErrorDescription{} }
func (m *ErrorDescription) String() string { return proto.CompactTextString(m) }
func (*ErrorDescription) ProtoMessage()    {}
func (*ErrorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *ErrorDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorDescription.Unmarshal(m, b)
}
func (m *ErrorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorDescription.Marshal(b, m, deterministic)
}
func (m *ErrorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorDescription.Merge(m, src)
}
func (m *ErrorDescription) XXX_Size() int {
	return xxx_messageInfo_ErrorDescription.Size(m)
}
func (m *ErrorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorDescription proto.InternalMessageInfo

func (m *ErrorDescription) GetSuccess() string {
	if m != nil {
		return m.Success
	}
	return ""
}

func (m *ErrorDescription) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type Body struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OriginId             string   `protobuf:"bytes,2,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

func (m *Body) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Body) GetOriginId() string {
	if m != nil {
		return m.OriginId
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail               *any.Any `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDetail() *any.Any {
	if m != nil {
		return m.Detail
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorDescription)(nil), "proto.ErrorDescription")
	proto.RegisterType((*Body)(nil), "proto.Body")
	proto.RegisterType((*Error)(nil), "proto.Error")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x90, 0x14, 0x72, 0x95, 0x10, 0x3a, 0x31, 0x04, 0x18, 0xa8, 0x32, 0x75, 0x40,
	0x8e, 0x44, 0x9f, 0x00, 0x04, 0x03, 0x03, 0x8b, 0x5f, 0x00, 0x25, 0xb6, 0x89, 0x4e, 0x6a, 0x7c,
	0x95, 0x9d, 0x0e, 0x79, 0xfb, 0xaa, 0xe7, 0x64, 0xb2, 0xbf, 0xf3, 0x7d, 0xd6, 0xff, 0x03, 0x9e,
	0x02, 0x4f, 0xdc, 0x1a, 0x1e, 0x47, 0xf6, 0x4a, 0x00, 0x4b, 0x39, 0x9e, 0x9f, 0x06, 0xe6, 0xe1,
	0xe8, 0x5a, 0xa1, 0xfe, 0xfc, 0xdf, 0x76, 0x7e, 0x4e, 0x1b, 0xcd, 0x2f, 0x3c, 0x7c, 0x87, 0xc0,
	0xe1, 0xcb, 0x45, 0x13, 0xe8, 0x34, 0x11, 0x7b, 0xac, 0xe1, 0x36, 0x9e, 0x8d, 0x71, 0x31, 0xd6,
	0xd9, 0x2e, 0xdb, 0x57, 0x7a, 0x45, 0x7c, 0x85, 0xa2, 0x67, 0x3b, 0xd7, 0xf9, 0x2e, 0xdb, 0x6f,
	0xdf, 0xb7, 0xe9, 0x0f, 0xf5, 0xc9, 0x76, 0xd6, 0xf2, 0xd0, 0x1c, 0xa0, 0xb8, 0x12, 0xde, 0x43,
	0x4e, 0x76, 0xb1, 0x73, 0xb2, 0xf8, 0x02, 0x15, 0x07, 0x1a, 0xc8, 0xff, 0x91, 0x15, 0xbb, 0xd2,
	0x77, 0x69, 0xf0, 0x63, 0x1b, 0x03, 0xa5, 0x64, 0x40, 0x84, 0xc2, 0xb0, 0x75, 0xe2, 0x95, 0x5a,
	0xee, 0xd7, 0x30, 0xa3, 0x8b, 0xb1, 0x1b, 0xdc, 0xe2, 0xad, 0x88, 0x6f, 0xb0, 0xb1, 0x6e, 0xea,
	0xe8, 0x58, 0xdf, 0x48, 0x9c, 0x47, 0x95, 0x6a, 0xaa, 0xb5, 0xa6, 0xfa, 0xf0, 0xb3, 0x5e, 0x76,
	0xfa, 0x8d, 0x4c, 0x0f, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x94, 0xb6, 0x49, 0x27, 0x01,
	0x00, 0x00,
}
