// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: valset/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgAddExternalChainInfoForValidator struct {
	Creator    string               `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainInfos []*ExternalChainInfo `protobuf:"bytes,2,rep,name=chainInfos,proto3" json:"chainInfos,omitempty"`
}

func (m *MsgAddExternalChainInfoForValidator) Reset()         { *m = MsgAddExternalChainInfoForValidator{} }
func (m *MsgAddExternalChainInfoForValidator) String() string { return proto.CompactTextString(m) }
func (*MsgAddExternalChainInfoForValidator) ProtoMessage()    {}
func (*MsgAddExternalChainInfoForValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f173a384e172ec1f, []int{0}
}
func (m *MsgAddExternalChainInfoForValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddExternalChainInfoForValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddExternalChainInfoForValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddExternalChainInfoForValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddExternalChainInfoForValidator.Merge(m, src)
}
func (m *MsgAddExternalChainInfoForValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddExternalChainInfoForValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddExternalChainInfoForValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddExternalChainInfoForValidator proto.InternalMessageInfo

func (m *MsgAddExternalChainInfoForValidator) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddExternalChainInfoForValidator) GetChainInfos() []*ExternalChainInfo {
	if m != nil {
		return m.ChainInfos
	}
	return nil
}

type MsgAddExternalChainInfoForValidatorResponse struct {
}

func (m *MsgAddExternalChainInfoForValidatorResponse) Reset() {
	*m = MsgAddExternalChainInfoForValidatorResponse{}
}
func (m *MsgAddExternalChainInfoForValidatorResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgAddExternalChainInfoForValidatorResponse) ProtoMessage() {}
func (*MsgAddExternalChainInfoForValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f173a384e172ec1f, []int{1}
}
func (m *MsgAddExternalChainInfoForValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddExternalChainInfoForValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddExternalChainInfoForValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddExternalChainInfoForValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddExternalChainInfoForValidatorResponse.Merge(m, src)
}
func (m *MsgAddExternalChainInfoForValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddExternalChainInfoForValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddExternalChainInfoForValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddExternalChainInfoForValidatorResponse proto.InternalMessageInfo

type MsgKeepAlive struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgKeepAlive) Reset()         { *m = MsgKeepAlive{} }
func (m *MsgKeepAlive) String() string { return proto.CompactTextString(m) }
func (*MsgKeepAlive) ProtoMessage()    {}
func (*MsgKeepAlive) Descriptor() ([]byte, []int) {
	return fileDescriptor_f173a384e172ec1f, []int{2}
}
func (m *MsgKeepAlive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgKeepAlive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgKeepAlive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgKeepAlive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgKeepAlive.Merge(m, src)
}
func (m *MsgKeepAlive) XXX_Size() int {
	return m.Size()
}
func (m *MsgKeepAlive) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgKeepAlive.DiscardUnknown(m)
}

var xxx_messageInfo_MsgKeepAlive proto.InternalMessageInfo

func (m *MsgKeepAlive) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgKeepAliveResponse struct {
}

func (m *MsgKeepAliveResponse) Reset()         { *m = MsgKeepAliveResponse{} }
func (m *MsgKeepAliveResponse) String() string { return proto.CompactTextString(m) }
func (*MsgKeepAliveResponse) ProtoMessage()    {}
func (*MsgKeepAliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f173a384e172ec1f, []int{3}
}
func (m *MsgKeepAliveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgKeepAliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgKeepAliveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgKeepAliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgKeepAliveResponse.Merge(m, src)
}
func (m *MsgKeepAliveResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgKeepAliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgKeepAliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgKeepAliveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddExternalChainInfoForValidator)(nil), "volumefi.paloma.valset.MsgAddExternalChainInfoForValidator")
	proto.RegisterType((*MsgAddExternalChainInfoForValidatorResponse)(nil), "volumefi.paloma.valset.MsgAddExternalChainInfoForValidatorResponse")
	proto.RegisterType((*MsgKeepAlive)(nil), "volumefi.paloma.valset.MsgKeepAlive")
	proto.RegisterType((*MsgKeepAliveResponse)(nil), "volumefi.paloma.valset.MsgKeepAliveResponse")
}

func init() { proto.RegisterFile("valset/tx.proto", fileDescriptor_f173a384e172ec1f) }

var fileDescriptor_f173a384e172ec1f = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4b, 0xcc, 0x29,
	0x4e, 0x2d, 0xd1, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0xcb, 0xcf,
	0x29, 0xcd, 0x4d, 0x4d, 0xcb, 0xd4, 0x2b, 0x48, 0xcc, 0xc9, 0xcf, 0x4d, 0xd4, 0x83, 0x28, 0x90,
	0x12, 0x85, 0x2a, 0x2c, 0xce, 0x4b, 0x2c, 0x28, 0xce, 0xc8, 0x2f, 0x81, 0x28, 0x57, 0xea, 0x62,
	0xe4, 0x52, 0xf6, 0x2d, 0x4e, 0x77, 0x4c, 0x49, 0x71, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc,
	0x71, 0xce, 0x48, 0xcc, 0xcc, 0xf3, 0xcc, 0x4b, 0xcb, 0x77, 0xcb, 0x2f, 0x0a, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x12, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x05, 0x31, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x4f, 0x2e, 0xae, 0x64, 0x98, 0x96, 0x62, 0x09,
	0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x4d, 0x3d, 0xec, 0xae, 0xd0, 0xc3, 0xb0, 0x24, 0x08, 0x49,
	0xb3, 0x92, 0x2e, 0x97, 0x36, 0x11, 0x6e, 0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55,
	0xd2, 0xe0, 0xe2, 0xf1, 0x2d, 0x4e, 0xf7, 0x4e, 0x4d, 0x2d, 0x70, 0xcc, 0xc9, 0x2c, 0x4b, 0xc5,
	0xed, 0x46, 0x25, 0x31, 0x2e, 0x11, 0x64, 0x95, 0x30, 0x13, 0x8c, 0x7a, 0x99, 0xb8, 0x98, 0x7d,
	0x8b, 0xd3, 0x85, 0x96, 0x30, 0x72, 0x29, 0x10, 0x0c, 0x02, 0x6b, 0x5c, 0x9e, 0x22, 0xc2, 0xcd,
	0x52, 0xce, 0x14, 0x68, 0x86, 0x39, 0x57, 0x28, 0x9e, 0x8b, 0x13, 0xe1, 0x5b, 0x15, 0x3c, 0x26,
	0xc2, 0x55, 0x49, 0xe9, 0x10, 0xa3, 0x0a, 0x66, 0x81, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x43, 0x0c, 0x02, 0xc7, 0x1b, 0x94, 0xad, 0x5f, 0xa1, 0x0f, 0x4b, 0x87, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0xe0, 0xc4, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x94, 0x73, 0x86, 0xa1,
	0x9e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddExternalChainInfoForValidator(ctx context.Context, in *MsgAddExternalChainInfoForValidator, opts ...grpc.CallOption) (*MsgAddExternalChainInfoForValidatorResponse, error)
	KeepAlive(ctx context.Context, in *MsgKeepAlive, opts ...grpc.CallOption) (*MsgKeepAliveResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddExternalChainInfoForValidator(ctx context.Context, in *MsgAddExternalChainInfoForValidator, opts ...grpc.CallOption) (*MsgAddExternalChainInfoForValidatorResponse, error) {
	out := new(MsgAddExternalChainInfoForValidatorResponse)
	err := c.cc.Invoke(ctx, "/volumefi.paloma.valset.Msg/AddExternalChainInfoForValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) KeepAlive(ctx context.Context, in *MsgKeepAlive, opts ...grpc.CallOption) (*MsgKeepAliveResponse, error) {
	out := new(MsgKeepAliveResponse)
	err := c.cc.Invoke(ctx, "/volumefi.paloma.valset.Msg/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddExternalChainInfoForValidator(context.Context, *MsgAddExternalChainInfoForValidator) (*MsgAddExternalChainInfoForValidatorResponse, error)
	KeepAlive(context.Context, *MsgKeepAlive) (*MsgKeepAliveResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddExternalChainInfoForValidator(ctx context.Context, req *MsgAddExternalChainInfoForValidator) (*MsgAddExternalChainInfoForValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExternalChainInfoForValidator not implemented")
}
func (*UnimplementedMsgServer) KeepAlive(ctx context.Context, req *MsgKeepAlive) (*MsgKeepAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddExternalChainInfoForValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddExternalChainInfoForValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddExternalChainInfoForValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volumefi.paloma.valset.Msg/AddExternalChainInfoForValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddExternalChainInfoForValidator(ctx, req.(*MsgAddExternalChainInfoForValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgKeepAlive)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volumefi.paloma.valset.Msg/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).KeepAlive(ctx, req.(*MsgKeepAlive))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "volumefi.paloma.valset.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddExternalChainInfoForValidator",
			Handler:    _Msg_AddExternalChainInfoForValidator_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Msg_KeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "valset/tx.proto",
}

func (m *MsgAddExternalChainInfoForValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddExternalChainInfoForValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddExternalChainInfoForValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainInfos) > 0 {
		for iNdEx := len(m.ChainInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddExternalChainInfoForValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddExternalChainInfoForValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddExternalChainInfoForValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgKeepAlive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgKeepAlive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgKeepAlive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgKeepAliveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgKeepAliveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgKeepAliveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddExternalChainInfoForValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ChainInfos) > 0 {
		for _, e := range m.ChainInfos {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgAddExternalChainInfoForValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgKeepAlive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgKeepAliveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddExternalChainInfoForValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddExternalChainInfoForValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddExternalChainInfoForValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainInfos = append(m.ChainInfos, &ExternalChainInfo{})
			if err := m.ChainInfos[len(m.ChainInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddExternalChainInfoForValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddExternalChainInfoForValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddExternalChainInfoForValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgKeepAlive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgKeepAlive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgKeepAlive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgKeepAliveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgKeepAliveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgKeepAliveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
