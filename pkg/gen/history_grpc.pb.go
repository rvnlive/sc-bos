// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: history.proto

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

// HistoryAdminApiClient is the client API for HistoryAdminApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryAdminApiClient interface {
	CreateHistoryRecord(ctx context.Context, in *CreateHistoryRecordRequest, opts ...grpc.CallOption) (*HistoryRecord, error)
	ListHistoryRecords(ctx context.Context, in *ListHistoryRecordsRequest, opts ...grpc.CallOption) (*ListHistoryRecordsResponse, error)
}

type historyAdminApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryAdminApiClient(cc grpc.ClientConnInterface) HistoryAdminApiClient {
	return &historyAdminApiClient{cc}
}

func (c *historyAdminApiClient) CreateHistoryRecord(ctx context.Context, in *CreateHistoryRecordRequest, opts ...grpc.CallOption) (*HistoryRecord, error) {
	out := new(HistoryRecord)
	err := c.cc.Invoke(ctx, "/smartcore.bos.HistoryAdminApi/CreateHistoryRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyAdminApiClient) ListHistoryRecords(ctx context.Context, in *ListHistoryRecordsRequest, opts ...grpc.CallOption) (*ListHistoryRecordsResponse, error) {
	out := new(ListHistoryRecordsResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.HistoryAdminApi/ListHistoryRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryAdminApiServer is the server API for HistoryAdminApi service.
// All implementations must embed UnimplementedHistoryAdminApiServer
// for forward compatibility
type HistoryAdminApiServer interface {
	CreateHistoryRecord(context.Context, *CreateHistoryRecordRequest) (*HistoryRecord, error)
	ListHistoryRecords(context.Context, *ListHistoryRecordsRequest) (*ListHistoryRecordsResponse, error)
	mustEmbedUnimplementedHistoryAdminApiServer()
}

// UnimplementedHistoryAdminApiServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryAdminApiServer struct {
}

func (UnimplementedHistoryAdminApiServer) CreateHistoryRecord(context.Context, *CreateHistoryRecordRequest) (*HistoryRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistoryRecord not implemented")
}
func (UnimplementedHistoryAdminApiServer) ListHistoryRecords(context.Context, *ListHistoryRecordsRequest) (*ListHistoryRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHistoryRecords not implemented")
}
func (UnimplementedHistoryAdminApiServer) mustEmbedUnimplementedHistoryAdminApiServer() {}

// UnsafeHistoryAdminApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryAdminApiServer will
// result in compilation errors.
type UnsafeHistoryAdminApiServer interface {
	mustEmbedUnimplementedHistoryAdminApiServer()
}

func RegisterHistoryAdminApiServer(s grpc.ServiceRegistrar, srv HistoryAdminApiServer) {
	s.RegisterService(&HistoryAdminApi_ServiceDesc, srv)
}

func _HistoryAdminApi_CreateHistoryRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryAdminApiServer).CreateHistoryRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.HistoryAdminApi/CreateHistoryRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryAdminApiServer).CreateHistoryRecord(ctx, req.(*CreateHistoryRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryAdminApi_ListHistoryRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHistoryRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryAdminApiServer).ListHistoryRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.HistoryAdminApi/ListHistoryRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryAdminApiServer).ListHistoryRecords(ctx, req.(*ListHistoryRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HistoryAdminApi_ServiceDesc is the grpc.ServiceDesc for HistoryAdminApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoryAdminApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.HistoryAdminApi",
	HandlerType: (*HistoryAdminApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHistoryRecord",
			Handler:    _HistoryAdminApi_CreateHistoryRecord_Handler,
		},
		{
			MethodName: "ListHistoryRecords",
			Handler:    _HistoryAdminApi_ListHistoryRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}

// MeterHistoryClient is the client API for MeterHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeterHistoryClient interface {
	ListMeterReadingHistory(ctx context.Context, in *ListMeterReadingHistoryRequest, opts ...grpc.CallOption) (*ListMeterReadingHistoryResponse, error)
}

type meterHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewMeterHistoryClient(cc grpc.ClientConnInterface) MeterHistoryClient {
	return &meterHistoryClient{cc}
}

func (c *meterHistoryClient) ListMeterReadingHistory(ctx context.Context, in *ListMeterReadingHistoryRequest, opts ...grpc.CallOption) (*ListMeterReadingHistoryResponse, error) {
	out := new(ListMeterReadingHistoryResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.MeterHistory/ListMeterReadingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeterHistoryServer is the server API for MeterHistory service.
// All implementations must embed UnimplementedMeterHistoryServer
// for forward compatibility
type MeterHistoryServer interface {
	ListMeterReadingHistory(context.Context, *ListMeterReadingHistoryRequest) (*ListMeterReadingHistoryResponse, error)
	mustEmbedUnimplementedMeterHistoryServer()
}

// UnimplementedMeterHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedMeterHistoryServer struct {
}

func (UnimplementedMeterHistoryServer) ListMeterReadingHistory(context.Context, *ListMeterReadingHistoryRequest) (*ListMeterReadingHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeterReadingHistory not implemented")
}
func (UnimplementedMeterHistoryServer) mustEmbedUnimplementedMeterHistoryServer() {}

// UnsafeMeterHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeterHistoryServer will
// result in compilation errors.
type UnsafeMeterHistoryServer interface {
	mustEmbedUnimplementedMeterHistoryServer()
}

func RegisterMeterHistoryServer(s grpc.ServiceRegistrar, srv MeterHistoryServer) {
	s.RegisterService(&MeterHistory_ServiceDesc, srv)
}

func _MeterHistory_ListMeterReadingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMeterReadingHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeterHistoryServer).ListMeterReadingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.MeterHistory/ListMeterReadingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeterHistoryServer).ListMeterReadingHistory(ctx, req.(*ListMeterReadingHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MeterHistory_ServiceDesc is the grpc.ServiceDesc for MeterHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeterHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.MeterHistory",
	HandlerType: (*MeterHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMeterReadingHistory",
			Handler:    _MeterHistory_ListMeterReadingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}

// ElectricHistoryClient is the client API for ElectricHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectricHistoryClient interface {
	ListElectricDemandHistory(ctx context.Context, in *ListElectricDemandHistoryRequest, opts ...grpc.CallOption) (*ListElectricDemandHistoryResponse, error)
}

type electricHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewElectricHistoryClient(cc grpc.ClientConnInterface) ElectricHistoryClient {
	return &electricHistoryClient{cc}
}

func (c *electricHistoryClient) ListElectricDemandHistory(ctx context.Context, in *ListElectricDemandHistoryRequest, opts ...grpc.CallOption) (*ListElectricDemandHistoryResponse, error) {
	out := new(ListElectricDemandHistoryResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.ElectricHistory/ListElectricDemandHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectricHistoryServer is the server API for ElectricHistory service.
// All implementations must embed UnimplementedElectricHistoryServer
// for forward compatibility
type ElectricHistoryServer interface {
	ListElectricDemandHistory(context.Context, *ListElectricDemandHistoryRequest) (*ListElectricDemandHistoryResponse, error)
	mustEmbedUnimplementedElectricHistoryServer()
}

// UnimplementedElectricHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedElectricHistoryServer struct {
}

func (UnimplementedElectricHistoryServer) ListElectricDemandHistory(context.Context, *ListElectricDemandHistoryRequest) (*ListElectricDemandHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListElectricDemandHistory not implemented")
}
func (UnimplementedElectricHistoryServer) mustEmbedUnimplementedElectricHistoryServer() {}

// UnsafeElectricHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectricHistoryServer will
// result in compilation errors.
type UnsafeElectricHistoryServer interface {
	mustEmbedUnimplementedElectricHistoryServer()
}

func RegisterElectricHistoryServer(s grpc.ServiceRegistrar, srv ElectricHistoryServer) {
	s.RegisterService(&ElectricHistory_ServiceDesc, srv)
}

func _ElectricHistory_ListElectricDemandHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListElectricDemandHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectricHistoryServer).ListElectricDemandHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.ElectricHistory/ListElectricDemandHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectricHistoryServer).ListElectricDemandHistory(ctx, req.(*ListElectricDemandHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ElectricHistory_ServiceDesc is the grpc.ServiceDesc for ElectricHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElectricHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.ElectricHistory",
	HandlerType: (*ElectricHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListElectricDemandHistory",
			Handler:    _ElectricHistory_ListElectricDemandHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}

// OccupancySensorHistoryClient is the client API for OccupancySensorHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OccupancySensorHistoryClient interface {
	ListOccupancyHistory(ctx context.Context, in *ListOccupancyHistoryRequest, opts ...grpc.CallOption) (*ListOccupancyHistoryResponse, error)
}

type occupancySensorHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewOccupancySensorHistoryClient(cc grpc.ClientConnInterface) OccupancySensorHistoryClient {
	return &occupancySensorHistoryClient{cc}
}

func (c *occupancySensorHistoryClient) ListOccupancyHistory(ctx context.Context, in *ListOccupancyHistoryRequest, opts ...grpc.CallOption) (*ListOccupancyHistoryResponse, error) {
	out := new(ListOccupancyHistoryResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.OccupancySensorHistory/ListOccupancyHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OccupancySensorHistoryServer is the server API for OccupancySensorHistory service.
// All implementations must embed UnimplementedOccupancySensorHistoryServer
// for forward compatibility
type OccupancySensorHistoryServer interface {
	ListOccupancyHistory(context.Context, *ListOccupancyHistoryRequest) (*ListOccupancyHistoryResponse, error)
	mustEmbedUnimplementedOccupancySensorHistoryServer()
}

// UnimplementedOccupancySensorHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedOccupancySensorHistoryServer struct {
}

func (UnimplementedOccupancySensorHistoryServer) ListOccupancyHistory(context.Context, *ListOccupancyHistoryRequest) (*ListOccupancyHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOccupancyHistory not implemented")
}
func (UnimplementedOccupancySensorHistoryServer) mustEmbedUnimplementedOccupancySensorHistoryServer() {
}

// UnsafeOccupancySensorHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OccupancySensorHistoryServer will
// result in compilation errors.
type UnsafeOccupancySensorHistoryServer interface {
	mustEmbedUnimplementedOccupancySensorHistoryServer()
}

func RegisterOccupancySensorHistoryServer(s grpc.ServiceRegistrar, srv OccupancySensorHistoryServer) {
	s.RegisterService(&OccupancySensorHistory_ServiceDesc, srv)
}

func _OccupancySensorHistory_ListOccupancyHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOccupancyHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccupancySensorHistoryServer).ListOccupancyHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.OccupancySensorHistory/ListOccupancyHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccupancySensorHistoryServer).ListOccupancyHistory(ctx, req.(*ListOccupancyHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OccupancySensorHistory_ServiceDesc is the grpc.ServiceDesc for OccupancySensorHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OccupancySensorHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.OccupancySensorHistory",
	HandlerType: (*OccupancySensorHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOccupancyHistory",
			Handler:    _OccupancySensorHistory_ListOccupancyHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "history.proto",
}
