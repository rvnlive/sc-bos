// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: services.proto

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

const (
	ServicesApi_GetService_FullMethodName          = "/smartcore.bos.ServicesApi/GetService"
	ServicesApi_PullService_FullMethodName         = "/smartcore.bos.ServicesApi/PullService"
	ServicesApi_CreateService_FullMethodName       = "/smartcore.bos.ServicesApi/CreateService"
	ServicesApi_DeleteService_FullMethodName       = "/smartcore.bos.ServicesApi/DeleteService"
	ServicesApi_ListServices_FullMethodName        = "/smartcore.bos.ServicesApi/ListServices"
	ServicesApi_PullServices_FullMethodName        = "/smartcore.bos.ServicesApi/PullServices"
	ServicesApi_StartService_FullMethodName        = "/smartcore.bos.ServicesApi/StartService"
	ServicesApi_ConfigureService_FullMethodName    = "/smartcore.bos.ServicesApi/ConfigureService"
	ServicesApi_StopService_FullMethodName         = "/smartcore.bos.ServicesApi/StopService"
	ServicesApi_GetServiceMetadata_FullMethodName  = "/smartcore.bos.ServicesApi/GetServiceMetadata"
	ServicesApi_PullServiceMetadata_FullMethodName = "/smartcore.bos.ServicesApi/PullServiceMetadata"
)

// ServicesApiClient is the client API for ServicesApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServicesApiClient interface {
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error)
	PullService(ctx context.Context, in *PullServiceRequest, opts ...grpc.CallOption) (ServicesApi_PullServiceClient, error)
	CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error)
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	PullServices(ctx context.Context, in *PullServicesRequest, opts ...grpc.CallOption) (ServicesApi_PullServicesClient, error)
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*Service, error)
	ConfigureService(ctx context.Context, in *ConfigureServiceRequest, opts ...grpc.CallOption) (*Service, error)
	StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*Service, error)
	// Get service metadata: how many service are there, what types exist, etc.
	GetServiceMetadata(ctx context.Context, in *GetServiceMetadataRequest, opts ...grpc.CallOption) (*ServiceMetadata, error)
	PullServiceMetadata(ctx context.Context, in *PullServiceMetadataRequest, opts ...grpc.CallOption) (ServicesApi_PullServiceMetadataClient, error)
}

type servicesApiClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesApiClient(cc grpc.ClientConnInterface) ServicesApiClient {
	return &servicesApiClient{cc}
}

