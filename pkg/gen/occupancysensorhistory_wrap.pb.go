// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapOccupancySensorHistory	adapts a OccupancySensorHistoryServer	and presents it as a OccupancySensorHistoryClient
func WrapOccupancySensorHistory(server OccupancySensorHistoryServer) *OccupancySensorHistoryWrapper {
	conn := wrap.ServerToClient(OccupancySensorHistory_ServiceDesc, server)
	client := NewOccupancySensorHistoryClient(conn)
	return &OccupancySensorHistoryWrapper{
		OccupancySensorHistoryClient: client,
		server:                       server,
		conn:                         conn,
		desc:                         OccupancySensorHistory_ServiceDesc,
	}
}

type OccupancySensorHistoryWrapper struct {
	OccupancySensorHistoryClient

	server OccupancySensorHistoryServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *OccupancySensorHistoryWrapper) UnwrapServer() OccupancySensorHistoryServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *OccupancySensorHistoryWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *OccupancySensorHistoryWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
