// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_status_codes.proto

package Ydb

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

// reserved range [400000, 400999]
type StatusIds_StatusCode int32

const (
	StatusIds_STATUS_CODE_UNSPECIFIED StatusIds_StatusCode = 0
	StatusIds_SUCCESS                 StatusIds_StatusCode = 400000
	StatusIds_BAD_REQUEST             StatusIds_StatusCode = 400010
	StatusIds_UNAUTHORIZED            StatusIds_StatusCode = 400020
	StatusIds_INTERNAL_ERROR          StatusIds_StatusCode = 400030
	StatusIds_ABORTED                 StatusIds_StatusCode = 400040
	StatusIds_UNAVAILABLE             StatusIds_StatusCode = 400050
	StatusIds_OVERLOADED              StatusIds_StatusCode = 400060
	StatusIds_SCHEME_ERROR            StatusIds_StatusCode = 400070
	StatusIds_GENERIC_ERROR           StatusIds_StatusCode = 400080
	StatusIds_TIMEOUT                 StatusIds_StatusCode = 400090
	StatusIds_BAD_SESSION             StatusIds_StatusCode = 400100
	StatusIds_PRECONDITION_FAILED     StatusIds_StatusCode = 400120
	StatusIds_ALREADY_EXISTS          StatusIds_StatusCode = 400130
	StatusIds_NOT_FOUND               StatusIds_StatusCode = 400140
	StatusIds_SESSION_EXPIRED         StatusIds_StatusCode = 400150
	StatusIds_CANCELLED               StatusIds_StatusCode = 400160
	StatusIds_UNDETERMINED            StatusIds_StatusCode = 400170
	StatusIds_UNSUPPORTED             StatusIds_StatusCode = 400180
)

var StatusIds_StatusCode_name = map[int32]string{
	0:      "STATUS_CODE_UNSPECIFIED",
	400000: "SUCCESS",
	400010: "BAD_REQUEST",
	400020: "UNAUTHORIZED",
	400030: "INTERNAL_ERROR",
	400040: "ABORTED",
	400050: "UNAVAILABLE",
	400060: "OVERLOADED",
	400070: "SCHEME_ERROR",
	400080: "GENERIC_ERROR",
	400090: "TIMEOUT",
	400100: "BAD_SESSION",
	400120: "PRECONDITION_FAILED",
	400130: "ALREADY_EXISTS",
	400140: "NOT_FOUND",
	400150: "SESSION_EXPIRED",
	400160: "CANCELLED",
	400170: "UNDETERMINED",
	400180: "UNSUPPORTED",
}

var StatusIds_StatusCode_value = map[string]int32{
	"STATUS_CODE_UNSPECIFIED": 0,
	"SUCCESS":                 400000,
	"BAD_REQUEST":             400010,
	"UNAUTHORIZED":            400020,
	"INTERNAL_ERROR":          400030,
	"ABORTED":                 400040,
	"UNAVAILABLE":             400050,
	"OVERLOADED":              400060,
	"SCHEME_ERROR":            400070,
	"GENERIC_ERROR":           400080,
	"TIMEOUT":                 400090,
	"BAD_SESSION":             400100,
	"PRECONDITION_FAILED":     400120,
	"ALREADY_EXISTS":          400130,
	"NOT_FOUND":               400140,
	"SESSION_EXPIRED":         400150,
	"CANCELLED":               400160,
	"UNDETERMINED":            400170,
	"UNSUPPORTED":             400180,
}

func (x StatusIds_StatusCode) String() string {
	return proto.EnumName(StatusIds_StatusCode_name, int32(x))
}

func (StatusIds_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0, 0}
}

type StatusIds struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusIds) Reset()         { *m = StatusIds{} }
func (m *StatusIds) String() string { return proto.CompactTextString(m) }
func (*StatusIds) ProtoMessage()    {}
func (*StatusIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0}
}

func (m *StatusIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusIds.Unmarshal(m, b)
}
func (m *StatusIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusIds.Marshal(b, m, deterministic)
}
func (m *StatusIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusIds.Merge(m, src)
}
func (m *StatusIds) XXX_Size() int {
	return xxx_messageInfo_StatusIds.Size(m)
}
func (m *StatusIds) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusIds.DiscardUnknown(m)
}

var xxx_messageInfo_StatusIds proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Ydb.StatusIds_StatusCode", StatusIds_StatusCode_name, StatusIds_StatusCode_value)
	proto.RegisterType((*StatusIds)(nil), "Ydb.StatusIds")
}

func init() { proto.RegisterFile("ydb_status_codes.proto", fileDescriptor_f81e45819472f2bf) }

