// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapMeterInfo	adapts a MeterInfoServer	and presents it as a MeterInfoClient
func WrapMeterInfo(server MeterInfoServer) *MeterInfoWrapper {
	conn := wrap.ServerToClient(MeterInfo_ServiceDesc, server)
	client := NewMeterInfoClient(conn)
	return &MeterInfoWrapper{
		MeterInfoClient: client,
		server:          server,
		conn:            conn,
		desc:            MeterInfo_ServiceDesc,
	}
}

type MeterInfoWrapper struct {
	MeterInfoClient

	server MeterInfoServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *MeterInfoWrapper) UnwrapServer() MeterInfoServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *MeterInfoWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *MeterInfoWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
