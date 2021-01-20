// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1alpha1/ca.proto

// Keep this package for backward compatibility.

package v1alpha1

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Certificate request message. The authentication should be based on:
// 1. Bearer tokens carried in the side channel;
// 2. Client-side certificate via Mutual TLS handshake.
// Note: the service implementation is REQUIRED to verify the authenticated caller is authorize to
// all SANs in the CSR. The server side may overwrite any requested certificate field based on its
// policies.
type IstioCertificateRequest struct {
	// PEM-encoded certificate request.
	// The public key in the CSR is used to generate the certificate,
	// and other fields in the generated certificate may be overwritten by the CA.
	Csr string `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
	// Optional: requested certificate validity period, in seconds.
	ValidityDuration     int64    `protobuf:"varint,3,opt,name=validity_duration,json=validityDuration,proto3" json:"validity_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioCertificateRequest) Reset()         { *m = IstioCertificateRequest{} }
func (m *IstioCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*IstioCertificateRequest) ProtoMessage()    {}
func (*IstioCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af28ba83f79edf69, []int{0}
}
func (m *IstioCertificateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioCertificateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioCertificateRequest.Merge(m, src)
}
func (m *IstioCertificateRequest) XXX_Size() int {
	return m.Size()
}
func (m *IstioCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IstioCertificateRequest proto.InternalMessageInfo

func (m *IstioCertificateRequest) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

func (m *IstioCertificateRequest) GetValidityDuration() int64 {
	if m != nil {
		return m.ValidityDuration
	}
	return 0
}

// Certificate response message.
type IstioCertificateResponse struct {
	// PEM-encoded certificate chain.
	// The leaf cert is the first element, and the root cert is the last element.
	CertChain            []string `protobuf:"bytes,1,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioCertificateResponse) Reset()         { *m = IstioCertificateResponse{} }
func (m *IstioCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*IstioCertificateResponse) ProtoMessage()    {}
func (*IstioCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af28ba83f79edf69, []int{1}
}
func (m *IstioCertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioCertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioCertificateResponse.Merge(m, src)
}
func (m *IstioCertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *IstioCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IstioCertificateResponse proto.InternalMessageInfo

func (m *IstioCertificateResponse) GetCertChain() []string {
	if m != nil {
		return m.CertChain
	}
	return nil
}

func init() {
	proto.RegisterType((*IstioCertificateRequest)(nil), "istio.v1.auth.IstioCertificateRequest")
	proto.RegisterType((*IstioCertificateResponse)(nil), "istio.v1.auth.IstioCertificateResponse")
}

func init() { proto.RegisterFile("security/v1alpha1/ca.proto", fileDescriptor_af28ba83f79edf69) }

var fileDescriptor_af28ba83f79edf69 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x5d, 0x02, 0x42, 0x16, 0x84, 0xbb, 0x6d, 0x0c, 0x07, 0x86, 0x90, 0x42, 0x03, 0xca,
	0x86, 0x68, 0x65, 0x6b, 0x6c, 0x6c, 0x63, 0x23, 0x36, 0xc7, 0xb8, 0x37, 0x47, 0x06, 0x8e, 0x6c,
	0xdc, 0x9d, 0x04, 0xae, 0xf4, 0xdf, 0x59, 0xfa, 0x13, 0x24, 0xbf, 0x44, 0x62, 0x14, 0x94, 0x20,
	0x76, 0xc3, 0xf7, 0x66, 0x86, 0xf7, 0x9e, 0x5c, 0x79, 0x34, 0x9d, 0x23, 0xde, 0xe7, 0x7d, 0x01,
	0xbb, 0xb6, 0x86, 0x22, 0x37, 0xa0, 0x5b, 0x67, 0xd9, 0xaa, 0x23, 0xf2, 0x4c, 0x56, 0xf7, 0x85,
	0x86, 0x8e, 0xeb, 0xf4, 0x41, 0x1e, 0xdf, 0x8d, 0xa0, 0x44, 0xc7, 0xb4, 0x25, 0x03, 0x8c, 0x15,
	0x3e, 0x77, 0xe8, 0x59, 0x2d, 0x64, 0x60, 0xbc, 0x8b, 0x44, 0x22, 0xb2, 0xb0, 0x1a, 0x47, 0x75,
	0x2e, 0x97, 0x3d, 0xec, 0x68, 0x43, 0xbc, 0x5f, 0x6f, 0x3a, 0x07, 0x4c, 0xb6, 0x89, 0x82, 0x44,
	0x64, 0x41, 0xb5, 0xf8, 0x16, 0x6e, 0xbf, 0x78, 0x7a, 0x2d, 0xa3, 0xf9, 0x67, 0xdf, 0xda, 0xc6,
	0xa3, 0x3a, 0x91, 0xd2, 0xa0, 0xe3, 0xb5, 0xa9, 0x81, 0x9a, 0x48, 0x24, 0x41, 0x16, 0x56, 0xe1,
	0x48, 0xca, 0x11, 0x5c, 0xbe, 0x88, 0xb9, 0xab, 0x7b, 0x74, 0x3d, 0x19, 0x54, 0x5b, 0xb9, 0x2c,
	0x1d, 0x02, 0xe3, 0x0f, 0x4d, 0x9d, 0xea, 0x5f, 0xa9, 0xf4, 0x1f, 0x91, 0x56, 0x67, 0xff, 0xee,
	0x4d, 0x06, 0xd3, 0x83, 0x9b, 0x8b, 0xd7, 0x21, 0x16, 0x6f, 0x43, 0x2c, 0xde, 0x87, 0x58, 0x3c,
	0xc6, 0xd3, 0x1d, 0xd9, 0x1c, 0x5a, 0xca, 0x67, 0xf5, 0x3e, 0x1d, 0x7e, 0x96, 0x7b, 0xf5, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x67, 0x4f, 0xd9, 0x1c, 0x7a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IstioCertificateServiceClient is the client API for IstioCertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IstioCertificateServiceClient interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(ctx context.Context, in *IstioCertificateRequest, opts ...grpc.CallOption) (*IstioCertificateResponse, error)
}

type istioCertificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewIstioCertificateServiceClient(cc *grpc.ClientConn) IstioCertificateServiceClient {
	return &istioCertificateServiceClient{cc}
}

func (c *istioCertificateServiceClient) CreateCertificate(ctx context.Context, in *IstioCertificateRequest, opts ...grpc.CallOption) (*IstioCertificateResponse, error) {
	out := new(IstioCertificateResponse)
	err := c.cc.Invoke(ctx, "/istio.v1.auth.IstioCertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IstioCertificateServiceServer is the server API for IstioCertificateService service.
type IstioCertificateServiceServer interface {
	// Using provided CSR, returns a signed certificate.
	CreateCertificate(context.Context, *IstioCertificateRequest) (*IstioCertificateResponse, error)
}

// UnimplementedIstioCertificateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIstioCertificateServiceServer struct {
}

func (*UnimplementedIstioCertificateServiceServer) CreateCertificate(ctx context.Context, req *IstioCertificateRequest) (*IstioCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}

func RegisterIstioCertificateServiceServer(s *grpc.Server, srv IstioCertificateServiceServer) {
	s.RegisterService(&_IstioCertificateService_serviceDesc, srv)
}

func _IstioCertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IstioCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IstioCertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/istio.v1.auth.IstioCertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IstioCertificateServiceServer).CreateCertificate(ctx, req.(*IstioCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IstioCertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "istio.v1.auth.IstioCertificateService",
	HandlerType: (*IstioCertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _IstioCertificateService_CreateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/v1alpha1/ca.proto",
}

func (m *IstioCertificateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioCertificateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioCertificateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ValidityDuration != 0 {
		i = encodeVarintCa(dAtA, i, uint64(m.ValidityDuration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Csr) > 0 {
		i -= len(m.Csr)
		copy(dAtA[i:], m.Csr)
		i = encodeVarintCa(dAtA, i, uint64(len(m.Csr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IstioCertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioCertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioCertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CertChain) > 0 {
		for iNdEx := len(m.CertChain) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CertChain[iNdEx])
			copy(dAtA[i:], m.CertChain[iNdEx])
			i = encodeVarintCa(dAtA, i, uint64(len(m.CertChain[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCa(dAtA []byte, offset int, v uint64) int {
	offset -= sovCa(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioCertificateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Csr)
	if l > 0 {
		n += 1 + l + sovCa(uint64(l))
	}
	if m.ValidityDuration != 0 {
		n += 1 + sovCa(uint64(m.ValidityDuration))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioCertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CertChain) > 0 {
		for _, s := range m.CertChain {
			l = len(s)
			n += 1 + l + sovCa(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCa(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCa(x uint64) (n int) {
	return sovCa(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioCertificateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
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
			return fmt.Errorf("proto: IstioCertificateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioCertificateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
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
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCa
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidityDuration", wireType)
			}
			m.ValidityDuration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValidityDuration |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCa(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCa
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
func (m *IstioCertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCa
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
			return fmt.Errorf("proto: IstioCertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioCertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCa
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
				return ErrInvalidLengthCa
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCa
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertChain = append(m.CertChain, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCa(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCa
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCa
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
func skipCa(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCa
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
					return 0, ErrIntOverflowCa
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
					return 0, ErrIntOverflowCa
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
				return 0, ErrInvalidLengthCa
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCa
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCa
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
				next, err := skipCa(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCa
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
	ErrInvalidLengthCa = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCa   = fmt.Errorf("proto: integer overflow")
)