func (c *servicesApiClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_GetService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) PullService(ctx context.Context, in *PullServiceRequest, opts ...grpc.CallOption) (ServicesApi_PullServiceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicesApi_ServiceDesc.Streams[0], ServicesApi_PullService_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesApiPullServiceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicesApi_PullServiceClient interface {
	Recv() (*PullServiceResponse, error)
	grpc.ClientStream
}

type servicesApiPullServiceClient struct {
	grpc.ClientStream
}

func (x *servicesApiPullServiceClient) Recv() (*PullServiceResponse, error) {
	m := new(PullServiceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicesApiClient) CreateService(ctx context.Context, in *CreateServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_CreateService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_DeleteService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, ServicesApi_ListServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) PullServices(ctx context.Context, in *PullServicesRequest, opts ...grpc.CallOption) (ServicesApi_PullServicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicesApi_ServiceDesc.Streams[1], ServicesApi_PullServices_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesApiPullServicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicesApi_PullServicesClient interface {
	Recv() (*PullServicesResponse, error)
	grpc.ClientStream
}

type servicesApiPullServicesClient struct {
	grpc.ClientStream
}

func (x *servicesApiPullServicesClient) Recv() (*PullServicesResponse, error) {
	m := new(PullServicesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *servicesApiClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_StartService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) ConfigureService(ctx context.Context, in *ConfigureServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_ConfigureService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) StopService(ctx context.Context, in *StopServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, ServicesApi_StopService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) GetServiceMetadata(ctx context.Context, in *GetServiceMetadataRequest, opts ...grpc.CallOption) (*ServiceMetadata, error) {
	out := new(ServiceMetadata)
	err := c.cc.Invoke(ctx, ServicesApi_GetServiceMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesApiClient) PullServiceMetadata(ctx context.Context, in *PullServiceMetadataRequest, opts ...grpc.CallOption) (ServicesApi_PullServiceMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServicesApi_ServiceDesc.Streams[2], ServicesApi_PullServiceMetadata_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &servicesApiPullServiceMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServicesApi_PullServiceMetadataClient interface {
	Recv() (*PullServiceMetadataResponse, error)
	grpc.ClientStream
}

type servicesApiPullServiceMetadataClient struct {
	grpc.ClientStream
}

func (x *servicesApiPullServiceMetadataClient) Recv() (*PullServiceMetadataResponse, error) {
	m := new(PullServiceMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServicesApiServer is the server API for ServicesApi service.
// All implementations must embed UnimplementedServicesApiServer
// for forward compatibility
type ServicesApiServer interface {
	GetService(context.Context, *GetServiceRequest) (*Service, error)
	PullService(*PullServiceRequest, ServicesApi_PullServiceServer) error
	CreateService(context.Context, *CreateServiceRequest) (*Service, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*Service, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	PullServices(*PullServicesRequest, ServicesApi_PullServicesServer) error
	StartService(context.Context, *StartServiceRequest) (*Service, error)
	ConfigureService(context.Context, *ConfigureServiceRequest) (*Service, error)
	StopService(context.Context, *StopServiceRequest) (*Service, error)
	// Get service metadata: how many service are there, what types exist, etc.
	GetServiceMetadata(context.Context, *GetServiceMetadataRequest) (*ServiceMetadata, error)
	PullServiceMetadata(*PullServiceMetadataRequest, ServicesApi_PullServiceMetadataServer) error
	mustEmbedUnimplementedServicesApiServer()
}

// UnimplementedServicesApiServer must be embedded to have forward compatible implementations.
type UnimplementedServicesApiServer struct {
}

func (UnimplementedServicesApiServer) GetService(context.Context, *GetServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedServicesApiServer) PullService(*PullServiceRequest, ServicesApi_PullServiceServer) error {
	return status.Errorf(codes.Unimplemented, "method PullService not implemented")
}
func (UnimplementedServicesApiServer) CreateService(context.Context, *CreateServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServicesApiServer) DeleteService(context.Context, *DeleteServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServicesApiServer) ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedServicesApiServer) PullServices(*PullServicesRequest, ServicesApi_PullServicesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullServices not implemented")
}
func (UnimplementedServicesApiServer) StartService(context.Context, *StartServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (UnimplementedServicesApiServer) ConfigureService(context.Context, *ConfigureServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureService not implemented")
}
func (UnimplementedServicesApiServer) StopService(context.Context, *StopServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopService not implemented")
}
func (UnimplementedServicesApiServer) GetServiceMetadata(context.Context, *GetServiceMetadataRequest) (*ServiceMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceMetadata not implemented")
}
func (UnimplementedServicesApiServer) PullServiceMetadata(*PullServiceMetadataRequest, ServicesApi_PullServiceMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method PullServiceMetadata not implemented")
}
func (UnimplementedServicesApiServer) mustEmbedUnimplementedServicesApiServer() {}

// UnsafeServicesApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesApiServer will
// result in compilation errors.
type UnsafeServicesApiServer interface {
	mustEmbedUnimplementedServicesApiServer()
}

func RegisterServicesApiServer(s grpc.ServiceRegistrar, srv ServicesApiServer) {
	s.RegisterService(&ServicesApi_ServiceDesc, srv)
}

func _ServicesApi_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_GetService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_PullService_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullServiceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicesApiServer).PullService(m, &servicesApiPullServiceServer{stream})
}

type ServicesApi_PullServiceServer interface {
	Send(*PullServiceResponse) error
	grpc.ServerStream
}

type servicesApiPullServiceServer struct {
	grpc.ServerStream
}

func (x *servicesApiPullServiceServer) Send(m *PullServiceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServicesApi_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).CreateService(ctx, req.(*CreateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_DeleteService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_ListServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_PullServices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullServicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicesApiServer).PullServices(m, &servicesApiPullServicesServer{stream})
}

type ServicesApi_PullServicesServer interface {
	Send(*PullServicesResponse) error
	grpc.ServerStream
}

type servicesApiPullServicesServer struct {
	grpc.ServerStream
}

func (x *servicesApiPullServicesServer) Send(m *PullServicesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ServicesApi_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_StartService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_ConfigureService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).ConfigureService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_ConfigureService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).ConfigureService(ctx, req.(*ConfigureServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_StopService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).StopService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_StopService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).StopService(ctx, req.(*StopServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_GetServiceMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesApiServer).GetServiceMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServicesApi_GetServiceMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesApiServer).GetServiceMetadata(ctx, req.(*GetServiceMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicesApi_PullServiceMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullServiceMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServicesApiServer).PullServiceMetadata(m, &servicesApiPullServiceMetadataServer{stream})
}

type ServicesApi_PullServiceMetadataServer interface {
	Send(*PullServiceMetadataResponse) error
	grpc.ServerStream
}

type servicesApiPullServiceMetadataServer struct {
	grpc.ServerStream
}

func (x *servicesApiPullServiceMetadataServer) Send(m *PullServiceMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ServicesApi_ServiceDesc is the grpc.ServiceDesc for ServicesApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.ServicesApi",
	HandlerType: (*ServicesApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetService",
			Handler:    _ServicesApi_GetService_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _ServicesApi_CreateService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _ServicesApi_DeleteService_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _ServicesApi_ListServices_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _ServicesApi_StartService_Handler,
		},
		{
			MethodName: "ConfigureService",
			Handler:    _ServicesApi_ConfigureService_Handler,
		},
		{
			MethodName: "StopService",
			Handler:    _ServicesApi_StopService_Handler,
		},
		{
			MethodName: "GetServiceMetadata",
			Handler:    _ServicesApi_GetServiceMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullService",
			Handler:       _ServicesApi_PullService_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullServices",
			Handler:       _ServicesApi_PullServices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullServiceMetadata",
			Handler:       _ServicesApi_PullServiceMetadata_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services.proto",
}
