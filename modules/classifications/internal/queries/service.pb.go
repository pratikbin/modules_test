// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/classifications/internal/queries/service.proto

package queries

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	classification "github.com/persistenceOne/persistenceSDK/modules/classifications/internal/queries/classification"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	proto.RegisterFile("persistence_sdk/modules/classifications/internal/queries/service.proto", fileDescriptor_33af74aa9a0b235e)
}

var fileDescriptor_33af74aa9a0b235e = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2b, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2c, 0x4d,
	0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xb2, 0x40, 0x33, 0x47, 0x0f, 0x6a, 0x8e, 0x1e, 0x9a, 0x39, 0x7a, 0x30, 0x73,
	0xf4, 0xa0, 0xe6, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x0d, 0xd1, 0x07, 0xb1, 0x20, 0xe6,
	0x49, 0x05, 0x93, 0xed, 0x2e, 0x54, 0x05, 0x60, 0xe1, 0x4a, 0xa8, 0xa1, 0x32, 0xe9, 0xf9, 0xf9,
	0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x50, 0x37, 0x81,
	0x65, 0x8d, 0xbe, 0x31, 0x72, 0xb1, 0x06, 0x82, 0x54, 0x0b, 0xbd, 0x61, 0xe4, 0xe2, 0x73, 0x46,
	0x31, 0x46, 0x28, 0x54, 0x8f, 0x5c, 0x0f, 0xa2, 0x29, 0xd0, 0x03, 0x5b, 0x11, 0x94, 0x5a, 0x58,
	0x9a, 0x5a, 0x5c, 0x22, 0x15, 0x46, 0x6d, 0x63, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0xb4,
	0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x2a, 0xa4, 0xac, 0x0f, 0x0a, 0x3b, 0xf4, 0x30, 0x43, 0xe5,
	0x3b, 0x65, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x54, 0x60, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x92, 0x43, 0xfd, 0xf3, 0x52, 0x91, 0xb9,
	0xc1, 0x2e, 0xde, 0x44, 0x47, 0x4f, 0x12, 0x1b, 0x38, 0xb0, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x3c, 0x45, 0x3d, 0x79, 0x02, 0x00, 0x00,
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
	Classification(ctx context.Context, in *classification.QueryRequest, opts ...grpc.CallOption) (*classification.QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Classification(ctx context.Context, in *classification.QueryRequest, opts ...grpc.CallOption) (*classification.QueryResponse, error) {
	out := new(classification.QueryResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.classifications.internal.queries.Query/Classification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Classification(context.Context, *classification.QueryRequest) (*classification.QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Classification(ctx context.Context, req *classification.QueryRequest) (*classification.QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classification not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Classification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(classification.QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Classification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.classifications.internal.queries.Query/Classification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Classification(ctx, req.(*classification.QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.classifications.internal.queries.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classification",
			Handler:    _Query_Classification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/classifications/internal/queries/service.proto",
}
