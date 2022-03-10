// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package create_data

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CreateDataCollectClient is the client API for CreateDataCollect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateDataCollectClient interface {
	CreateData(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataResponse, error)
}

type createDataCollectClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateDataCollectClient(cc grpc.ClientConnInterface) CreateDataCollectClient {
	return &createDataCollectClient{cc}
}

func (c *createDataCollectClient) CreateData(ctx context.Context, in *CreateDataRequest, opts ...grpc.CallOption) (*CreateDataResponse, error) {
	out := new(CreateDataResponse)
	err := c.cc.Invoke(ctx, "/rpctest.CreateDataCollect/CreateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateDataCollectServer is the server API for CreateDataCollect service.
// All implementations must embed UnimplementedCreateDataCollectServer
// for forward compatibility
type CreateDataCollectServer interface {
	CreateData(context.Context, *CreateDataRequest) (*CreateDataResponse, error)
	mustEmbedUnimplementedCreateDataCollectServer()
}

// UnimplementedCreateDataCollectServer must be embedded to have forward compatible implementations.
type UnimplementedCreateDataCollectServer struct {
}

func (UnimplementedCreateDataCollectServer) CreateData(context.Context, *CreateDataRequest) (*CreateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateData not implemented")
}
func (UnimplementedCreateDataCollectServer) mustEmbedUnimplementedCreateDataCollectServer() {}

// UnsafeCreateDataCollectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateDataCollectServer will
// result in compilation errors.
type UnsafeCreateDataCollectServer interface {
	mustEmbedUnimplementedCreateDataCollectServer()
}

func RegisterCreateDataCollectServer(s grpc.ServiceRegistrar, srv CreateDataCollectServer) {
	s.RegisterService(&CreateDataCollect_ServiceDesc, srv)
}

func _CreateDataCollect_CreateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateDataCollectServer).CreateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctest.CreateDataCollect/CreateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateDataCollectServer).CreateData(ctx, req.(*CreateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateDataCollect_ServiceDesc is the grpc.ServiceDesc for CreateDataCollect service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateDataCollect_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpctest.CreateDataCollect",
	HandlerType: (*CreateDataCollectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateData",
			Handler:    _CreateDataCollect_CreateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpctest.proto",
}