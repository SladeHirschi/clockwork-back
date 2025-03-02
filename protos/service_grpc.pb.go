// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: protos/service.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ClockWorkService_ClockIn_FullMethodName = "/clockwork.ClockWorkService/ClockIn"
)

// ClockWorkServiceClient is the client API for ClockWorkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClockWorkServiceClient interface {
	ClockIn(ctx context.Context, in *ClockInRequest, opts ...grpc.CallOption) (*ClockInResponse, error)
}

type clockWorkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClockWorkServiceClient(cc grpc.ClientConnInterface) ClockWorkServiceClient {
	return &clockWorkServiceClient{cc}
}

func (c *clockWorkServiceClient) ClockIn(ctx context.Context, in *ClockInRequest, opts ...grpc.CallOption) (*ClockInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClockInResponse)
	err := c.cc.Invoke(ctx, ClockWorkService_ClockIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockWorkServiceServer is the server API for ClockWorkService service.
// All implementations must embed UnimplementedClockWorkServiceServer
// for forward compatibility.
type ClockWorkServiceServer interface {
	ClockIn(context.Context, *ClockInRequest) (*ClockInResponse, error)
	mustEmbedUnimplementedClockWorkServiceServer()
}

// UnimplementedClockWorkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClockWorkServiceServer struct{}

func (UnimplementedClockWorkServiceServer) ClockIn(context.Context, *ClockInRequest) (*ClockInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClockIn not implemented")
}
func (UnimplementedClockWorkServiceServer) mustEmbedUnimplementedClockWorkServiceServer() {}
func (UnimplementedClockWorkServiceServer) testEmbeddedByValue()                          {}

// UnsafeClockWorkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClockWorkServiceServer will
// result in compilation errors.
type UnsafeClockWorkServiceServer interface {
	mustEmbedUnimplementedClockWorkServiceServer()
}

func RegisterClockWorkServiceServer(s grpc.ServiceRegistrar, srv ClockWorkServiceServer) {
	// If the following call pancis, it indicates UnimplementedClockWorkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClockWorkService_ServiceDesc, srv)
}

func _ClockWorkService_ClockIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClockInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockWorkServiceServer).ClockIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClockWorkService_ClockIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockWorkServiceServer).ClockIn(ctx, req.(*ClockInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClockWorkService_ServiceDesc is the grpc.ServiceDesc for ClockWorkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClockWorkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clockwork.ClockWorkService",
	HandlerType: (*ClockWorkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClockIn",
			Handler:    _ClockWorkService_ClockIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service.proto",
}
