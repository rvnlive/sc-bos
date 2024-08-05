// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	context "context"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapWasteApi	adapts a WasteApiServer	and presents it as a WasteApiClient
func WrapWasteApi(server WasteApiServer) WasteApiClient {
	return &wasteApiWrapper{server}
}

type wasteApiWrapper struct {
	server WasteApiServer
}

// compile time check that we implement the interface we need
var _ WasteApiClient = (*wasteApiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *wasteApiWrapper) UnwrapServer() WasteApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *wasteApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *wasteApiWrapper) ListWasteRecords(ctx context.Context, req *ListWasteRecordsRequest, _ ...grpc.CallOption) (*ListWasteRecordsResponse, error) {
	return w.server.ListWasteRecords(ctx, req)
}

func (w *wasteApiWrapper) PullWasteRecords(ctx context.Context, in *PullWasteRecordsRequest, opts ...grpc.CallOption) (WasteApi_PullWasteRecordsClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullWasteRecordsWasteApiServerWrapper{stream.Server()}
	client := &pullWasteRecordsWasteApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullWasteRecords(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullWasteRecordsWasteApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullWasteRecordsWasteApiClientWrapper) Recv() (*PullWasteRecordsResponse, error) {
	m := new(PullWasteRecordsResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullWasteRecordsWasteApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullWasteRecordsWasteApiServerWrapper) Send(response *PullWasteRecordsResponse) error {
	return s.ServerStream.SendMsg(response)
}