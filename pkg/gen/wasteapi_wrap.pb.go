// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapWasteApi	adapts a WasteApiServer	and presents it as a WasteApiClient
func WrapWasteApi(server WasteApiServer) *WasteApiWrapper {
	conn := wrap.ServerToClient(WasteApi_ServiceDesc, server)
	client := NewWasteApiClient(conn)
	return &WasteApiWrapper{
		WasteApiClient: client,
		server:         server,
		conn:           conn,
		desc:           WasteApi_ServiceDesc,
	}
}

type WasteApiWrapper struct {
	WasteApiClient

	server WasteApiServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *WasteApiWrapper) UnwrapServer() WasteApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *WasteApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *WasteApiWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
