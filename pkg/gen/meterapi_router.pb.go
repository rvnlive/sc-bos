// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
	io "io"
)

// MeterApiRouter is a gen.MeterApiServer that allows routing named requests to specific gen.MeterApiClient
type MeterApiRouter struct {
	UnimplementedMeterApiServer

	router.Router
}

// compile time check that we implement the interface we need
var _ MeterApiServer = (*MeterApiRouter)(nil)

func NewMeterApiRouter(opts ...router.Option) *MeterApiRouter {
	return &MeterApiRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithMeterApiClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithMeterApiClientFactory(f func(name string) (MeterApiClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *MeterApiRouter) Register(server grpc.ServiceRegistrar) {
	RegisterMeterApiServer(server, r)
}

// Add extends Router.Add to panic if client is not of type gen.MeterApiClient.
func (r *MeterApiRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a gen.MeterApiClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *MeterApiRouter) HoldsType(client any) bool {
	_, ok := client.(MeterApiClient)
	return ok
}

func (r *MeterApiRouter) AddMeterApiClient(name string, client MeterApiClient) MeterApiClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(MeterApiClient)
}

func (r *MeterApiRouter) RemoveMeterApiClient(name string) MeterApiClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(MeterApiClient)
}

func (r *MeterApiRouter) GetMeterApiClient(name string) (MeterApiClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(MeterApiClient), nil
}

func (r *MeterApiRouter) GetMeterReading(ctx context.Context, request *GetMeterReadingRequest) (*MeterReading, error) {
	child, err := r.GetMeterApiClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.GetMeterReading(ctx, request)
}

func (r *MeterApiRouter) PullMeterReadings(request *PullMeterReadingsRequest, server MeterApi_PullMeterReadingsServer) error {
	child, err := r.GetMeterApiClient(request.Name)
	if err != nil {
		return err
	}

	// so we can cancel our forwarding request if we can't send responses to our caller
	reqCtx, reqDone := context.WithCancel(server.Context())
	// issue the request
	stream, err := child.PullMeterReadings(reqCtx, request)
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

		var msg *PullMeterReadingsResponse
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
