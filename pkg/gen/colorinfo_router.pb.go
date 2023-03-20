// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// ColorInfoRouter is a ColorInfoServer that allows routing named requests to specific ColorInfoClient
type ColorInfoRouter struct {
	UnimplementedColorInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ ColorInfoServer = (*ColorInfoRouter)(nil)

func NewColorInfoRouter(opts ...router.Option) *ColorInfoRouter {
	return &ColorInfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithColorInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithColorInfoClientFactory(f func(name string) (ColorInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ColorInfoRouter) Register(server *grpc.Server) {
	RegisterColorInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type ColorInfoClient.
func (r *ColorInfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a ColorInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ColorInfoRouter) HoldsType(client any) bool {
	_, ok := client.(ColorInfoClient)
	return ok
}

func (r *ColorInfoRouter) AddColorInfoClient(name string, client ColorInfoClient) ColorInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(ColorInfoClient)
}

func (r *ColorInfoRouter) RemoveColorInfoClient(name string) ColorInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(ColorInfoClient)
}

func (r *ColorInfoRouter) GetColorInfoClient(name string) (ColorInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ColorInfoClient), nil
}

func (r *ColorInfoRouter) DescribeColor(ctx context.Context, request *DescribeColorRequest) (*ColorSupport, error) {
	child, err := r.GetColorInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeColor(ctx, request)
}