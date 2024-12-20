// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapColorApi	adapts a ColorApiServer	and presents it as a ColorApiClient
func WrapColorApi(server ColorApiServer) *ColorApiWrapper {
	conn := wrap.ServerToClient(ColorApi_ServiceDesc, server)
	client := NewColorApiClient(conn)
	return &ColorApiWrapper{
		ColorApiClient: client,
		server:         server,
		conn:           conn,
		desc:           ColorApi_ServiceDesc,
	}
}

type ColorApiWrapper struct {
	ColorApiClient

	server ColorApiServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *ColorApiWrapper) UnwrapServer() ColorApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *ColorApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *ColorApiWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