var fileDescriptor_f81e45819472f2bf = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xbf, 0x8b, 0x13, 0x41,
	0x14, 0x07, 0x70, 0x21, 0xa0, 0xdc, 0xd3, 0xbb, 0x7b, 0x37, 0xe7, 0x8f, 0x11, 0xbb, 0xb3, 0x4f,
	0xe3, 0x5f, 0x30, 0x3b, 0xf3, 0xe2, 0x0d, 0x6c, 0x66, 0xd6, 0xf9, 0x11, 0x12, 0x9b, 0x21, 0x9b,
	0x4d, 0xa9, 0x2b, 0x6e, 0x04, 0xd3, 0x89, 0x60, 0x23, 0x76, 0x8a, 0xa5, 0x58, 0x8a, 0xa5, 0xb8,
	0x5b, 0x59, 0x5b, 0x5b, 0x5b, 0xfb, 0x47, 0x58, 0x1e, 0xc9, 0x2e, 0xa4, 0x9c, 0x1f, 0x7c, 0x78,
	0xef, 0xfb, 0x85, 0xbb, 0xdb, 0xaa, 0x4c, 0xcd, 0x66, 0xb9, 0x79, 0xd5, 0xa4, 0x55, 0x5d, 0xad,
	0x9b, 0xf1, 0x8b, 0x97, 0xf5, 0xa6, 0x66, 0xa3, 0x45, 0x55, 0x5e, 0x7c, 0x1c, 0xc1, 0x91, 0xdf,
	0xbf, 0xe9, 0xaa, 0xb9, 0x78, 0x37, 0x02, 0xe8, 0x4f, 0xb2, 0xae, 0xd6, 0xec, 0x01, 0xdc, 0xf3,
	0x41, 0x84, 0xe8, 0x93, 0xb4, 0x8a, 0x52, 0x34, 0xbe, 0x20, 0xa9, 0x27, 0x9a, 0x14, 0x5e, 0x63,
	0xc7, 0x70, 0xc3, 0x47, 0x29, 0xc9, 0x7b, 0x7c, 0xd3, 0x72, 0x76, 0x06, 0x37, 0x33, 0xa1, 0x92,
	0xa3, 0x27, 0x91, 0x7c, 0xc0, 0xf7, 0x2d, 0x67, 0x0c, 0x6e, 0x45, 0x23, 0x62, 0xb8, 0xb4, 0x4e,
	0x3f, 0x25, 0x85, 0x9f, 0x5a, 0xce, 0x6e, 0xc3, 0x89, 0x36, 0x81, 0x9c, 0x11, 0x79, 0x22, 0xe7,
	0xac, 0xc3, 0x2f, 0x2d, 0xdf, 0x59, 0x22, 0xb3, 0x2e, 0x90, 0xc2, 0x6f, 0xbd, 0x15, 0x8d, 0x98,
	0x09, 0x9d, 0x8b, 0x2c, 0x27, 0xfc, 0xd1, 0x72, 0x86, 0x00, 0x76, 0x46, 0x2e, 0xb7, 0x42, 0x91,
	0xc2, 0x5f, 0xbd, 0xee, 0xe5, 0x25, 0x4d, 0x69, 0x70, 0x7e, 0xb7, 0x9c, 0x9d, 0xc3, 0xf1, 0x63,
	0x32, 0xe4, 0xb4, 0x1c, 0x2e, 0xff, 0xf4, 0x78, 0xd0, 0x53, 0xb2, 0x31, 0xe0, 0xdf, 0xc3, 0xa0,
	0x9e, 0xbc, 0xd7, 0xd6, 0xe0, 0xbf, 0x96, 0xb3, 0xfb, 0x70, 0x5e, 0x38, 0x92, 0xd6, 0x28, 0x1d,
	0xb4, 0x35, 0x69, 0x22, 0x74, 0x4e, 0x0a, 0xff, 0xf7, 0xf3, 0x8a, 0xdc, 0x91, 0x50, 0x8b, 0x44,
	0x73, 0xed, 0x83, 0xc7, 0xb7, 0x1d, 0x67, 0xa7, 0x70, 0x64, 0x6c, 0x48, 0x13, 0x1b, 0x8d, 0xc2,
	0x0f, 0x1d, 0x67, 0x77, 0xe0, 0x74, 0x00, 0x13, 0xcd, 0x0b, 0xed, 0x48, 0xe1, 0xe7, 0xfe, 0x9f,
	0x14, 0x46, 0x52, 0xbe, 0xe3, 0xbe, 0x76, 0x43, 0x24, 0x8a, 0x02, 0xb9, 0xa9, 0x36, 0xa4, 0xf0,
	0x7b, 0x37, 0x6c, 0xeb, 0x63, 0x51, 0xf4, 0x01, 0xfc, 0xec, 0x78, 0xf6, 0x10, 0x4e, 0x56, 0xf5,
	0xb3, 0xf1, 0x76, 0xf9, 0xbc, 0x5a, 0xbf, 0x1e, 0x6f, 0xab, 0x32, 0x3b, 0x3b, 0xd4, 0xd2, 0x14,
	0xbb, 0xfa, 0x9a, 0xf2, 0xfa, 0xbe, 0xc6, 0x47, 0x57, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x54,
	0x93, 0xd7, 0xe0, 0x01, 0x00, 0x00,
}

const ()
