// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: anpr_camera.proto

package gen

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
	AnprCameraApi_ListAnprEvents_FullMethodName = "/smartcore.bos.AnprCameraApi/ListAnprEvents"
	AnprCameraApi_PullAnprEvents_FullMethodName = "/smartcore.bos.AnprCameraApi/PullAnprEvents"
)

// AnprCameraApiClient is the client API for AnprCameraApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AnprCameraApi describes the capability to retrieve detection events from ANPR cameras.
type AnprCameraApiClient interface {
	ListAnprEvents(ctx context.Context, in *ListAnprEventsRequest, opts ...grpc.CallOption) (*ListAnprEventsResponse, error)
	PullAnprEvents(ctx context.Context, in *PullAnprEventsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullAnprEventsResponse], error)
}

type anprCameraApiClient struct {
	cc grpc.ClientConnInterface
}

func NewAnprCameraApiClient(cc grpc.ClientConnInterface) AnprCameraApiClient {
	return &anprCameraApiClient{cc}
}

func (c *anprCameraApiClient) ListAnprEvents(ctx context.Context, in *ListAnprEventsRequest, opts ...grpc.CallOption) (*ListAnprEventsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAnprEventsResponse)
	err := c.cc.Invoke(ctx, AnprCameraApi_ListAnprEvents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anprCameraApiClient) PullAnprEvents(ctx context.Context, in *PullAnprEventsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PullAnprEventsResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AnprCameraApi_ServiceDesc.Streams[0], AnprCameraApi_PullAnprEvents_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PullAnprEventsRequest, PullAnprEventsResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnprCameraApi_PullAnprEventsClient = grpc.ServerStreamingClient[PullAnprEventsResponse]

// AnprCameraApiServer is the server API for AnprCameraApi service.
// All implementations must embed UnimplementedAnprCameraApiServer
// for forward compatibility.
//
// AnprCameraApi describes the capability to retrieve detection events from ANPR cameras.
type AnprCameraApiServer interface {
	ListAnprEvents(context.Context, *ListAnprEventsRequest) (*ListAnprEventsResponse, error)
	PullAnprEvents(*PullAnprEventsRequest, grpc.ServerStreamingServer[PullAnprEventsResponse]) error
	mustEmbedUnimplementedAnprCameraApiServer()
}

// UnimplementedAnprCameraApiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnprCameraApiServer struct{}

func (UnimplementedAnprCameraApiServer) ListAnprEvents(context.Context, *ListAnprEventsRequest) (*ListAnprEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnprEvents not implemented")
}
func (UnimplementedAnprCameraApiServer) PullAnprEvents(*PullAnprEventsRequest, grpc.ServerStreamingServer[PullAnprEventsResponse]) error {
	return status.Errorf(codes.Unimplemented, "method PullAnprEvents not implemented")
}
func (UnimplementedAnprCameraApiServer) mustEmbedUnimplementedAnprCameraApiServer() {}
func (UnimplementedAnprCameraApiServer) testEmbeddedByValue()                       {}

// UnsafeAnprCameraApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnprCameraApiServer will
// result in compilation errors.
type UnsafeAnprCameraApiServer interface {
	mustEmbedUnimplementedAnprCameraApiServer()
}

func RegisterAnprCameraApiServer(s grpc.ServiceRegistrar, srv AnprCameraApiServer) {
	// If the following call pancis, it indicates UnimplementedAnprCameraApiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnprCameraApi_ServiceDesc, srv)
}

func _AnprCameraApi_ListAnprEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAnprEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnprCameraApiServer).ListAnprEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnprCameraApi_ListAnprEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnprCameraApiServer).ListAnprEvents(ctx, req.(*ListAnprEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnprCameraApi_PullAnprEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullAnprEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnprCameraApiServer).PullAnprEvents(m, &grpc.GenericServerStream[PullAnprEventsRequest, PullAnprEventsResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AnprCameraApi_PullAnprEventsServer = grpc.ServerStreamingServer[PullAnprEventsResponse]

// AnprCameraApi_ServiceDesc is the grpc.ServiceDesc for AnprCameraApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnprCameraApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.AnprCameraApi",
	HandlerType: (*AnprCameraApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAnprEvents",
			Handler:    _AnprCameraApi_ListAnprEvents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullAnprEvents",
			Handler:       _AnprCameraApi_PullAnprEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "anpr_camera.proto",
}
