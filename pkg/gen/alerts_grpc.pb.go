// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: alerts.proto

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

// AlertApiClient is the client API for AlertApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertApiClient interface {
	ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error)
	PullAlerts(ctx context.Context, in *PullAlertsRequest, opts ...grpc.CallOption) (AlertApi_PullAlertsClient, error)
	// Acknowledge an existing alert.
	// Acknowledgement allows tracking of which alerts have been seen by an actor that may be able to react to the alert.
	AcknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*Alert, error)
	UnacknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*Alert, error)
	// Get alert metadata: how many alerts are there, what zones exist, etc.
	GetAlertMetadata(ctx context.Context, in *GetAlertMetadataRequest, opts ...grpc.CallOption) (*AlertMetadata, error)
	PullAlertMetadata(ctx context.Context, in *PullAlertMetadataRequest, opts ...grpc.CallOption) (AlertApi_PullAlertMetadataClient, error)
}

type alertApiClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertApiClient(cc grpc.ClientConnInterface) AlertApiClient {
	return &alertApiClient{cc}
}

func (c *alertApiClient) ListAlerts(ctx context.Context, in *ListAlertsRequest, opts ...grpc.CallOption) (*ListAlertsResponse, error) {
	out := new(ListAlertsResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertApi/ListAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertApiClient) PullAlerts(ctx context.Context, in *PullAlertsRequest, opts ...grpc.CallOption) (AlertApi_PullAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AlertApi_ServiceDesc.Streams[0], "/smartcore.bos.AlertApi/PullAlerts", opts...)
	if err != nil {
		return nil, err
	}
	x := &alertApiPullAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlertApi_PullAlertsClient interface {
	Recv() (*PullAlertsResponse, error)
	grpc.ClientStream
}

type alertApiPullAlertsClient struct {
	grpc.ClientStream
}

func (x *alertApiPullAlertsClient) Recv() (*PullAlertsResponse, error) {
	m := new(PullAlertsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *alertApiClient) AcknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertApi/AcknowledgeAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertApiClient) UnacknowledgeAlert(ctx context.Context, in *AcknowledgeAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertApi/UnacknowledgeAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertApiClient) GetAlertMetadata(ctx context.Context, in *GetAlertMetadataRequest, opts ...grpc.CallOption) (*AlertMetadata, error) {
	out := new(AlertMetadata)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertApi/GetAlertMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertApiClient) PullAlertMetadata(ctx context.Context, in *PullAlertMetadataRequest, opts ...grpc.CallOption) (AlertApi_PullAlertMetadataClient, error) {
	stream, err := c.cc.NewStream(ctx, &AlertApi_ServiceDesc.Streams[1], "/smartcore.bos.AlertApi/PullAlertMetadata", opts...)
	if err != nil {
		return nil, err
	}
	x := &alertApiPullAlertMetadataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AlertApi_PullAlertMetadataClient interface {
	Recv() (*PullAlertMetadataResponse, error)
	grpc.ClientStream
}

type alertApiPullAlertMetadataClient struct {
	grpc.ClientStream
}

func (x *alertApiPullAlertMetadataClient) Recv() (*PullAlertMetadataResponse, error) {
	m := new(PullAlertMetadataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AlertApiServer is the server API for AlertApi service.
// All implementations must embed UnimplementedAlertApiServer
// for forward compatibility
type AlertApiServer interface {
	ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error)
	PullAlerts(*PullAlertsRequest, AlertApi_PullAlertsServer) error
	// Acknowledge an existing alert.
	// Acknowledgement allows tracking of which alerts have been seen by an actor that may be able to react to the alert.
	AcknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*Alert, error)
	UnacknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*Alert, error)
	// Get alert metadata: how many alerts are there, what zones exist, etc.
	GetAlertMetadata(context.Context, *GetAlertMetadataRequest) (*AlertMetadata, error)
	PullAlertMetadata(*PullAlertMetadataRequest, AlertApi_PullAlertMetadataServer) error
	mustEmbedUnimplementedAlertApiServer()
}

// UnimplementedAlertApiServer must be embedded to have forward compatible implementations.
type UnimplementedAlertApiServer struct {
}

func (UnimplementedAlertApiServer) ListAlerts(context.Context, *ListAlertsRequest) (*ListAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlerts not implemented")
}
func (UnimplementedAlertApiServer) PullAlerts(*PullAlertsRequest, AlertApi_PullAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method PullAlerts not implemented")
}
func (UnimplementedAlertApiServer) AcknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeAlert not implemented")
}
func (UnimplementedAlertApiServer) UnacknowledgeAlert(context.Context, *AcknowledgeAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnacknowledgeAlert not implemented")
}
func (UnimplementedAlertApiServer) GetAlertMetadata(context.Context, *GetAlertMetadataRequest) (*AlertMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertMetadata not implemented")
}
func (UnimplementedAlertApiServer) PullAlertMetadata(*PullAlertMetadataRequest, AlertApi_PullAlertMetadataServer) error {
	return status.Errorf(codes.Unimplemented, "method PullAlertMetadata not implemented")
}
func (UnimplementedAlertApiServer) mustEmbedUnimplementedAlertApiServer() {}

// UnsafeAlertApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertApiServer will
// result in compilation errors.
type UnsafeAlertApiServer interface {
	mustEmbedUnimplementedAlertApiServer()
}

func RegisterAlertApiServer(s grpc.ServiceRegistrar, srv AlertApiServer) {
	s.RegisterService(&AlertApi_ServiceDesc, srv)
}

