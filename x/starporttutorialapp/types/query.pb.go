// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: starporttutorialapp/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("starporttutorialapp/query.proto", fileDescriptor_0c81d6a39928889d) }

var fileDescriptor_0c81d6a39928889d = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0xa1, 0x50, 0x13, 0x4a, 0x4b, 0x62, 0xd6, 0xde, 0x04, 0x26, 0xe0, 0x1b, 0xd8, 0x58,
	0xdb, 0xda, 0xcd, 0xe2, 0x66, 0xd9, 0x08, 0x3b, 0xeb, 0xee, 0xac, 0x91, 0xb7, 0xf0, 0xb1, 0xae,
	0xa4, 0xbc, 0xf2, 0x02, 0x2f, 0x72, 0x01, 0x42, 0xae, 0xa1, 0x9d, 0x7c, 0x99, 0xff, 0xfb, 0xb2,
	0xe7, 0xc0, 0xe8, 0x1d, 0x79, 0xe6, 0xc8, 0xe4, 0x0d, 0x76, 0xe8, 0x1c, 0xfc, 0x44, 0xe5, 0x87,
	0xd2, 0x79, 0x62, 0x7a, 0xac, 0xbf, 0xd0, 0x73, 0xdb, 0xc5, 0x6f, 0xb4, 0xe5, 0x01, 0x7b, 0x74,
	0xcb, 0x9f, 0x34, 0x91, 0xee, 0x14, 0xa0, 0x33, 0x80, 0xd6, 0x12, 0x23, 0x1b, 0xb2, 0x61, 0xfb,
	0x98, 0xbf, 0x34, 0x14, 0x7a, 0x0a, 0x20, 0x31, 0xa8, 0x6d, 0x0a, 0x7e, 0x2b, 0xa9, 0x18, 0x2b,
	0x70, 0xa8, 0x8d, 0x5d, 0xe1, 0x8d, 0xad, 0x1f, 0xb2, 0xbb, 0x8f, 0x85, 0x78, 0xc3, 0xd3, 0x24,
	0xd2, 0x71, 0x12, 0xe9, 0x65, 0x12, 0xe9, 0xff, 0x2c, 0x92, 0x71, 0x16, 0xc9, 0x79, 0x16, 0xc9,
	0xe7, 0xbb, 0x36, 0xdc, 0x46, 0x59, 0x36, 0xd4, 0xc3, 0xcd, 0x15, 0x76, 0xaf, 0x62, 0x17, 0x2b,
	0x96, 0xb2, 0x3f, 0x38, 0xea, 0xe5, 0xc1, 0xa9, 0x20, 0xef, 0xd7, 0xc9, 0xd7, 0x6b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x19, 0xa2, 0x33, 0x43, 0x13, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "darthlukan.starporttutorialapp.starporttutorialapp.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "starporttutorialapp/query.proto",
}
