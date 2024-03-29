// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/tx.proto

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

// this line is used by starport scaffolding # proto/tx/message
type MsgSendBuyOrder struct {
	Sender           string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Port             string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	ChannelID        string `protobuf:"bytes,3,opt,name=channelID,proto3" json:"channelID,omitempty"`
	TimeoutTimestamp uint64 `protobuf:"varint,4,opt,name=timeoutTimestamp,proto3" json:"timeoutTimestamp,omitempty"`
	AmountDenom      string `protobuf:"bytes,5,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	Amount           int32  `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	PriceDenom       string `protobuf:"bytes,7,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Price            int32  `protobuf:"varint,8,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MsgSendBuyOrder) Reset()         { *m = MsgSendBuyOrder{} }
func (m *MsgSendBuyOrder) String() string { return proto.CompactTextString(m) }
func (*MsgSendBuyOrder) ProtoMessage()    {}
func (*MsgSendBuyOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb3024a2a1faa2e9, []int{0}
}
func (m *MsgSendBuyOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendBuyOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendBuyOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendBuyOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendBuyOrder.Merge(m, src)
}
func (m *MsgSendBuyOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendBuyOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendBuyOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendBuyOrder proto.InternalMessageInfo

func (m *MsgSendBuyOrder) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSendBuyOrder) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *MsgSendBuyOrder) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *MsgSendBuyOrder) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgSendBuyOrder) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *MsgSendBuyOrder) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *MsgSendBuyOrder) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *MsgSendBuyOrder) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

type MsgSendBuyOrderResponse struct {
}

func (m *MsgSendBuyOrderResponse) Reset()         { *m = MsgSendBuyOrderResponse{} }
func (m *MsgSendBuyOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendBuyOrderResponse) ProtoMessage()    {}
func (*MsgSendBuyOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb3024a2a1faa2e9, []int{1}
}
func (m *MsgSendBuyOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendBuyOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendBuyOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendBuyOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendBuyOrderResponse.Merge(m, src)
}
func (m *MsgSendBuyOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendBuyOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendBuyOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendBuyOrderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendBuyOrder)(nil), "darthlukan.starporttutorialapp.ibcdex.MsgSendBuyOrder")
	proto.RegisterType((*MsgSendBuyOrderResponse)(nil), "darthlukan.starporttutorialapp.ibcdex.MsgSendBuyOrderResponse")
}

func init() { proto.RegisterFile("ibcdex/tx.proto", fileDescriptor_fb3024a2a1faa2e9) }