func _AlertApi_ListAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertApiServer).ListAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertApi/ListAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertApiServer).ListAlerts(ctx, req.(*ListAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertApi_PullAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullAlertsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlertApiServer).PullAlerts(m, &alertApiPullAlertsServer{stream})
}

type AlertApi_PullAlertsServer interface {
	Send(*PullAlertsResponse) error
	grpc.ServerStream
}

type alertApiPullAlertsServer struct {
	grpc.ServerStream
}

func (x *alertApiPullAlertsServer) Send(m *PullAlertsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AlertApi_AcknowledgeAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertApiServer).AcknowledgeAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertApi/AcknowledgeAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertApiServer).AcknowledgeAlert(ctx, req.(*AcknowledgeAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertApi_UnacknowledgeAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowledgeAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertApiServer).UnacknowledgeAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertApi/UnacknowledgeAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertApiServer).UnacknowledgeAlert(ctx, req.(*AcknowledgeAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertApi_GetAlertMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertApiServer).GetAlertMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertApi/GetAlertMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertApiServer).GetAlertMetadata(ctx, req.(*GetAlertMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertApi_PullAlertMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PullAlertMetadataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AlertApiServer).PullAlertMetadata(m, &alertApiPullAlertMetadataServer{stream})
}

type AlertApi_PullAlertMetadataServer interface {
	Send(*PullAlertMetadataResponse) error
	grpc.ServerStream
}

type alertApiPullAlertMetadataServer struct {
	grpc.ServerStream
}

func (x *alertApiPullAlertMetadataServer) Send(m *PullAlertMetadataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AlertApi_ServiceDesc is the grpc.ServiceDesc for AlertApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.AlertApi",
	HandlerType: (*AlertApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlerts",
			Handler:    _AlertApi_ListAlerts_Handler,
		},
		{
			MethodName: "AcknowledgeAlert",
			Handler:    _AlertApi_AcknowledgeAlert_Handler,
		},
		{
			MethodName: "UnacknowledgeAlert",
			Handler:    _AlertApi_UnacknowledgeAlert_Handler,
		},
		{
			MethodName: "GetAlertMetadata",
			Handler:    _AlertApi_GetAlertMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullAlerts",
			Handler:       _AlertApi_PullAlerts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullAlertMetadata",
			Handler:       _AlertApi_PullAlertMetadata_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "alerts.proto",
}

// AlertAdminApiClient is the client API for AlertAdminApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertAdminApiClient interface {
	CreateAlert(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*Alert, error)
	UpdateAlert(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*Alert, error)
	DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error)
}

type alertAdminApiClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertAdminApiClient(cc grpc.ClientConnInterface) AlertAdminApiClient {
	return &alertAdminApiClient{cc}
}

func (c *alertAdminApiClient) CreateAlert(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertAdminApi/CreateAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertAdminApiClient) UpdateAlert(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*Alert, error) {
	out := new(Alert)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertAdminApi/UpdateAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertAdminApiClient) DeleteAlert(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error) {
	out := new(DeleteAlertResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.AlertAdminApi/DeleteAlert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertAdminApiServer is the server API for AlertAdminApi service.
// All implementations must embed UnimplementedAlertAdminApiServer
// for forward compatibility
type AlertAdminApiServer interface {
	CreateAlert(context.Context, *CreateAlertRequest) (*Alert, error)
	UpdateAlert(context.Context, *UpdateAlertRequest) (*Alert, error)
	DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error)
	mustEmbedUnimplementedAlertAdminApiServer()
}

// UnimplementedAlertAdminApiServer must be embedded to have forward compatible implementations.
type UnimplementedAlertAdminApiServer struct {
}

func (UnimplementedAlertAdminApiServer) CreateAlert(context.Context, *CreateAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlert not implemented")
}
func (UnimplementedAlertAdminApiServer) UpdateAlert(context.Context, *UpdateAlertRequest) (*Alert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlert not implemented")
}
func (UnimplementedAlertAdminApiServer) DeleteAlert(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlert not implemented")
}
func (UnimplementedAlertAdminApiServer) mustEmbedUnimplementedAlertAdminApiServer() {}

// UnsafeAlertAdminApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertAdminApiServer will
// result in compilation errors.
type UnsafeAlertAdminApiServer interface {
	mustEmbedUnimplementedAlertAdminApiServer()
}

func RegisterAlertAdminApiServer(s grpc.ServiceRegistrar, srv AlertAdminApiServer) {
	s.RegisterService(&AlertAdminApi_ServiceDesc, srv)
}

func _AlertAdminApi_CreateAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertAdminApiServer).CreateAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertAdminApi/CreateAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertAdminApiServer).CreateAlert(ctx, req.(*CreateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertAdminApi_UpdateAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertAdminApiServer).UpdateAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertAdminApi/UpdateAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertAdminApiServer).UpdateAlert(ctx, req.(*UpdateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertAdminApi_DeleteAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertAdminApiServer).DeleteAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.AlertAdminApi/DeleteAlert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertAdminApiServer).DeleteAlert(ctx, req.(*DeleteAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertAdminApi_ServiceDesc is the grpc.ServiceDesc for AlertAdminApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertAdminApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.AlertAdminApi",
	HandlerType: (*AlertAdminApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlert",
			Handler:    _AlertAdminApi_CreateAlert_Handler,
		},
		{
			MethodName: "UpdateAlert",
			Handler:    _AlertAdminApi_UpdateAlert_Handler,
		},
		{
			MethodName: "DeleteAlert",
			Handler:    _AlertAdminApi_DeleteAlert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alerts.proto",
}
