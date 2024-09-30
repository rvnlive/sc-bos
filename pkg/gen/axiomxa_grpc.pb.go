// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: axiomxa.proto

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

// AxiomXaDriverServiceClient is the client API for AxiomXaDriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AxiomXaDriverServiceClient interface {
	SaveQRCredential(ctx context.Context, in *SaveQRCredentialRequest, opts ...grpc.CallOption) (*SaveQRCredentialResponse, error)
}

type axiomXaDriverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAxiomXaDriverServiceClient(cc grpc.ClientConnInterface) AxiomXaDriverServiceClient {
	return &axiomXaDriverServiceClient{cc}
}

func (c *axiomXaDriverServiceClient) SaveQRCredential(ctx context.Context, in *SaveQRCredentialRequest, opts ...grpc.CallOption) (*SaveQRCredentialResponse, error) {
	out := new(SaveQRCredentialResponse)
	err := c.cc.Invoke(ctx, "/smartcore.bos.driver.axiomxa.AxiomXaDriverService/SaveQRCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AxiomXaDriverServiceServer is the server API for AxiomXaDriverService service.
// All implementations must embed UnimplementedAxiomXaDriverServiceServer
// for forward compatibility
type AxiomXaDriverServiceServer interface {
	SaveQRCredential(context.Context, *SaveQRCredentialRequest) (*SaveQRCredentialResponse, error)
	mustEmbedUnimplementedAxiomXaDriverServiceServer()
}

// UnimplementedAxiomXaDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAxiomXaDriverServiceServer struct {
}

func (UnimplementedAxiomXaDriverServiceServer) SaveQRCredential(context.Context, *SaveQRCredentialRequest) (*SaveQRCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQRCredential not implemented")
}
func (UnimplementedAxiomXaDriverServiceServer) mustEmbedUnimplementedAxiomXaDriverServiceServer() {}

// UnsafeAxiomXaDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AxiomXaDriverServiceServer will
// result in compilation errors.
type UnsafeAxiomXaDriverServiceServer interface {
	mustEmbedUnimplementedAxiomXaDriverServiceServer()
}

func RegisterAxiomXaDriverServiceServer(s grpc.ServiceRegistrar, srv AxiomXaDriverServiceServer) {
	s.RegisterService(&AxiomXaDriverService_ServiceDesc, srv)
}

func _AxiomXaDriverService_SaveQRCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQRCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AxiomXaDriverServiceServer).SaveQRCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smartcore.bos.driver.axiomxa.AxiomXaDriverService/SaveQRCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AxiomXaDriverServiceServer).SaveQRCredential(ctx, req.(*SaveQRCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AxiomXaDriverService_ServiceDesc is the grpc.ServiceDesc for AxiomXaDriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AxiomXaDriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smartcore.bos.driver.axiomxa.AxiomXaDriverService",
	HandlerType: (*AxiomXaDriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQRCredential",
			Handler:    _AxiomXaDriverService_SaveQRCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axiomxa.proto",
}
