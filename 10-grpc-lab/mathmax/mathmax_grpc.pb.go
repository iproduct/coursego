// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mathmax

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

// MathClient is the hello-bidi-streaming-client API for Math service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MathClient interface {
	Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error)
}

type mathClient struct {
	cc grpc.ClientConnInterface
}

func NewMathClient(cc grpc.ClientConnInterface) MathClient {
	return &mathClient{cc}
}

func (c *mathClient) Max(ctx context.Context, opts ...grpc.CallOption) (Math_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &Math_ServiceDesc.Streams[0], "/protobuf.Math/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &mathMaxClient{stream}
	return x, nil
}

type Math_MaxClient interface {
	Send(*IntRequest) error
	Recv() (*IntResponse, error)
	grpc.ClientStream
}

type mathMaxClient struct {
	grpc.ClientStream
}

func (x *mathMaxClient) Send(m *IntRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mathMaxClient) Recv() (*IntResponse, error) {
	m := new(IntResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MathServer is the hello-server API for Math service.
// All implementations must embed UnimplementedMathServer
// for forward compatibility
type MathServer interface {
	Max(Math_MaxServer) error
	mustEmbedUnimplementedMathServer()
}

// UnimplementedMathServer must be embedded to have forward compatible implementations.
type UnimplementedMathServer struct {
}

func (UnimplementedMathServer) Max(Math_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedMathServer) mustEmbedUnimplementedMathServer() {}

// UnsafeMathServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MathServer will
// result in compilation errors.
type UnsafeMathServer interface {
	mustEmbedUnimplementedMathServer()
}

func RegisterMathServer(s grpc.ServiceRegistrar, srv MathServer) {
	s.RegisterService(&Math_ServiceDesc, srv)
}

func _Math_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MathServer).Max(&mathMaxServer{stream})
}

type Math_MaxServer interface {
	Send(*IntResponse) error
	Recv() (*IntRequest, error)
	grpc.ServerStream
}

type mathMaxServer struct {
	grpc.ServerStream
}

func (x *mathMaxServer) Send(m *IntResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mathMaxServer) Recv() (*IntRequest, error) {
	m := new(IntRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Math_ServiceDesc is the grpc.ServiceDesc for Math service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Math_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Math",
	HandlerType: (*MathServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Max",
			Handler:       _Math_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mathmax.proto",
}
