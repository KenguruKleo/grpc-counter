// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package counter

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

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	UpdateCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) UpdateCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/counter.CounterService/UpdateCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations should embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	UpdateCount(context.Context, *Empty) (*Response, error)
}

// UnimplementedCounterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) UpdateCount(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCount not implemented")
}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_UpdateCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).UpdateCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.CounterService/UpdateCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).UpdateCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "counter.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateCount",
			Handler:    _CounterService_UpdateCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/services/counter/counter.proto",
}
