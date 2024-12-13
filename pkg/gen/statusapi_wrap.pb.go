// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapStatusApi	adapts a StatusApiServer	and presents it as a StatusApiClient
func WrapStatusApi(server StatusApiServer) *StatusApiWrapper {
	conn := wrap.ServerToClient(StatusApi_ServiceDesc, server)
	client := NewStatusApiClient(conn)
	return &StatusApiWrapper{
		StatusApiClient: client,
		server:          server,
		conn:            conn,
		desc:            StatusApi_ServiceDesc,
	}
}

type StatusApiWrapper struct {
	StatusApiClient

	server StatusApiServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *StatusApiWrapper) UnwrapServer() StatusApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *StatusApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *StatusApiWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
