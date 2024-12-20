// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// WasteInfoRouter is a WasteInfoServer that allows routing named requests to specific WasteInfoClient
type WasteInfoRouter struct {
	UnimplementedWasteInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ WasteInfoServer = (*WasteInfoRouter)(nil)

func NewWasteInfoRouter(opts ...router.Option) *WasteInfoRouter {
	return &WasteInfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithWasteInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithWasteInfoClientFactory(f func(name string) (WasteInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *WasteInfoRouter) Register(server grpc.ServiceRegistrar) {
	RegisterWasteInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type WasteInfoClient.
func (r *WasteInfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a WasteInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *WasteInfoRouter) HoldsType(client any) bool {
	_, ok := client.(WasteInfoClient)
	return ok
}

func (r *WasteInfoRouter) AddWasteInfoClient(name string, client WasteInfoClient) WasteInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(WasteInfoClient)
}

func (r *WasteInfoRouter) RemoveWasteInfoClient(name string) WasteInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(WasteInfoClient)
}

func (r *WasteInfoRouter) GetWasteInfoClient(name string) (WasteInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(WasteInfoClient), nil
}

func (r *WasteInfoRouter) DescribeWasteRecord(ctx context.Context, request *DescribeWasteRecordRequest) (*WasteRecordSupport, error) {
	child, err := r.GetWasteInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeWasteRecord(ctx, request)
}
