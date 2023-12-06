// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: point.proto

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

// PointApiClient is the client API for PointApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointApiClient interface {
	GetPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*Points, error)
	PullPoints(ctx context.Context, in *PullPointsRequest, opts ...grpc.CallOption) (PointApi_PullPointsClient, error)
}

type pointApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPointApiClient(cc grpc.ClientConnInterface) PointApiClient {
	return &pointApiClient{cc}
}

func (c *pointApiClient) GetPoints(ctx context.Context, in *GetPointsRequest, opts ...grpc.CallOption) (*Points, error) {
	out := new(Points)
	err := c.cc.Invoke(ctx, "/smartcore.bos.PointApi/GetPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointApiClient) PullPoints(ctx context.Context, in *PullPointsRequest, opts ...grpc.CallOption) (PointApi_PullPointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PointApi_ServiceDesc.Streams[0], "/smartcore.bos.PointApi/PullPoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &pointApiPullPointsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PointApi_PullPointsClient interface {
	Recv() (*PullPointsResponse, error)
	grpc.ClientStream
}

type pointApiPullPointsClient struct {
	grpc.ClientStream
}

func (x *pointApiPullPointsClient) Recv() (*PullPointsResponse, error) {
	m := new(PullPointsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PointApiServer is the server API for PointApi service.
// All implementations must embed UnimplementedPointApiServer
// for forward compatibility
type PointApiServer interface {
	GetPoints(context.Context, *GetPointsRequest) (*Points, error)
	PullPoints(*PullPointsRequest, PointApi_PullPointsServer) error
	mustEmbedUnimplementedPointApiServer()
}

// UnimplementedPointApiServer must be embedded to have forward compatible implementations.
type UnimplementedPointApiServer struct {
}

func (UnimplementedPointApiServer) GetPoints(context.Context, *GetPointsRequest) (*Points, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoints not implemented")
}
func (UnimplementedPointApiServer) PullPoints(*PullPointsRequest, PointApi_PullPointsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullPoints not implemented")
}
func (UnimplementedPointApiServer) mustEmbedUnimplementedPointApiServer() {}

// UnsafePointApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointApiServer will
// result in compilation errors.
type UnsafePointApiServer interface {
	mustEmbedUnimplementedPointApiServer()
}

func RegisterPointApiServer(s grpc.ServiceRegistrar, srv PointApiServer) {
	s.RegisterService(&PointApi_ServiceDesc, srv)
}

func _PointApi_GetPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointApiServer).GetPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.PointApi/GetPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointApiServer).GetPoints(ctx, req.(*GetPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointApi_PullPoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullPointsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PointApiServer).PullPoints(m, &pointApiPullPointsServer{stream})
}

type PointApi_PullPointsServer interface {
	Send(*PullPointsResponse) error
	grpc.ServerStream
}

type pointApiPullPointsServer struct {
	grpc.ServerStream
}

func (x *pointApiPullPointsServer) Send(m *PullPointsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PointApi_ServiceDesc is the grpc.ServiceDesc for PointApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.PointApi",
	HandlerType: (*PointApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoints",
			Handler:    _PointApi_GetPoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullPoints",
			Handler:       _PointApi_PullPoints_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "point.proto",
}

// PointInfoClient is the client API for PointInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PointInfoClient interface {
	DescribePoints(ctx context.Context, in *DescribePointsRequest, opts ...grpc.CallOption) (*PointsSupport, error)
}

type pointInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPointInfoClient(cc grpc.ClientConnInterface) PointInfoClient {
	return &pointInfoClient{cc}
}

func (c *pointInfoClient) DescribePoints(ctx context.Context, in *DescribePointsRequest, opts ...grpc.CallOption) (*PointsSupport, error) {
	out := new(PointsSupport)
	err := c.cc.Invoke(ctx, "/smartcore.bos.PointInfo/DescribePoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointInfoServer is the server API for PointInfo service.
// All implementations must embed UnimplementedPointInfoServer
// for forward compatibility
type PointInfoServer interface {
	DescribePoints(context.Context, *DescribePointsRequest) (*PointsSupport, error)
	mustEmbedUnimplementedPointInfoServer()
}

// UnimplementedPointInfoServer must be embedded to have forward compatible implementations.
type UnimplementedPointInfoServer struct {
}

func (UnimplementedPointInfoServer) DescribePoints(context.Context, *DescribePointsRequest) (*PointsSupport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePoints not implemented")
}
func (UnimplementedPointInfoServer) mustEmbedUnimplementedPointInfoServer() {}

// UnsafePointInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointInfoServer will
// result in compilation errors.
type UnsafePointInfoServer interface {
	mustEmbedUnimplementedPointInfoServer()
}

func RegisterPointInfoServer(s grpc.ServiceRegistrar, srv PointInfoServer) {
	s.RegisterService(&PointInfo_ServiceDesc, srv)
}

func _PointInfo_DescribePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointInfoServer).DescribePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.PointInfo/DescribePoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointInfoServer).DescribePoints(ctx, req.(*DescribePointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PointInfo_ServiceDesc is the grpc.ServiceDesc for PointInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.PointInfo",
	HandlerType: (*PointInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePoints",
			Handler:    _PointInfo_DescribePoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "point.proto",
}
