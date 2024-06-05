// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/dclgenutil/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/dclgenutil/query.proto", fileDescriptor_405203a00d51f2d1)
}

var fileDescriptor_405203a00d51f2d1 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xac, 0xca, 0x4c, 0x4f,
	0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4f, 0xc9, 0x2c, 0x2e, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4,
	0xa7, 0x16, 0xe9, 0xa7, 0x24, 0xe7, 0xa4, 0xa7, 0xe6, 0x95, 0x96, 0x64, 0xe6, 0xe8, 0x17, 0x96,
	0xa6, 0x16, 0x55, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x99, 0xa2, 0x1a, 0xa1, 0x87, 0xc7,
	0x08, 0x3d, 0x84, 0x11, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99,
	0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x43, 0x8d, 0xd8,
	0xb9, 0x58, 0x03, 0x41, 0x76, 0x38, 0xa5, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x77, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x09,
	0xba, 0xd8, 0xbc, 0xa1, 0x8b, 0x70, 0x84, 0x2e, 0xd4, 0x23, 0x15, 0xc8, 0x5e, 0x29, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x5b, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x06, 0x8e, 0x75,
	0x08, 0x10, 0x01, 0x00, 0x00,
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
	ServiceName: "zigbeealliance.distributedcomplianceledger.dclgenutil.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "zigbeealliance/distributedcomplianceledger/dclgenutil/query.proto",
}
