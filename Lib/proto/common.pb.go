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
	Aid                  string   `protobuf:"bytes,3,opt,name=aid,proto3" json:"aid,omitempty"`
	AccessToken          string   `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	State                string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	ReceivingTime        string   `protobuf:"bytes,8,opt,name=receiving_time,json=receivingTime,proto3" json:"receiving_time,omitempty"`
	PrintTime            string   `protobuf:"bytes,9,opt,name=print_time,json=printTime,proto3" json:"print_time,omitempty"`
	Status               string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Content              string   `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
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

func (m *Body) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *Body) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Body) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Body) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Body) GetReceivingTime() string {
	if m != nil {
		return m.ReceivingTime
	}
	return ""
}

func (m *Body) GetPrintTime() string {
	if m != nil {
		return m.PrintTime
	}
	return ""
}

func (m *Body) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Body) GetContent() string {
	if m != nil {
		return m.Content
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
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x91, 0x41, 0x6b, 0xe3, 0x30,
	0x10, 0x85, 0x89, 0x63, 0x7b, 0xe3, 0x71, 0x12, 0x82, 0x08, 0x8b, 0x76, 0x4b, 0x69, 0x9a, 0x52,
	0xc8, 0xa1, 0x38, 0x90, 0xfe, 0x82, 0x96, 0xf6, 0xd0, 0x43, 0x2f, 0x26, 0xf7, 0xe0, 0x48, 0x13,
	0x57, 0x34, 0x96, 0x82, 0xa4, 0x14, 0xfc, 0x5b, 0xfa, 0x67, 0x8b, 0x47, 0x4e, 0x4f, 0xd6, 0xfb,
	0xde, 0x9b, 0x31, 0xbc, 0x01, 0x76, 0xb2, 0xc6, 0x9b, 0xb5, 0x30, 0x4d, 0x63, 0x74, 0x41, 0x82,
	0x25, 0xf4, 0xf9, 0xff, 0xaf, 0x36, 0xa6, 0x3e, 0xe2, 0x9a, 0xd4, 0xfe, 0x7c, 0x58, 0x57, 0xba,
	0x0d, 0x89, 0xe5, 0x3b, 0xcc, 0x5e, 0xad, 0x35, 0xf6, 0x05, 0x9d, 0xb0, 0xea, 0xe4, 0x95, 0xd1,
	0x8c, 0xc3, 0x1f, 0x77, 0x16, 0x02, 0x9d, 0xe3, 0x83, 0xc5, 0x60, 0x95, 0x95, 0x17, 0xc9, 0x6e,
	0x20, 0xde, 0x1b, 0xd9, 0xf2, 0x68, 0x31, 0x58, 0xe5, 0x9b, 0x3c, 0xec, 0x28, 0x9e, 0x8d, 0x6c,
	0x4b, 0x32, 0x96, 0xdf, 0x11, 0xc4, 0x9d, 0x64, 0x53, 0x88, 0x94, 0xec, 0xc7, 0x23, 0x25, 0xd9,
	0x15, 0x64, 0xc6, 0xaa, 0x5a, 0xe9, 0x9d, 0x92, 0x34, 0x9e, 0x95, 0xa3, 0x00, 0xde, 0x24, 0x9b,
	0xc1, 0xb0, 0x52, 0x92, 0x0f, 0x09, 0x77, 0x4f, 0x76, 0x0b, 0xe3, 0x8a, 0x7e, 0xb9, 0xf3, 0xe6,
	0x13, 0x35, 0x8f, 0xc9, 0xca, 0x03, 0xdb, 0x76, 0x88, 0xdd, 0xc1, 0xc4, 0xe2, 0xc1, 0xa2, 0xfb,
	0xe8, 0x33, 0x09, 0x65, 0xc6, 0x3d, 0x0c, 0xa1, 0x39, 0x24, 0xce, 0x57, 0x1e, 0x79, 0x4a, 0x66,
	0x10, 0xec, 0x1e, 0xa6, 0x16, 0x05, 0xaa, 0x2f, 0xa5, 0xeb, 0x9d, 0x57, 0x0d, 0xf2, 0x11, 0xd9,
	0x93, 0x5f, 0xba, 0x55, 0x0d, 0xb2, 0x6b, 0x80, 0x93, 0x55, 0xda, 0x87, 0x48, 0x46, 0x91, 0x8c,
	0x08, 0xd9, 0x7f, 0x21, 0xed, 0xd6, 0x9d, 0x1d, 0x07, 0xb2, 0x7a, 0xd5, 0xd5, 0x27, 0x8c, 0xf6,
	0xa8, 0x3d, 0xcf, 0x43, 0x7d, 0xbd, 0x5c, 0x0a, 0x48, 0xa8, 0x6c, 0xc6, 0x20, 0x16, 0x46, 0x22,
	0xf5, 0x93, 0x94, 0xf4, 0xee, 0xc6, 0x1a, 0x74, 0xae, 0xaa, 0xb1, 0xef, 0xe7, 0x22, 0xd9, 0x03,
	0xa4, 0x12, 0x7d, 0xa5, 0x8e, 0xd4, 0x50, 0xbe, 0x99, 0x17, 0xe1, 0x9e, 0xc5, 0xe5, 0x9e, 0xc5,
	0x93, 0x6e, 0xcb, 0x3e, 0xb3, 0x4f, 0x89, 0x3e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xc1,
	0x3d, 0x35, 0x10, 0x02, 0x00, 0x00,
}
