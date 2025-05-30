// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapTenantApi	adapts a TenantApiServer	and presents it as a TenantApiClient
func WrapTenantApi(server TenantApiServer) *TenantApiWrapper {
	conn := wrap.ServerToClient(TenantApi_ServiceDesc, server)
	client := NewTenantApiClient(conn)
	return &TenantApiWrapper{
		TenantApiClient: client,
		server:          server,
		conn:            conn,
		desc:            TenantApi_ServiceDesc,
	}
}

type TenantApiWrapper struct {
	TenantApiClient

	server TenantApiServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *TenantApiWrapper) UnwrapServer() TenantApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *TenantApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *TenantApiWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
