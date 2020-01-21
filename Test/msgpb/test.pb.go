// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package msgpb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//*
// @brief base_score_info
type Test struct {
	WinCount             *int32   `protobuf:"varint,1,opt,name=win_count,json=winCount" json:"win_count,omitempty"`
	LoseCount            *int32   `protobuf:"varint,2,opt,name=lose_count,json=loseCount" json:"lose_count,omitempty"`
	ExceptionCount       *int32   `protobuf:"varint,3,opt,name=exception_count,json=exceptionCount" json:"exception_count,omitempty"`
	KillCount            *int32   `protobuf:"varint,4,opt,name=kill_count,json=killCount" json:"kill_count,omitempty"`
	DeathCount           *int32   `protobuf:"varint,5,opt,name=death_count,json=deathCount" json:"death_count,omitempty"`
	AssistCount          *int32   `protobuf:"varint,6,opt,name=assist_count,json=assistCount" json:"assist_count,omitempty"`
	Rating               *int64   `protobuf:"varint,7,opt,name=rating" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetWinCount() int32 {
	if m != nil && m.WinCount != nil {
		return *m.WinCount
	}
	return 0
}

func (m *Test) GetLoseCount() int32 {
	if m != nil && m.LoseCount != nil {
		return *m.LoseCount
	}
	return 0
}

func (m *Test) GetExceptionCount() int32 {
	if m != nil && m.ExceptionCount != nil {
		return *m.ExceptionCount
	}
	return 0
}

func (m *Test) GetKillCount() int32 {
	if m != nil && m.KillCount != nil {
		return *m.KillCount
	}
	return 0
}

func (m *Test) GetDeathCount() int32 {
	if m != nil && m.DeathCount != nil {
		return *m.DeathCount
	}
	return 0
}

func (m *Test) GetAssistCount() int32 {
	if m != nil && m.AssistCount != nil {
		return *m.AssistCount
	}
	return 0
}

func (m *Test) GetRating() int64 {
	if m != nil && m.Rating != nil {
		return *m.Rating
	}
	return 0
}

func init() {
	proto.RegisterType((*Test)(nil), "msgpb.test")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8c, 0x5d, 0x0a, 0x82, 0x40,
	0x14, 0x46, 0x99, 0xfc, 0x29, 0xaf, 0x51, 0xe0, 0x43, 0x08, 0x11, 0x59, 0x2f, 0xf9, 0xd4, 0x26,
	0xda, 0x81, 0x1b, 0x08, 0xb3, 0xc1, 0x86, 0x6c, 0x46, 0xbc, 0x37, 0x6c, 0xc5, 0xad, 0x23, 0x66,
	0xee, 0xe4, 0xe3, 0x77, 0xce, 0xe1, 0x03, 0x20, 0x89, 0x74, 0xee, 0x07, 0x43, 0x26, 0x8b, 0x5e,
	0xd8, 0xf6, 0xb7, 0xe3, 0x57, 0x40, 0x68, 0x69, 0xb6, 0x85, 0x64, 0x54, 0xfa, 0xda, 0x98, 0xb7,
	0xa6, 0x5c, 0x14, 0xa2, 0x8c, 0xaa, 0xc5, 0xa8, 0xf4, 0xc5, 0xee, 0x6c, 0x07, 0xd0, 0x19, 0x94,
	0xde, 0xce, 0x9c, 0x4d, 0x2c, 0x61, 0x7d, 0x82, 0xb5, 0xfc, 0x34, 0xb2, 0x27, 0x65, 0xfe, 0x0f,
	0x81, 0x6b, 0x56, 0x13, 0x9e, 0x7e, 0x9e, 0xaa, 0xeb, 0x7c, 0x13, 0xf2, 0x8f, 0x25, 0xac, 0xf7,
	0x90, 0xde, 0x65, 0x4d, 0x0f, 0xef, 0x23, 0xe7, 0xc1, 0x21, 0x0e, 0x0e, 0xb0, 0xac, 0x11, 0x15,
	0x92, 0x2f, 0x62, 0x57, 0xa4, 0xcc, 0x38, 0xd9, 0x40, 0x3c, 0xd4, 0xa4, 0x74, 0x9b, 0xcf, 0x0b,
	0x51, 0x06, 0x95, 0x5f, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0xa0, 0x89, 0x59, 0xfc, 0x00,
	0x00, 0x00,
}
