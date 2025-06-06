// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: mental_wellness.proto

package proto

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
	MentalWellnessBot_GenerateResponse_FullMethodName = "/mental_wellness_bot.MentalWellnessBot/GenerateResponse"
)

// MentalWellnessBotClient is the client API for MentalWellnessBot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MentalWellnessBotClient interface {
	GenerateResponse(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*BotResponse, error)
}

type mentalWellnessBotClient struct {
	cc grpc.ClientConnInterface
}

func NewMentalWellnessBotClient(cc grpc.ClientConnInterface) MentalWellnessBotClient {
	return &mentalWellnessBotClient{cc}
}

func (c *mentalWellnessBotClient) GenerateResponse(ctx context.Context, in *UserInput, opts ...grpc.CallOption) (*BotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BotResponse)
	err := c.cc.Invoke(ctx, MentalWellnessBot_GenerateResponse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MentalWellnessBotServer is the server API for MentalWellnessBot service.
// All implementations must embed UnimplementedMentalWellnessBotServer
// for forward compatibility.
type MentalWellnessBotServer interface {
	GenerateResponse(context.Context, *UserInput) (*BotResponse, error)
	mustEmbedUnimplementedMentalWellnessBotServer()
}

// UnimplementedMentalWellnessBotServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMentalWellnessBotServer struct{}

func (UnimplementedMentalWellnessBotServer) GenerateResponse(context.Context, *UserInput) (*BotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateResponse not implemented")
}
func (UnimplementedMentalWellnessBotServer) mustEmbedUnimplementedMentalWellnessBotServer() {}
func (UnimplementedMentalWellnessBotServer) testEmbeddedByValue()                           {}

// UnsafeMentalWellnessBotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MentalWellnessBotServer will
// result in compilation errors.
type UnsafeMentalWellnessBotServer interface {
	mustEmbedUnimplementedMentalWellnessBotServer()
}

func RegisterMentalWellnessBotServer(s grpc.ServiceRegistrar, srv MentalWellnessBotServer) {
	// If the following call pancis, it indicates UnimplementedMentalWellnessBotServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MentalWellnessBot_ServiceDesc, srv)
}

func _MentalWellnessBot_GenerateResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MentalWellnessBotServer).GenerateResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MentalWellnessBot_GenerateResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MentalWellnessBotServer).GenerateResponse(ctx, req.(*UserInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MentalWellnessBot_ServiceDesc is the grpc.ServiceDesc for MentalWellnessBot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MentalWellnessBot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mental_wellness_bot.MentalWellnessBot",
	HandlerType: (*MentalWellnessBotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateResponse",
			Handler:    _MentalWellnessBot_GenerateResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mental_wellness.proto",
}
