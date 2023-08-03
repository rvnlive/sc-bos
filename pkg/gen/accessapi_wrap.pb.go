// Code generated by protoc-gen-wrapper. DO NOT EDIT.

package gen

import (
	context "context"
	wrap "github.com/smart-core-os/sc-golang/pkg/wrap"
	grpc "google.golang.org/grpc"
)

// WrapAccessApi	adapts a AccessApiServer	and presents it as a AccessApiClient
func WrapAccessApi(server AccessApiServer) AccessApiClient {
	return &accessApiWrapper{server}
}

type accessApiWrapper struct {
	server AccessApiServer
}

// compile time check that we implement the interface we need
var _ AccessApiClient = (*accessApiWrapper)(nil)

// UnwrapServer returns the underlying server instance.
func (w *accessApiWrapper) UnwrapServer() AccessApiServer {
	return w.server
}

// Unwrap implements wrap.Unwrapper and returns the underlying server instance as an unknown type.
func (w *accessApiWrapper) Unwrap() any {
	return w.UnwrapServer()
}

func (w *accessApiWrapper) GetLastAccessAttempt(ctx context.Context, req *GetLastAccessAttemptRequest, _ ...grpc.CallOption) (*AccessAttempt, error) {
	return w.server.GetLastAccessAttempt(ctx, req)
}

func (w *accessApiWrapper) PullAccessAttempts(ctx context.Context, in *PullAccessAttemptsRequest, opts ...grpc.CallOption) (AccessApi_PullAccessAttemptsClient, error) {
	stream := wrap.NewClientServerStream(ctx)
	server := &pullAccessAttemptsAccessApiServerWrapper{stream.Server()}
	client := &pullAccessAttemptsAccessApiClientWrapper{stream.Client()}
	go func() {
		err := w.server.PullAccessAttempts(in, server)
		stream.Close(err)
	}()
	return client, nil
}

type pullAccessAttemptsAccessApiClientWrapper struct {
	grpc.ClientStream
}

func (w *pullAccessAttemptsAccessApiClientWrapper) Recv() (*PullAccessAttemptsResponse, error) {
	m := new(PullAccessAttemptsResponse)
	if err := w.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type pullAccessAttemptsAccessApiServerWrapper struct {
	grpc.ServerStream
}

func (s *pullAccessAttemptsAccessApiServerWrapper) Send(response *PullAccessAttemptsResponse) error {
	return s.ServerStream.SendMsg(response)
}
