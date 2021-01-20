// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta/v1alpha1/status.proto

package v1alpha1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1alpha1 "istio.io/api/analysis/v1alpha1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type IstioStatus struct {
	// Current service state of pod.
	// More info: https://istio.io/docs/reference/config/config-status/
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*IstioCondition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Includes any errors or warnings detected by Istio's analyzers.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	ValidationMessages   []*v1alpha1.AnalysisMessageBase `protobuf:"bytes,2,rep,name=validation_messages,json=validationMessages,proto3" json:"validation_messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IstioStatus) Reset()         { *m = IstioStatus{} }
func (m *IstioStatus) String() string { return proto.CompactTextString(m) }
func (*IstioStatus) ProtoMessage()    {}
func (*IstioStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c58e523a81edc6, []int{0}
}
func (m *IstioStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioStatus.Merge(m, src)
}
func (m *IstioStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioStatus proto.InternalMessageInfo

func (m *IstioStatus) GetConditions() []*IstioCondition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *IstioStatus) GetValidationMessages() []*v1alpha1.AnalysisMessageBase {
	if m != nil {
		return m.ValidationMessages
	}
	return nil
}

type IstioCondition struct {
	// Type is the type of the condition.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime *types.Timestamp `protobuf:"bytes,3,opt,name=last_probe_time,json=lastProbeTime,proto3" json:"last_probe_time,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *types.Timestamp `protobuf:"bytes,4,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioCondition) Reset()         { *m = IstioCondition{} }
func (m *IstioCondition) String() string { return proto.CompactTextString(m) }
func (*IstioCondition) ProtoMessage()    {}
func (*IstioCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c58e523a81edc6, []int{1}
}
func (m *IstioCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioCondition.Merge(m, src)
}
func (m *IstioCondition) XXX_Size() int {
	return m.Size()
}
func (m *IstioCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioCondition.DiscardUnknown(m)
}

var xxx_messageInfo_IstioCondition proto.InternalMessageInfo

func (m *IstioCondition) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *IstioCondition) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *IstioCondition) GetLastProbeTime() *types.Timestamp {
	if m != nil {
		return m.LastProbeTime
	}
	return nil
}

func (m *IstioCondition) GetLastTransitionTime() *types.Timestamp {
	if m != nil {
		return m.LastTransitionTime
	}
	return nil
}

func (m *IstioCondition) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *IstioCondition) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IstioStatus)(nil), "istio.meta.v1alpha1.IstioStatus")
	proto.RegisterType((*IstioCondition)(nil), "istio.meta.v1alpha1.IstioCondition")
}

func init() { proto.RegisterFile("meta/v1alpha1/status.proto", fileDescriptor_05c58e523a81edc6) }

