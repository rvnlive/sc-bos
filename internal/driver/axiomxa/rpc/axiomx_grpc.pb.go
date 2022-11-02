// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: internal/driver/axiomxa/rpc/axiomxa.proto

package rpc

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

// AxiomXDriverServiceClient is the client API for AxiomXDriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AxiomXDriverServiceClient interface {
	SaveQRCredential(ctx context.Context, in *SaveQRCredentialRequest, opts ...grpc.CallOption) (*SaveQRCredentialResponse, error)
}

type axiomXDriverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAxiomXDriverServiceClient(cc grpc.ClientConnInterface) AxiomXDriverServiceClient {
	return &axiomXDriverServiceClient{cc}
}

func (c *axiomXDriverServiceClient) SaveQRCredential(ctx context.Context, in *SaveQRCredentialRequest, opts ...grpc.CallOption) (*SaveQRCredentialResponse, error) {
	out := new(SaveQRCredentialResponse)
	err := c.cc.Invoke(ctx, "/vanti.bsp.ew.driver.axiomxa.AxiomXDriverService/SaveQRCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AxiomXDriverServiceServer is the server API for AxiomXDriverService service.
// All implementations must embed UnimplementedAxiomXDriverServiceServer
// for forward compatibility
type AxiomXDriverServiceServer interface {
	SaveQRCredential(context.Context, *SaveQRCredentialRequest) (*SaveQRCredentialResponse, error)
	mustEmbedUnimplementedAxiomXDriverServiceServer()
}

// UnimplementedAxiomXDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAxiomXDriverServiceServer struct {
}

func (UnimplementedAxiomXDriverServiceServer) SaveQRCredential(context.Context, *SaveQRCredentialRequest) (*SaveQRCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQRCredential not implemented")
}
func (UnimplementedAxiomXDriverServiceServer) mustEmbedUnimplementedAxiomXDriverServiceServer() {}

// UnsafeAxiomXDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AxiomXDriverServiceServer will
// result in compilation errors.
type UnsafeAxiomXDriverServiceServer interface {
	mustEmbedUnimplementedAxiomXDriverServiceServer()
}

func RegisterAxiomXDriverServiceServer(s grpc.ServiceRegistrar, srv AxiomXDriverServiceServer) {
	s.RegisterService(&AxiomXDriverService_ServiceDesc, srv)
}

func _AxiomXDriverService_SaveQRCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQRCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxiomXDriverServiceServer).SaveQRCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vanti.bsp.ew.driver.axiomxa.AxiomXDriverService/SaveQRCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxiomXDriverServiceServer).SaveQRCredential(ctx, req.(*SaveQRCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AxiomXDriverService_ServiceDesc is the grpc.ServiceDesc for AxiomXDriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AxiomXDriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vanti.bsp.ew.driver.axiomxa.AxiomXDriverService",
	HandlerType: (*AxiomXDriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQRCredential",
			Handler:    _AxiomXDriverService_SaveQRCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/driver/axiomxa/rpc/axiomxa.proto",
}
