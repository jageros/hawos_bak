// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router.proto

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

type ReqArg struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	MsgID                int32    `protobuf:"varint,2,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqArg) Reset()         { *m = ReqArg{} }
func (m *ReqArg) String() string { return proto.CompactTextString(m) }
func (*ReqArg) ProtoMessage()    {}
func (*ReqArg) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}
func (m *ReqArg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReqArg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReqArg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReqArg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqArg.Merge(m, src)
}
func (m *ReqArg) XXX_Size() int {
	return m.Size()
}
func (m *ReqArg) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqArg.DiscardUnknown(m)
}

var xxx_messageInfo_ReqArg proto.InternalMessageInfo

func (m *ReqArg) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqArg) GetMsgID() int32 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *ReqArg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type RespMsg struct {
	MsgID                int32    `protobuf:"varint,1,opt,name=MsgID,proto3" json:"MsgID,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespMsg) Reset()         { *m = RespMsg{} }
func (m *RespMsg) String() string { return proto.CompactTextString(m) }
func (*RespMsg) ProtoMessage()    {}
func (*RespMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}
func (m *RespMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RespMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RespMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RespMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespMsg.Merge(m, src)
}
func (m *RespMsg) XXX_Size() int {
	return m.Size()
}
func (m *RespMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RespMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RespMsg proto.InternalMessageInfo

func (m *RespMsg) GetMsgID() int32 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *RespMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type QueueMsg struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Targets              *Target  `protobuf:"bytes,2,opt,name=Targets,proto3" json:"Targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueMsg) Reset()         { *m = QueueMsg{} }
func (m *QueueMsg) String() string { return proto.CompactTextString(m) }
func (*QueueMsg) ProtoMessage()    {}
func (*QueueMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}
func (m *QueueMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueueMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueueMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueueMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueMsg.Merge(m, src)
}
func (m *QueueMsg) XXX_Size() int {
	return m.Size()
}
func (m *QueueMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueMsg.DiscardUnknown(m)
}

var xxx_messageInfo_QueueMsg proto.InternalMessageInfo

func (m *QueueMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueueMsg) GetTargets() *Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

type Target struct {
	GroupId              string   `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	Uids                 []string `protobuf:"bytes,2,rep,name=Uids,proto3" json:"Uids,omitempty"`
	UnlessUids           []string `protobuf:"bytes,3,rep,name=UnlessUids,proto3" json:"UnlessUids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}
func (m *Target) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Target.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return m.Size()
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Target) GetUids() []string {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *Target) GetUnlessUids() []string {
	if m != nil {
		return m.UnlessUids
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqArg)(nil), "pbf.ReqArg")
	proto.RegisterType((*RespMsg)(nil), "pbf.RespMsg")
	proto.RegisterType((*QueueMsg)(nil), "pbf.QueueMsg")
	proto.RegisterType((*Target)(nil), "pbf.Target")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0x7f, 0xda, 0xe4, 0xef, 0x6d, 0x16, 0xe5, 0xe2, 0x22, 0xb8, 0x08, 0x21, 0xa0,
	0x64, 0x15, 0xa4, 0x3e, 0x81, 0xb6, 0x22, 0x15, 0x02, 0x3a, 0x18, 0xf7, 0x13, 0x32, 0x0d, 0x85,
	0xe0, 0xa4, 0x33, 0xc9, 0xc2, 0x37, 0xf1, 0x91, 0x5c, 0xfa, 0x08, 0x12, 0x5f, 0x44, 0xe6, 0xa6,
	0xc1, 0xae, 0xdc, 0x9d, 0x73, 0x0f, 0xf7, 0x4b, 0xee, 0x19, 0xf0, 0xb5, 0xea, 0x5a, 0xa9, 0xd3,
	0x46, 0xab, 0x56, 0xa1, 0xd3, 0x14, 0xbb, 0xf8, 0x01, 0x5c, 0x2e, 0x0f, 0x37, 0xba, 0xc2, 0x25,
	0x38, 0xf9, 0xbe, 0x0c, 0x58, 0xc4, 0x92, 0x39, 0xb7, 0x12, 0xcf, 0x60, 0x96, 0x99, 0x6a, 0xbb,
	0x09, 0xfe, 0x45, 0x2c, 0x99, 0xf1, 0xc1, 0x60, 0x00, 0xde, 0xa3, 0x78, 0xab, 0x95, 0x28, 0x03,
	0x27, 0x62, 0x89, 0xcf, 0x47, 0x1b, 0x67, 0xe0, 0x71, 0x69, 0x9a, 0xcc, 0x54, 0xbf, 0xab, 0xec,
	0x74, 0x15, 0x61, 0xba, 0x56, 0xa5, 0x3c, 0xf2, 0x48, 0xff, 0x81, 0xbb, 0x83, 0xff, 0x4f, 0x9d,
	0xec, 0xa4, 0xe5, 0x21, 0x4c, 0x37, 0xa2, 0x15, 0x84, 0xf3, 0x39, 0x69, 0xbc, 0x00, 0xef, 0x59,
	0xe8, 0x4a, 0xb6, 0x86, 0x80, 0x8b, 0xd5, 0x22, 0x6d, 0x8a, 0x5d, 0x3a, 0xcc, 0xf8, 0x98, 0xc5,
	0x2f, 0xe0, 0x0e, 0xd2, 0x7e, 0xea, 0x5e, 0xab, 0xae, 0xd9, 0x8e, 0x57, 0x8e, 0xd6, 0xe2, 0xf3,
	0x7d, 0x69, 0x39, 0x4e, 0x32, 0xe7, 0xa4, 0x31, 0x04, 0xc8, 0x5f, 0x6b, 0x69, 0x0c, 0x25, 0x0e,
	0x25, 0x27, 0x93, 0xd5, 0x15, 0xb8, 0x9c, 0xea, 0xc4, 0x4b, 0x7b, 0xf7, 0x61, 0x2d, 0xea, 0x1a,
	0x87, 0x5f, 0x18, 0x1a, 0x3d, 0xf7, 0x8f, 0x86, 0x2a, 0x89, 0x27, 0xb7, 0xcb, 0x8f, 0x3e, 0x64,
	0x9f, 0x7d, 0xc8, 0xbe, 0xfa, 0x90, 0xbd, 0x7f, 0x87, 0x93, 0xc2, 0xa5, 0x97, 0xb8, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x39, 0xfc, 0x93, 0x60, 0x99, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	ReqCall(ctx context.Context, in *ReqArg, opts ...grpc.CallOption) (*RespMsg, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) ReqCall(ctx context.Context, in *ReqArg, opts ...grpc.CallOption) (*RespMsg, error) {
	out := new(RespMsg)
	err := c.cc.Invoke(ctx, "/pbf.Router/ReqCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	ReqCall(context.Context, *ReqArg) (*RespMsg, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) ReqCall(ctx context.Context, req *ReqArg) (*RespMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqCall not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_ReqCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqArg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ReqCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbf.Router/ReqCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ReqCall(ctx, req.(*ReqArg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbf.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReqCall",
			Handler:    _Router_ReqCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}

func (m *ReqArg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReqArg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReqArg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintRouter(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MsgID != 0 {
		i = encodeVarintRouter(dAtA, i, uint64(m.MsgID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintRouter(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RespMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RespMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RespMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintRouter(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Code != 0 {
		i = encodeVarintRouter(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x10
	}
	if m.MsgID != 0 {
		i = encodeVarintRouter(dAtA, i, uint64(m.MsgID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueueMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueueMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueueMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Targets != nil {
		{
			size, err := m.Targets.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRouter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintRouter(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Target) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Target) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Target) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UnlessUids) > 0 {
		for iNdEx := len(m.UnlessUids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UnlessUids[iNdEx])
			copy(dAtA[i:], m.UnlessUids[iNdEx])
			i = encodeVarintRouter(dAtA, i, uint64(len(m.UnlessUids[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Uids) > 0 {
		for iNdEx := len(m.Uids) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Uids[iNdEx])
			copy(dAtA[i:], m.Uids[iNdEx])
			i = encodeVarintRouter(dAtA, i, uint64(len(m.Uids[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.GroupId) > 0 {
		i -= len(m.GroupId)
		copy(dAtA[i:], m.GroupId)
		i = encodeVarintRouter(dAtA, i, uint64(len(m.GroupId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRouter(dAtA []byte, offset int, v uint64) int {
	offset -= sovRouter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReqArg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.MsgID != 0 {
		n += 1 + sovRouter(uint64(m.MsgID))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RespMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgID != 0 {
		n += 1 + sovRouter(uint64(m.MsgID))
	}
	if m.Code != 0 {
		n += 1 + sovRouter(uint64(m.Code))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueueMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.Targets != nil {
		l = m.Targets.Size()
		n += 1 + l + sovRouter(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Target) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GroupId)
	if l > 0 {
		n += 1 + l + sovRouter(uint64(l))
	}
	if len(m.Uids) > 0 {
		for _, s := range m.Uids {
			l = len(s)
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	if len(m.UnlessUids) > 0 {
		for _, s := range m.UnlessUids {
			l = len(s)
			n += 1 + l + sovRouter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRouter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRouter(x uint64) (n int) {
	return sovRouter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReqArg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
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
			return fmt.Errorf("proto: ReqArg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReqArg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			m.MsgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRouter
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
func (m *RespMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
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
			return fmt.Errorf("proto: RespMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RespMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			m.MsgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRouter
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
func (m *QueueMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
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
			return fmt.Errorf("proto: QueueMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueueMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Targets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Targets == nil {
				m.Targets = &Target{}
			}
			if err := m.Targets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRouter
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
func (m *Target) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRouter
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
			return fmt.Errorf("proto: Target: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Target: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uids = append(m.Uids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlessUids", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRouter
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
				return ErrInvalidLengthRouter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRouter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnlessUids = append(m.UnlessUids, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRouter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRouter
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
func skipRouter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRouter
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
					return 0, ErrIntOverflowRouter
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
					return 0, ErrIntOverflowRouter
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
				return 0, ErrInvalidLengthRouter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRouter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRouter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRouter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRouter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRouter = fmt.Errorf("proto: unexpected end of group")
)
