// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"

	"github.com/smart-core-os/sc-golang/pkg/router"
)

// TemperatureApiRouter is a gen.TemperatureApiServer that allows routing named requests to specific gen.TemperatureApiClient
type TemperatureApiRouter struct {
	UnimplementedTemperatureApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ TemperatureApiServer = (*TemperatureApiRouter)(nil)

func NewTemperatureApiRouter(opts ...router.Option) *TemperatureApiRouter {
	return &TemperatureApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithTemperatureApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithTemperatureApiClientFactory(f func(name string) (TemperatureApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *TemperatureApiRouter) Register(server grpc.ServiceRegistrar) {
	RegisterTemperatureApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type gen.TemperatureApiClient.
func (r *TemperatureApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a gen.TemperatureApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *TemperatureApiRouter) HoldsType(client any) bool {
	_, ok := client.(TemperatureApiClient)
	return ok
}

func (r *TemperatureApiRouter) AddTemperatureApiClient(name string, client TemperatureApiClient) TemperatureApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(TemperatureApiClient)
}

func (r *TemperatureApiRouter) RemoveTemperatureApiClient(name string) TemperatureApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(TemperatureApiClient)
}

func (r *TemperatureApiRouter) GetTemperatureApiClient(name string) (TemperatureApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TemperatureApiClient), nil
}

func (r *TemperatureApiRouter) GetTemperature(ctx context.Context, request *GetTemperatureRequest) (*Temperature, error) {
	child, err := r.GetTemperatureApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetTemperature(ctx, request)
}

func (r *TemperatureApiRouter) PullTemperature(request *PullTemperatureRequest, server TemperatureApi_PullTemperatureServer) error {
	child, err := r.GetTemperatureApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullTemperature(reqCtx, request)
	if err != nil {
		return err
	}

	// send the stream header
	header, err := stream.Header()
	if err != nil {
		return err
	}
	if err = server.SendHeader(header); err != nil {
		return err
	}

	// send all the messages
	// false means the error is from the child, true means the error is from the caller
	var callerError bool
	for {
		// Impl note: we could improve throughput here by issuing the Recv and Send in different goroutines, but we're doing
		// it synchronously until we have a need to change the behaviour

		var msg *PullTemperatureResponse
		msg, err = stream.Recv()
		if err != nil {
			break
		}

		err = server.Send(msg)
		if err != nil {
			callerError = true
			break
		}
	}

	// err is guaranteed to be non-nil as it's the only way to exit the loop
	if callerError {
		// cancel the request
		reqDone()
		return err
	} else {
		if trailer := stream.Trailer(); trailer != nil {
			server.SetTrailer(trailer)
		}
		if err == io.EOF {
			return nil
		}
		return err
	}
}

func (r *TemperatureApiRouter) UpdateTemperature(ctx context.Context, request *UpdateTemperatureRequest) (*Temperature, error) {
	child, err := r.GetTemperatureApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.UpdateTemperature(ctx, request)
}
