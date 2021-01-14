// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package account

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

type AccountInfo struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{0}
}

func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type NewAccountInfo struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAccountInfo) Reset()         { *m = NewAccountInfo{} }
func (m *NewAccountInfo) String() string { return proto.CompactTextString(m) }
func (*NewAccountInfo) ProtoMessage()    {}
func (*NewAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{1}
}

func (m *NewAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountInfo.Unmarshal(m, b)
}
func (m *NewAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountInfo.Marshal(b, m, deterministic)
}
func (m *NewAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountInfo.Merge(m, src)
}
func (m *NewAccountInfo) XXX_Size() int {
	return xxx_messageInfo_NewAccountInfo.Size(m)
}
func (m *NewAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountInfo proto.InternalMessageInfo

func (m *NewAccountInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *NewAccountInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AccountResponse struct {
	StatusCode           uint32   `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	JWTtoken             string   `protobuf:"bytes,2,opt,name=JWTtoken,proto3" json:"JWTtoken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{2}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *AccountResponse) GetJWTtoken() string {
	if m != nil {
		return m.JWTtoken
	}
	return ""
}

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{4}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{6}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{7}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{8}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{9}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountInfo)(nil), "account.AccountInfo")
	proto.RegisterType((*NewAccountInfo)(nil), "account.NewAccountInfo")
	proto.RegisterType((*AccountResponse)(nil), "account.AccountResponse")
	proto.RegisterType((*Message)(nil), "account.Message")
	proto.RegisterType((*Request)(nil), "account.Request")
	proto.RegisterType((*Response)(nil), "account.Response")
	proto.RegisterType((*StreamingRequest)(nil), "account.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "account.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "account.Ping")
	proto.RegisterType((*Pong)(nil), "account.Pong")
}

func init() {
	proto.RegisterFile("proto/account.proto", fileDescriptor_477cbf5ae5b46edf)
}

var fileDescriptor_477cbf5ae5b46edf = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xc0, 0x47, 0xf9, 0xe6, 0xfb, 0x2a, 0xb0, 0x12, 0xc5, 0xaa, 0xc4, 0xec, 0x09, 0x0f,
	0xa2, 0xd1, 0xa3, 0x26, 0x0a, 0x78, 0xd1, 0x04, 0x63, 0xaa, 0xc6, 0xc4, 0xdb, 0x0a, 0x6b, 0x6d,
	0xa4, 0xdd, 0xda, 0xdd, 0x86, 0xf8, 0xbb, 0xfd, 0x03, 0x6e, 0xb7, 0xdb, 0xa6, 0x12, 0xf4, 0xa0,
	0xb7, 0x79, 0xf3, 0xde, 0xbc, 0xbc, 0x99, 0x6e, 0x61, 0x35, 0x8c, 0x98, 0x60, 0xfb, 0x64, 0x32,
	0x61, 0x71, 0x20, 0xfa, 0x0a, 0x21, 0x53, 0x43, 0x7c, 0x0a, 0xff, 0x06, 0x69, 0x79, 0x11, 0x3c,
	0x31, 0xd4, 0x86, 0x3f, 0xd4, 0x27, 0xde, 0xac, 0x63, 0xec, 0x18, 0xbd, 0xbf, 0x4e, 0x0a, 0x90,
	0x0d, 0xf5, 0x90, 0x70, 0x3e, 0x67, 0xd1, 0xb4, 0x53, 0x56, 0x44, 0x8e, 0xf1, 0x10, 0x56, 0xae,
	0xe8, 0xfc, 0x77, 0x1e, 0x63, 0x68, 0x68, 0x03, 0x87, 0xf2, 0x90, 0x05, 0x9c, 0xa2, 0x2e, 0x00,
	0x17, 0x44, 0xc4, 0x7c, 0xc4, 0xa6, 0x54, 0x39, 0x59, 0x4e, 0xa1, 0x93, 0xd8, 0x5d, 0xde, 0xdf,
	0x0a, 0xf6, 0x42, 0x83, 0xcc, 0x2e, 0xc3, 0x78, 0x13, 0xcc, 0x31, 0xe5, 0x9c, 0xb8, 0x14, 0x35,
	0xa1, 0xc2, 0xc9, 0x9b, 0x4e, 0x92, 0x94, 0x78, 0x1b, 0x4c, 0x87, 0xbe, 0xc6, 0x94, 0x0b, 0x84,
	0xa0, 0x1a, 0x10, 0x9f, 0x6a, 0x56, 0xd5, 0x78, 0x0b, 0xea, 0x79, 0x06, 0x39, 0xec, 0x73, 0x37,
	0x1b, 0x96, 0x25, 0xee, 0x41, 0xf3, 0x46, 0x44, 0x94, 0xf8, 0x5e, 0xe0, 0x66, 0x2e, 0x72, 0x5d,
	0x15, 0x5d, 0xe9, 0x2a, 0x4e, 0x0a, 0xf0, 0x2e, 0xb4, 0x0a, 0x4a, 0x6d, 0xb8, 0x5c, 0xda, 0x85,
	0xea, 0xb5, 0x54, 0xa1, 0x35, 0xa8, 0x71, 0x11, 0xc9, 0x0d, 0x34, 0xad, 0x91, 0xe2, 0xd9, 0xd7,
	0xfc, 0xe1, 0x7b, 0x19, 0x4c, 0x7d, 0x3e, 0xb4, 0x07, 0xd5, 0x11, 0x99, 0xcd, 0x50, 0xb3, 0x9f,
	0x7d, 0x6f, 0x1d, 0xd3, 0x6e, 0x15, 0x3a, 0x69, 0x1c, 0x5c, 0x42, 0x23, 0xa8, 0xa5, 0x29, 0xd1,
	0x46, 0x4e, 0x2f, 0x2e, 0x68, 0xdb, 0xcb, 0xa8, 0xcc, 0xe2, 0xc0, 0x40, 0x7d, 0xa8, 0x27, 0xf9,
	0x55, 0x46, 0x2b, 0xd7, 0x26, 0x2d, 0xbb, 0x00, 0x25, 0x8b, 0x4b, 0x3d, 0x43, 0xea, 0x4f, 0x92,
	0x13, 0xbb, 0x1e, 0x17, 0x34, 0x42, 0xed, 0x5c, 0x50, 0x78, 0x41, 0x76, 0x67, 0xb1, 0x5b, 0x88,
	0x7c, 0x06, 0xff, 0x07, 0xb1, 0x78, 0xa6, 0x81, 0xf0, 0x26, 0x44, 0xd0, 0x1f, 0x38, 0x9c, 0x83,
	0x75, 0x17, 0x4e, 0xe5, 0x6c, 0x76, 0xb4, 0xf5, 0x5c, 0xfc, 0xf9, 0x25, 0x7f, 0xe7, 0x32, 0x6c,
	0x3c, 0x58, 0xea, 0x57, 0x3a, 0xd6, 0x92, 0xc7, 0x9a, 0x82, 0x47, 0x1f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x99, 0x46, 0x52, 0x70, 0x03, 0x00, 0x00,
}
