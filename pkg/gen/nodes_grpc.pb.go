// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: nodes.proto

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

// NodeApiClient is the client API for NodeApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeApiClient interface {
	GetNodeRegistration(ctx context.Context, in *GetNodeRegistrationRequest, opts ...grpc.CallOption) (*NodeRegistration, error)
	CreateNodeRegistration(ctx context.Context, in *CreateNodeRegistrationRequest, opts ...grpc.CallOption) (*NodeRegistration, error)
	ListNodeRegistrations(ctx context.Context, in *ListNodeRegistrationsRequest, opts ...grpc.CallOption) (*ListNodeRegistrationsResponse, error)
	TestNodeCommunication(ctx context.Context, in *TestNodeCommunicationRequest, opts ...grpc.CallOption) (*TestNodeCommunicationResponse, error)
}

type nodeApiClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeApiClient(cc grpc.ClientConnInterface) NodeApiClient {
	return &nodeApiClient{cc}
}

func (c *nodeApiClient) GetNodeRegistration(ctx context.Context, in *GetNodeRegistrationRequest, opts ...grpc.CallOption) (*NodeRegistration, error) {
	out := new(NodeRegistration)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.NodeApi/GetNodeRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeApiClient) CreateNodeRegistration(ctx context.Context, in *CreateNodeRegistrationRequest, opts ...grpc.CallOption) (*NodeRegistration, error) {
	out := new(NodeRegistration)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.NodeApi/CreateNodeRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeApiClient) ListNodeRegistrations(ctx context.Context, in *ListNodeRegistrationsRequest, opts ...grpc.CallOption) (*ListNodeRegistrationsResponse, error) {
	out := new(ListNodeRegistrationsResponse)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.NodeApi/ListNodeRegistrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeApiClient) TestNodeCommunication(ctx context.Context, in *TestNodeCommunicationRequest, opts ...grpc.CallOption) (*TestNodeCommunicationResponse, error) {
	out := new(TestNodeCommunicationResponse)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.NodeApi/TestNodeCommunication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeApiServer is the server API for NodeApi service.
// All implementations must embed UnimplementedNodeApiServer
// for forward compatibility
type NodeApiServer interface {
	GetNodeRegistration(context.Context, *GetNodeRegistrationRequest) (*NodeRegistration, error)
	CreateNodeRegistration(context.Context, *CreateNodeRegistrationRequest) (*NodeRegistration, error)
	ListNodeRegistrations(context.Context, *ListNodeRegistrationsRequest) (*ListNodeRegistrationsResponse, error)
	TestNodeCommunication(context.Context, *TestNodeCommunicationRequest) (*TestNodeCommunicationResponse, error)
	mustEmbedUnimplementedNodeApiServer()
}

// UnimplementedNodeApiServer must be embedded to have forward compatible implementations.
type UnimplementedNodeApiServer struct {
}

func (UnimplementedNodeApiServer) GetNodeRegistration(context.Context, *GetNodeRegistrationRequest) (*NodeRegistration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeRegistration not implemented")
}
func (UnimplementedNodeApiServer) CreateNodeRegistration(context.Context, *CreateNodeRegistrationRequest) (*NodeRegistration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNodeRegistration not implemented")
}
func (UnimplementedNodeApiServer) ListNodeRegistrations(context.Context, *ListNodeRegistrationsRequest) (*ListNodeRegistrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodeRegistrations not implemented")
}
func (UnimplementedNodeApiServer) TestNodeCommunication(context.Context, *TestNodeCommunicationRequest) (*TestNodeCommunicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestNodeCommunication not implemented")
}
func (UnimplementedNodeApiServer) mustEmbedUnimplementedNodeApiServer() {}

// UnsafeNodeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeApiServer will
// result in compilation errors.
type UnsafeNodeApiServer interface {
	mustEmbedUnimplementedNodeApiServer()
}

func RegisterNodeApiServer(s grpc.ServiceRegistrar, srv NodeApiServer) {
	s.RegisterService(&NodeApi_ServiceDesc, srv)
}

func _NodeApi_GetNodeRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeApiServer).GetNodeRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.NodeApi/GetNodeRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeApiServer).GetNodeRegistration(ctx, req.(*GetNodeRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeApi_CreateNodeRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeApiServer).CreateNodeRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.NodeApi/CreateNodeRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeApiServer).CreateNodeRegistration(ctx, req.(*CreateNodeRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeApi_ListNodeRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRegistrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeApiServer).ListNodeRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.NodeApi/ListNodeRegistrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeApiServer).ListNodeRegistrations(ctx, req.(*ListNodeRegistrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeApi_TestNodeCommunication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestNodeCommunicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeApiServer).TestNodeCommunication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.NodeApi/TestNodeCommunication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeApiServer).TestNodeCommunication(ctx, req.(*TestNodeCommunicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeApi_ServiceDesc is the grpc.ServiceDesc for NodeApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vanti.bsp.ew.NodeApi",
	HandlerType: (*NodeApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeRegistration",
			Handler:    _NodeApi_GetNodeRegistration_Handler,
		},
		{
			MethodName: "CreateNodeRegistration",
			Handler:    _NodeApi_CreateNodeRegistration_Handler,
		},
		{
			MethodName: "ListNodeRegistrations",
			Handler:    _NodeApi_ListNodeRegistrations_Handler,
		},
		{
			MethodName: "TestNodeCommunication",
			Handler:    _NodeApi_TestNodeCommunication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodes.proto",
}
