// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapAirQualitySensorHistory	adapts a gen.AirQualitySensorHistoryServer	and presents it as a gen.AirQualitySensorHistoryClient
func WrapAirQualitySensorHistory(server AirQualitySensorHistoryServer) *AirQualitySensorHistoryWrapper {
	conn := wrap.ServerToClient(AirQualitySensorHistory_ServiceDesc, server)
	client := NewAirQualitySensorHistoryClient(conn)
	return &AirQualitySensorHistoryWrapper{
		AirQualitySensorHistoryClient: client,
		server:                        server,
		conn:                          conn,
		desc:                          AirQualitySensorHistory_ServiceDesc,
	}
}

type AirQualitySensorHistoryWrapper struct {
	AirQualitySensorHistoryClient

	server AirQualitySensorHistoryServer
	conn   grpc.ClientConnInterface
	desc   grpc.ServiceDesc
}

// UnwrapServer returns the underlying server instance.
func (w *AirQualitySensorHistoryWrapper) UnwrapServer() AirQualitySensorHistoryServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *AirQualitySensorHistoryWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *AirQualitySensorHistoryWrapper) UnwrapService() (grpc.ClientConnInterface, grpc.ServiceDesc) {
	return w.conn, w.desc
}
