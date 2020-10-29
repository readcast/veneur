// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpsink

import (
	context "context"
	ssf "github.com/stripe/veneur/ssf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SpanSinkClient is the client API for SpanSink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpanSinkClient interface {
	SendSpan(ctx context.Context, in *ssf.SSFSpan, opts ...grpc.CallOption) (*Empty, error)
}

type spanSinkClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanSinkClient(cc grpc.ClientConnInterface) SpanSinkClient {
	return &spanSinkClient{cc}
}

func (c *spanSinkClient) SendSpan(ctx context.Context, in *ssf.SSFSpan, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpsink.SpanSink/SendSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpanSinkServer is the server API for SpanSink service.
// All implementations must embed UnimplementedSpanSinkServer
// for forward compatibility
type SpanSinkServer interface {
	SendSpan(context.Context, *ssf.SSFSpan) (*Empty, error)
	mustEmbedUnimplementedSpanSinkServer()
}

// UnimplementedSpanSinkServer must be embedded to have forward compatible implementations.
type UnimplementedSpanSinkServer struct {
}

func (UnimplementedSpanSinkServer) SendSpan(context.Context, *ssf.SSFSpan) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSpan not implemented")
}
func (UnimplementedSpanSinkServer) mustEmbedUnimplementedSpanSinkServer() {}

// UnsafeSpanSinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpanSinkServer will
// result in compilation errors.
type UnsafeSpanSinkServer interface {
	mustEmbedUnimplementedSpanSinkServer()
}

func RegisterSpanSinkServer(s *grpc.Server, srv SpanSinkServer) {
	s.RegisterService(&_SpanSink_serviceDesc, srv)
}

func _SpanSink_SendSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ssf.SSFSpan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanSinkServer).SendSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpsink.SpanSink/SendSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanSinkServer).SendSpan(ctx, req.(*ssf.SSFSpan))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpanSink_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpsink.SpanSink",
	HandlerType: (*SpanSinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSpan",
			Handler:    _SpanSink_SendSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sinks/grpsink/grpc_sink.proto",
}