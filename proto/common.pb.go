//@generated
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/protobuf/common.proto

package proto // import "github.com/yuwenhui/gin-starter/proto"

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

type Pagination struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	PageMax              int32    `protobuf:"varint,3,opt,name=page_max,json=pageMax,proto3" json:"page_max"`
	Total                int32    `protobuf:"varint,4,opt,name=total,proto3" json:"total"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_1d5a43586b511b4a, []int{0}
}
func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (dst *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(dst, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pagination) GetPageMax() int32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *Pagination) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type TaskItem struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskItem) Reset()         { *m = TaskItem{} }
func (m *TaskItem) String() string { return proto.CompactTextString(m) }
func (*TaskItem) ProtoMessage()    {}
func (*TaskItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_1d5a43586b511b4a, []int{1}
}
func (m *TaskItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskItem.Unmarshal(m, b)
}
func (m *TaskItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskItem.Marshal(b, m, deterministic)
}
func (dst *TaskItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskItem.Merge(dst, src)
}
func (m *TaskItem) XXX_Size() int {
	return xxx_messageInfo_TaskItem.Size(m)
}
func (m *TaskItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskItem.DiscardUnknown(m)
}

var xxx_messageInfo_TaskItem proto.InternalMessageInfo

func (m *TaskItem) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TaskItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskItem) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func init() {
	proto.RegisterType((*Pagination)(nil), "common.Pagination")
	proto.RegisterType((*TaskItem)(nil), "common.TaskItem")
}

func init() { proto.RegisterFile("proto/protobuf/common.proto", fileDescriptor_common_1d5a43586b511b4a) }

var fileDescriptor_common_1d5a43586b511b4a = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xa9, 0xb6, 0xbb, 0xdd, 0x39, 0x0e, 0x82, 0x91, 0x5e, 0xa4, 0x20, 0x7a, 0xb1, 0x7b,
	0xf0, 0x1f, 0x78, 0xeb, 0x41, 0x90, 0xad, 0x27, 0x2f, 0xcb, 0xb4, 0x8d, 0xdb, 0xc1, 0x26, 0x29,
	0xc9, 0x04, 0xdb, 0xfe, 0x7a, 0xd9, 0xd9, 0xf6, 0x92, 0xf7, 0xbe, 0xf7, 0x08, 0x0f, 0x06, 0x66,
	0x87, 0x18, 0x24, 0xd4, 0xfa, 0xae, 0xf3, 0x4f, 0xbd, 0x09, 0xce, 0x05, 0xbf, 0x50, 0xc6, 0x62,
	0xa0, 0xb9, 0x07, 0xf8, 0xa4, 0x8e, 0x3d, 0x09, 0x07, 0x8f, 0x08, 0xe3, 0x03, 0x75, 0xd6, 0x8c,
	0x1e, 0x47, 0x2f, 0x93, 0x46, 0x3d, 0xce, 0xa0, 0xea, 0xb5, 0x4d, 0x7c, 0xb6, 0xe6, 0x46, 0x8b,
	0x69, 0x1f, 0xac, 0xf8, 0x6c, 0xf1, 0x01, 0xd4, 0xb7, 0x8e, 0x8e, 0xe6, 0x56, 0xbb, 0xb2, 0xe7,
	0x0f, 0x3a, 0xe2, 0x1d, 0x4c, 0x24, 0x08, 0xed, 0xcd, 0x58, 0xf3, 0x01, 0xe6, 0x2b, 0x98, 0x7e,
	0x51, 0xfa, 0x5d, 0x8a, 0x75, 0x78, 0x0f, 0x65, 0x4e, 0x36, 0xb6, 0xbc, 0xbd, 0x0c, 0x16, 0x3d,
	0x2e, 0xb7, 0xfa, 0x95, 0x65, 0x3f, 0xcc, 0x55, 0xcd, 0x00, 0x68, 0xa0, 0x4c, 0xd9, 0x39, 0x8a,
	0x27, 0x9d, 0xaa, 0x9a, 0x2b, 0xbe, 0x3f, 0x7f, 0x3f, 0x75, 0x2c, 0xbb, 0xbc, 0x5e, 0x6c, 0x82,
	0xab, 0x4f, 0xf9, 0xcf, 0xfa, 0x5d, 0xe6, 0xba, 0x63, 0xff, 0x9a, 0x84, 0xa2, 0xd8, 0x78, 0xb9,
	0x42, 0xa1, 0xf2, 0xf6, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xca, 0x46, 0xea, 0xfb, 0x1b, 0x01, 0x00,
	0x00,
}