var fileDescriptor_fb3024a2a1faa2e9 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x1b, 0xfb, 0x47, 0x1b, 0x85, 0x4a, 0x10, 0x8d, 0x22, 0xe1, 0x28, 0x08, 0x45, 0xe8,
	0x1d, 0x28, 0x38, 0x38, 0x38, 0x94, 0x2e, 0x0e, 0x45, 0x38, 0x3b, 0xb9, 0xa5, 0x77, 0x2f, 0xed,
	0x61, 0x2f, 0x09, 0x49, 0x0e, 0xda, 0x2f, 0xe0, 0x28, 0x7e, 0x2c, 0xc7, 0x8e, 0x8e, 0xd2, 0x7e,
	0x08, 0x57, 0x69, 0xae, 0x47, 0x6b, 0x5d, 0xc4, 0x2d, 0xcf, 0x2f, 0x79, 0xde, 0xe7, 0xcd, 0xcb,
	0x8b, 0x1b, 0xc9, 0x20, 0x8a, 0x61, 0x12, 0xd8, 0x89, 0xaf, 0xb4, 0xb4, 0x92, 0x5c, 0xc4, 0x5c,
	0xdb, 0xd1, 0x38, 0x7b, 0xe6, 0xc2, 0x37, 0x96, 0x6b, 0x25, 0xb5, 0xb5, 0x99, 0x95, 0x3a, 0xe1,
	0x63, 0xae, 0x94, 0x9f, 0xbf, 0x6f, 0x7e, 0x21, 0xdc, 0xe8, 0x99, 0xe1, 0x23, 0x88, 0xb8, 0x93,
	0x4d, 0x1f, 0x74, 0x0c, 0x9a, 0x1c, 0xe3, 0x9a, 0x01, 0x11, 0x83, 0xa6, 0xc8, 0x43, 0xad, 0x7a,
	0xb8, 0x52, 0x84, 0xe0, 0xca, 0xb2, 0x0a, 0xdd, 0x71, 0xd4, 0x9d, 0xc9, 0x39, 0xae, 0x47, 0x23,
	0x2e, 0x04, 0x8c, 0xef, 0xbb, 0xb4, 0xec, 0x2e, 0xd6, 0x80, 0x5c, 0xe2, 0x43, 0x9b, 0xa4, 0x20,
	0x33, 0xdb, 0x4f, 0x52, 0x30, 0x96, 0xa7, 0x8a, 0x56, 0x3c, 0xd4, 0xaa, 0x84, 0xbf, 0x38, 0xf1,
	0xf0, 0x3e, 0x4f, 0x65, 0x26, 0x6c, 0x17, 0x84, 0x4c, 0x69, 0xd5, 0xd5, 0xda, 0x44, 0xcb, 0xbe,
	0x72, 0x49, 0x6b, 0x1e, 0x6a, 0x55, 0xc3, 0x95, 0x22, 0x0c, 0x63, 0xa5, 0x93, 0x08, 0x72, 0xe3,
	0xae, 0x33, 0x6e, 0x10, 0x72, 0x84, 0xab, 0x4e, 0xd1, 0x3d, 0x67, 0xcb, 0x45, 0xf3, 0x14, 0x9f,
	0x6c, 0x7d, 0x3c, 0x04, 0xa3, 0xa4, 0x30, 0x70, 0xf5, 0x8a, 0x70, 0xb9, 0x67, 0x86, 0xe4, 0x05,
	0xe1, 0x83, 0x1f, 0x93, 0xb9, 0xf1, 0xff, 0x34, 0x55, 0x7f, 0xab, 0xf0, 0xd9, 0xdd, 0xff, 0x7c,
	0x45, 0x43, 0x9d, 0xfe, 0xfb, 0x9c, 0xa1, 0xd9, 0x9c, 0xa1, 0xcf, 0x39, 0x43, 0x6f, 0x0b, 0x56,
	0x9a, 0x2d, 0x58, 0xe9, 0x63, 0xc1, 0x4a, 0x4f, 0xb7, 0xc3, 0xc4, 0x8e, 0xb2, 0x81, 0x1f, 0xc9,
	0x34, 0x58, 0x67, 0x04, 0x45, 0x46, 0xbb, 0x08, 0x69, 0x73, 0xa5, 0x82, 0x49, 0x50, 0x6c, 0xc9,
	0x54, 0x81, 0x19, 0xd4, 0xdc, 0xa6, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x32, 0x2d,
	0x5a, 0x3c, 0x02, 0x00, 0x00,
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
	// this line is used by starport scaffolding # proto/tx/rpc
	SendBuyOrder(ctx context.Context, in *MsgSendBuyOrder, opts ...grpc.CallOption) (*MsgSendBuyOrderResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendBuyOrder(ctx context.Context, in *MsgSendBuyOrder, opts ...grpc.CallOption) (*MsgSendBuyOrderResponse, error) {
	out := new(MsgSendBuyOrderResponse)
	err := c.cc.Invoke(ctx, "/darthlukan.starporttutorialapp.ibcdex.Msg/SendBuyOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	SendBuyOrder(context.Context, *MsgSendBuyOrder) (*MsgSendBuyOrderResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendBuyOrder(ctx context.Context, req *MsgSendBuyOrder) (*MsgSendBuyOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBuyOrder not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendBuyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendBuyOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendBuyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darthlukan.starporttutorialapp.ibcdex.Msg/SendBuyOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendBuyOrder(ctx, req.(*MsgSendBuyOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "darthlukan.starporttutorialapp.ibcdex.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBuyOrder",
			Handler:    _Msg_SendBuyOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibcdex/tx.proto",
}

func (m *MsgSendBuyOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendBuyOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendBuyOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x40
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Amount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendBuyOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendBuyOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendBuyOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSendBuyOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTx(uint64(m.Amount))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovTx(uint64(m.Price))
	}
	return n
}

func (m *MsgSendBuyOrderResponse) Size() (n int) {
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
func (m *MsgSendBuyOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendBuyOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendBuyOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
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
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
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
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
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
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
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
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgSendBuyOrderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSendBuyOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendBuyOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
