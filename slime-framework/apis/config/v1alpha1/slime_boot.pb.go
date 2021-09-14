// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apis/config/v1alpha1/slime_boot.proto

package v1alpha1

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SlimeBootStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlimeBootStatus) Reset()         { *m = SlimeBootStatus{} }
func (m *SlimeBootStatus) String() string { return proto.CompactTextString(m) }
func (*SlimeBootStatus) ProtoMessage()    {}
func (*SlimeBootStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0354666a6f23fc96, []int{0}
}

func (m *SlimeBootStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlimeBootStatus.Unmarshal(m, b)
}

func (m *SlimeBootStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlimeBootStatus.Marshal(b, m, deterministic)
}

func (m *SlimeBootStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlimeBootStatus.Merge(m, src)
}

func (m *SlimeBootStatus) XXX_Size() int {
	return xxx_messageInfo_SlimeBootStatus.Size(m)
}

func (m *SlimeBootStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SlimeBootStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SlimeBootStatus proto.InternalMessageInfo

type SlimeBootSpec struct {
	Module               []*Config  `protobuf:"bytes,1,rep,name=module,proto3" json:"module,omitempty"`
	Component            *Component `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SlimeBootSpec) Reset()         { *m = SlimeBootSpec{} }
func (m *SlimeBootSpec) String() string { return proto.CompactTextString(m) }
func (*SlimeBootSpec) ProtoMessage()    {}
func (*SlimeBootSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0354666a6f23fc96, []int{1}
}

func (m *SlimeBootSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlimeBootSpec.Unmarshal(m, b)
}

func (m *SlimeBootSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlimeBootSpec.Marshal(b, m, deterministic)
}

func (m *SlimeBootSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlimeBootSpec.Merge(m, src)
}

func (m *SlimeBootSpec) XXX_Size() int {
	return xxx_messageInfo_SlimeBootSpec.Size(m)
}

func (m *SlimeBootSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SlimeBootSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SlimeBootSpec proto.InternalMessageInfo

func (m *SlimeBootSpec) GetModule() []*Config {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *SlimeBootSpec) GetComponent() *Component {
	if m != nil {
		return m.Component
	}
	return nil
}

type Component struct {
	GlobalSidecarNamespace []string `protobuf:"bytes,1,rep,name=globalSidecarNamespace,proto3" json:"globalSidecarNamespace,omitempty"`
	ReportServer           bool     `protobuf:"varint,2,opt,name=reportServer,proto3" json:"reportServer,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Component) Reset()         { *m = Component{} }
func (m *Component) String() string { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()    {}
func (*Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_0354666a6f23fc96, []int{2}
}

func (m *Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Component.Unmarshal(m, b)
}

func (m *Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Component.Marshal(b, m, deterministic)
}

func (m *Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Component.Merge(m, src)
}

func (m *Component) XXX_Size() int {
	return xxx_messageInfo_Component.Size(m)
}

func (m *Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Component proto.InternalMessageInfo

func (m *Component) GetGlobalSidecarNamespace() []string {
	if m != nil {
		return m.GlobalSidecarNamespace
	}
	return nil
}

func (m *Component) GetReportServer() bool {
	if m != nil {
		return m.ReportServer
	}
	return false
}

func init() {
	proto.RegisterType((*SlimeBootStatus)(nil), "netease.config.v1alpha1.SlimeBootStatus")
	proto.RegisterType((*SlimeBootSpec)(nil), "netease.config.v1alpha1.SlimeBootSpec")
	proto.RegisterType((*Component)(nil), "netease.config.v1alpha1.Component")
}

func init() {
	proto.RegisterFile("pkg/apis/config/v1alpha1/slime_boot.proto", fileDescriptor_0354666a6f23fc96)
}

var fileDescriptor_0354666a6f23fc96 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x14, 0x84, 0xa9, 0xc2, 0x62, 0xb3, 0x8a, 0x98, 0x83, 0x2e, 0x5e, 0x5c, 0x02, 0xc2, 0x7a, 0x49,
	0xdc, 0x15, 0xf4, 0x2a, 0xeb, 0xdd, 0x43, 0x7b, 0xf3, 0x22, 0xaf, 0xd9, 0x67, 0x2d, 0x26, 0x79,
	0x21, 0x49, 0x17, 0xfc, 0x0b, 0xfe, 0x6a, 0xb1, 0xad, 0xae, 0x1e, 0xea, 0x31, 0x93, 0x6f, 0xe6,
	0x0d, 0xc3, 0xae, 0xfc, 0x5b, 0xad, 0xc0, 0x37, 0x51, 0x69, 0x72, 0x2f, 0x4d, 0xad, 0xb6, 0x4b,
	0x30, 0xfe, 0x15, 0x96, 0x2a, 0x9a, 0xc6, 0xe2, 0x73, 0x45, 0x94, 0xa4, 0x0f, 0x94, 0x88, 0x9f,
	0x39, 0x4c, 0x08, 0x11, 0x65, 0x4f, 0xca, 0x6f, 0xf2, 0xfc, 0x72, 0x34, 0x63, 0x20, 0x3b, 0xbf,
	0x38, 0x61, 0xc7, 0xe5, 0x57, 0xe6, 0x9a, 0x28, 0x95, 0x09, 0x52, 0x1b, 0xc5, 0x47, 0xc6, 0x8e,
	0x76, 0x9a, 0x47, 0xcd, 0xef, 0xd8, 0xc4, 0xd2, 0xa6, 0x35, 0x38, 0xcb, 0xe6, 0xfb, 0x8b, 0xe9,
	0xea, 0x42, 0x8e, 0x5c, 0x95, 0x0f, 0xdd, 0xbb, 0x18, 0x70, 0x7e, 0xcf, 0x72, 0x4d, 0xd6, 0x93,
	0x43, 0x97, 0x66, 0x7b, 0xf3, 0x6c, 0x31, 0x5d, 0x89, 0x7f, 0xbc, 0x03, 0x59, 0xec, 0x4c, 0xa2,
	0x66, 0xf9, 0x8f, 0xce, 0x6f, 0xd9, 0x69, 0x6d, 0xa8, 0x02, 0x53, 0x36, 0x1b, 0xd4, 0x10, 0x1e,
	0xc1, 0x62, 0xf4, 0xa0, 0xfb, 0x5e, 0x79, 0x31, 0xf2, 0xcb, 0x05, 0x3b, 0x0c, 0xe8, 0x29, 0xa4,
	0x12, 0xc3, 0x16, 0x43, 0xd7, 0xe4, 0xa0, 0xf8, 0xa3, 0xad, 0xaf, 0x9f, 0xe4, 0x7b, 0xeb, 0x7e,
	0x95, 0xb3, 0xfd, 0xd8, 0x6a, 0x6c, 0xc7, 0x6a, 0xd2, 0x2d, 0x78, 0xf3, 0x19, 0x00, 0x00, 0xff,
	0xff, 0xa3, 0x8a, 0x5e, 0x56, 0xae, 0x01, 0x00, 0x00,
}