var fileDescriptor_05c58e523a81edc6 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xc9, 0x36, 0x2b, 0x66, 0xa8, 0x90, 0xc9, 0x08, 0x3d, 0x6c, 0x63, 0x5e, 0x7a, 0x90,
	0x94, 0xcd, 0x27, 0x70, 0x3b, 0x09, 0x0a, 0x52, 0x77, 0x12, 0xa4, 0xfc, 0xeb, 0xb2, 0x2d, 0xd0,
	0x36, 0xa5, 0xc9, 0x06, 0xbb, 0xf9, 0x4a, 0xbe, 0x85, 0x47, 0x1f, 0x41, 0xf6, 0x24, 0x92, 0xa4,
	0x75, 0x4e, 0x04, 0x6f, 0xfd, 0xff, 0xfb, 0xfd, 0xbe, 0x7c, 0xfd, 0x1a, 0xec, 0x67, 0x5c, 0x43,
	0xb8, 0x19, 0x41, 0x5a, 0xac, 0x60, 0x14, 0x2a, 0x0d, 0x7a, 0xad, 0x58, 0x51, 0x4a, 0x2d, 0x49,
	0x47, 0x28, 0x2d, 0x24, 0x33, 0x0a, 0x56, 0x2b, 0xfc, 0x3e, 0xe4, 0x90, 0x6e, 0x95, 0x50, 0x7b,
	0x28, 0xe3, 0x4a, 0xc1, 0x92, 0x3b, 0xca, 0xef, 0x2f, 0xa5, 0x5c, 0xa6, 0x3c, 0xb4, 0x53, 0xb2,
	0x5e, 0x84, 0x5a, 0x64, 0x5c, 0x69, 0xc8, 0x8a, 0x5f, 0x02, 0x28, 0x44, 0xb8, 0x10, 0x3c, 0x9d,
	0xc7, 0x09, 0x5f, 0xc1, 0x46, 0xc8, 0xd2, 0x09, 0x86, 0x6f, 0x08, 0xb7, 0x6f, 0xcd, 0xd1, 0x8f,
	0x36, 0x0d, 0x99, 0x62, 0xfc, 0x22, 0xf3, 0xb9, 0xd0, 0x42, 0xe6, 0x8a, 0xa2, 0x41, 0x33, 0x68,
	0x8f, 0x2f, 0xd9, 0x1f, 0xe1, 0x98, 0xa5, 0xa6, 0xb5, 0x36, 0xfa, 0x81, 0x91, 0x67, 0xdc, 0xd9,
	0x40, 0x2a, 0xe6, 0x60, 0xc6, 0xb8, 0x8a, 0xac, 0x68, 0xc3, 0xba, 0x5d, 0x55, 0x6e, 0xf5, 0xb7,
	0xed, 0x1d, 0x6f, 0xaa, 0xcd, 0xbd, 0x03, 0x26, 0xa0, 0x78, 0x44, 0xf6, 0x46, 0xd5, 0x5a, 0x0d,
	0x5f, 0x1b, 0xf8, 0xec, 0xf0, 0x74, 0x42, 0x70, 0x4b, 0x6f, 0x0b, 0x4e, 0xd1, 0x00, 0x05, 0x27,
	0x91, 0x7d, 0x26, 0x5d, 0xec, 0xb9, 0x8a, 0x69, 0xc3, 0x6e, 0xab, 0x89, 0x4c, 0xf0, 0x79, 0x0a,
	0x4a, 0xc7, 0x45, 0x29, 0x13, 0x1e, 0x9b, 0xc6, 0x68, 0x73, 0x80, 0x82, 0xf6, 0xd8, 0x67, 0xae,
	0x2d, 0x56, 0xd7, 0xc9, 0x66, 0x75, 0x9d, 0xd1, 0xa9, 0x41, 0x1e, 0x0c, 0x61, 0x76, 0xe4, 0x0e,
	0x5f, 0x58, 0x0f, 0x5d, 0x42, 0xae, 0x6c, 0x04, 0x67, 0xd4, 0xfa, 0xd7, 0x88, 0x18, 0x6e, 0xf6,
	0x8d, 0x59, 0xb7, 0x2e, 0xf6, 0x4a, 0x0e, 0x4a, 0xe6, 0xf4, 0xc8, 0x25, 0x75, 0x13, 0xa1, 0xf8,
	0xb8, 0x2a, 0x8f, 0x7a, 0xf6, 0x45, 0x3d, 0x4e, 0x82, 0xf7, 0x5d, 0x0f, 0x7d, 0xec, 0x7a, 0xe8,
	0x73, 0xd7, 0x43, 0x4f, 0xbe, 0x6b, 0x54, 0x48, 0xfb, 0x9f, 0x0f, 0x6e, 0x59, 0xe2, 0xd9, 0x0c,
	0xd7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xfa, 0xdc, 0x0a, 0x7d, 0x02, 0x00, 0x00,
}

func (m *IstioStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ValidationMessages) > 0 {
		for iNdEx := len(m.ValidationMessages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidationMessages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStatus(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStatus(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *IstioCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x2a
	}
	if m.LastTransitionTime != nil {
		{
			size, err := m.LastTransitionTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.LastProbeTime != nil {
		{
			size, err := m.LastProbeTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStatus(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	offset -= sovStatus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if len(m.ValidationMessages) > 0 {
		for _, e := range m.ValidationMessages {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	if m.LastProbeTime != nil {
		l = m.LastProbeTime.Size()
		n += 1 + l + sovStatus(uint64(l))
	}
	if m.LastTransitionTime != nil {
		l = m.LastTransitionTime.Size()
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IstioStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &IstioCondition{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidationMessages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidationMessages = append(m.ValidationMessages, &v1alpha1.AnalysisMessageBase{})
			if err := m.ValidationMessages[len(m.ValidationMessages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IstioCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IstioCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastProbeTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastProbeTime == nil {
				m.LastProbeTime = &types.Timestamp{}
			}
			if err := m.LastProbeTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastTransitionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastTransitionTime == nil {
				m.LastTransitionTime = &types.Timestamp{}
			}
			if err := m.LastTransitionTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStatus
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStatus
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)
