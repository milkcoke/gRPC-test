// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greeting

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Enter(ctx context.Context, in *EnterMessage, opts ...grpc.CallOption) (*GreetMessage, error)
	Exit(ctx context.Context, in *ExitMessage, opts ...grpc.CallOption) (*GreetMessage, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Enter(ctx context.Context, in *EnterMessage, opts ...grpc.CallOption) (*GreetMessage, error) {
	out := new(GreetMessage)
	err := c.cc.Invoke(ctx, "/greeting.Greeter/Enter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Exit(ctx context.Context, in *ExitMessage, opts ...grpc.CallOption) (*GreetMessage, error) {
	out := new(GreetMessage)
	err := c.cc.Invoke(ctx, "/greeting.Greeter/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Enter(context.Context, *EnterMessage) (*GreetMessage, error)
	Exit(context.Context, *ExitMessage) (*GreetMessage, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Enter(context.Context, *EnterMessage) (*GreetMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enter not implemented")
}
func (UnimplementedGreeterServer) Exit(context.Context, *ExitMessage) (*GreetMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exit not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeting.Greeter/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Enter(ctx, req.(*EnterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExitMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeting.Greeter/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Exit(ctx, req.(*ExitMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeting.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enter",
			Handler:    _Greeter_Enter_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _Greeter_Exit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeting.proto",
}
