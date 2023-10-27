// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: temperature.proto

package gen

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

// TemperatureApiClient is the client API for TemperatureApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemperatureApiClient interface {
	GetTemperature(ctx context.Context, in *GetTemperatureRequest, opts ...grpc.CallOption) (*Temperature, error)
	PullTemperature(ctx context.Context, in *PullTemperatureRequest, opts ...grpc.CallOption) (TemperatureApi_PullTemperatureClient, error)
	UpdateTemperature(ctx context.Context, in *UpdateTemperatureRequest, opts ...grpc.CallOption) (*Temperature, error)
}

type temperatureApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTemperatureApiClient(cc grpc.ClientConnInterface) TemperatureApiClient {
	return &temperatureApiClient{cc}
}

func (c *temperatureApiClient) GetTemperature(ctx context.Context, in *GetTemperatureRequest, opts ...grpc.CallOption) (*Temperature, error) {
	out := new(Temperature)
	err := c.cc.Invoke(ctx, "/smartcore.bos.TemperatureApi/GetTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temperatureApiClient) PullTemperature(ctx context.Context, in *PullTemperatureRequest, opts ...grpc.CallOption) (TemperatureApi_PullTemperatureClient, error) {
	stream, err := c.cc.NewStream(ctx, &TemperatureApi_ServiceDesc.Streams[0], "/smartcore.bos.TemperatureApi/PullTemperature", opts...)
	if err != nil {
		return nil, err
	}
	x := &temperatureApiPullTemperatureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TemperatureApi_PullTemperatureClient interface {
	Recv() (*PullTemperatureResponse, error)
	grpc.ClientStream
}

type temperatureApiPullTemperatureClient struct {
	grpc.ClientStream
}

func (x *temperatureApiPullTemperatureClient) Recv() (*PullTemperatureResponse, error) {
	m := new(PullTemperatureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *temperatureApiClient) UpdateTemperature(ctx context.Context, in *UpdateTemperatureRequest, opts ...grpc.CallOption) (*Temperature, error) {
	out := new(Temperature)
	err := c.cc.Invoke(ctx, "/smartcore.bos.TemperatureApi/UpdateTemperature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemperatureApiServer is the server API for TemperatureApi service.
// All implementations must embed UnimplementedTemperatureApiServer
// for forward compatibility
type TemperatureApiServer interface {
	GetTemperature(context.Context, *GetTemperatureRequest) (*Temperature, error)
	PullTemperature(*PullTemperatureRequest, TemperatureApi_PullTemperatureServer) error
	UpdateTemperature(context.Context, *UpdateTemperatureRequest) (*Temperature, error)
	mustEmbedUnimplementedTemperatureApiServer()
}

// UnimplementedTemperatureApiServer must be embedded to have forward compatible implementations.
type UnimplementedTemperatureApiServer struct {
}

func (UnimplementedTemperatureApiServer) GetTemperature(context.Context, *GetTemperatureRequest) (*Temperature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemperature not implemented")
}
func (UnimplementedTemperatureApiServer) PullTemperature(*PullTemperatureRequest, TemperatureApi_PullTemperatureServer) error {
	return status.Errorf(codes.Unimplemented, "method PullTemperature not implemented")
}
func (UnimplementedTemperatureApiServer) UpdateTemperature(context.Context, *UpdateTemperatureRequest) (*Temperature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemperature not implemented")
}
func (UnimplementedTemperatureApiServer) mustEmbedUnimplementedTemperatureApiServer() {}

// UnsafeTemperatureApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemperatureApiServer will
// result in compilation errors.
type UnsafeTemperatureApiServer interface {
	mustEmbedUnimplementedTemperatureApiServer()
}

func RegisterTemperatureApiServer(s grpc.ServiceRegistrar, srv TemperatureApiServer) {
	s.RegisterService(&TemperatureApi_ServiceDesc, srv)
}

func _TemperatureApi_GetTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemperatureApiServer).GetTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.TemperatureApi/GetTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemperatureApiServer).GetTemperature(ctx, req.(*GetTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemperatureApi_PullTemperature_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullTemperatureRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TemperatureApiServer).PullTemperature(m, &temperatureApiPullTemperatureServer{stream})
}

type TemperatureApi_PullTemperatureServer interface {
	Send(*PullTemperatureResponse) error
	grpc.ServerStream
}

type temperatureApiPullTemperatureServer struct {
	grpc.ServerStream
}

func (x *temperatureApiPullTemperatureServer) Send(m *PullTemperatureResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TemperatureApi_UpdateTemperature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemperatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemperatureApiServer).UpdateTemperature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.TemperatureApi/UpdateTemperature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemperatureApiServer).UpdateTemperature(ctx, req.(*UpdateTemperatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemperatureApi_ServiceDesc is the grpc.ServiceDesc for TemperatureApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemperatureApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.TemperatureApi",
	HandlerType: (*TemperatureApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTemperature",
			Handler:    _TemperatureApi_GetTemperature_Handler,
		},
		{
			MethodName: "UpdateTemperature",
			Handler:    _TemperatureApi_UpdateTemperature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullTemperature",
			Handler:       _TemperatureApi_PullTemperature_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "temperature.proto",
}