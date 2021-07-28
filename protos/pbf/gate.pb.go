// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gate.proto

package pbf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type DefaultArg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefaultArg) Reset()         { *m = DefaultArg{} }
func (m *DefaultArg) String() string { return proto.CompactTextString(m) }
func (*DefaultArg) ProtoMessage()    {}
func (*DefaultArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}
func (m *DefaultArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultArg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultArg.Merge(m, src)
}
func (m *DefaultArg) XXX_Size() int {
	return m.Size()
}
func (m *DefaultArg) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultArg.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultArg proto.InternalMessageInfo

type GateInfo struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=Host,proto3" json:"Host,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GateInfo) Reset()         { *m = GateInfo{} }
func (m *GateInfo) String() string { return proto.CompactTextString(m) }
func (*GateInfo) ProtoMessage()    {}
func (*GateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}
func (m *GateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GateInfo.Merge(m, src)
}
func (m *GateInfo) XXX_Size() int {
	return m.Size()
}
func (m *GateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GateInfo proto.InternalMessageInfo

func (m *GateInfo) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *GateInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GateInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*DefaultArg)(nil), "pbf.DefaultArg")
	proto.RegisterType((*GateInfo)(nil), "pbf.GateInfo")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x48, 0x4a, 0x53, 0xe2, 0xe1, 0xe2, 0x72,
	0x49, 0x4d, 0x4b, 0x2c, 0xcd, 0x29, 0x71, 0x2c, 0x4a, 0x57, 0xf2, 0xe2, 0xe2, 0x70, 0x4f, 0x2c,
	0x49, 0xf5, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe3, 0x62, 0x0b, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x84, 0xb8, 0x58, 0x3c, 0xf2, 0x8b, 0x4b,
	0x24, 0x98, 0xc0, 0xa2, 0x60, 0x36, 0x48, 0x2c, 0x20, 0xb1, 0x24, 0x43, 0x82, 0x19, 0x22, 0x06,
	0x62, 0x1b, 0x99, 0x73, 0xb1, 0x80, 0xcc, 0x12, 0xd2, 0xe7, 0xe2, 0x76, 0x4f, 0x2d, 0x81, 0x1b,
	0xcb, 0xaf, 0x57, 0x90, 0x94, 0xa6, 0x87, 0xb0, 0x53, 0x8a, 0x17, 0x2c, 0x00, 0x93, 0x57, 0x62,
	0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67,
	0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x3b, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xc6,
	0x19, 0xa2, 0xbe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateClient is the client API for Gate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateClient interface {
	GetGateInfo(ctx context.Context, in *DefaultArg, opts ...grpc.CallOption) (*GateInfo, error)
}

type gateClient struct {
	cc *grpc.ClientConn
}

func NewGateClient(cc *grpc.ClientConn) GateClient {
	return &gateClient{cc}
}

func (c *gateClient) GetGateInfo(ctx context.Context, in *DefaultArg, opts ...grpc.CallOption) (*GateInfo, error) {
	out := new(GateInfo)
	err := c.cc.Invoke(ctx, "/pbf.Gate/GetGateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServer is the server API for Gate service.
type GateServer interface {
	GetGateInfo(context.Context, *DefaultArg) (*GateInfo, error)
}

// UnimplementedGateServer can be embedded to have forward compatible implementations.
type UnimplementedGateServer struct {
}

func (*UnimplementedGateServer) GetGateInfo(ctx context.Context, req *DefaultArg) (*GateInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGateInfo not implemented")
}

func RegisterGateServer(s *grpc.Server, srv GateServer) {
	s.RegisterService(&_Gate_serviceDesc, srv)
}

func _Gate_GetGateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).GetGateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbf.Gate/GetGateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).GetGateInfo(ctx, req.(*DefaultArg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbf.Gate",
	HandlerType: (*GateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGateInfo",
			Handler:    _Gate_GetGateInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gate.proto",
}

func (m *DefaultArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultArg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultArg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *GateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GateInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GateInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Scheme) > 0 {
		i -= len(m.Scheme)
		copy(dAtA[i:], m.Scheme)
		i = encodeVarintGate(dAtA, i, uint64(len(m.Scheme)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGate(dAtA []byte, offset int, v uint64) int {
	offset -= sovGate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DefaultArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Scheme)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovGate(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGate(x uint64) (n int) {
	return sovGate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DefaultArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGate
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
			return fmt.Errorf("proto: DefaultArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGate
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
func (m *GateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGate
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
			return fmt.Errorf("proto: GateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scheme", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scheme = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGate
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
				return ErrInvalidLengthGate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGate
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
func skipGate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGate
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
					return 0, ErrIntOverflowGate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGate
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
				return 0, ErrInvalidLengthGate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGate = fmt.Errorf("proto: unexpected end of group")
)