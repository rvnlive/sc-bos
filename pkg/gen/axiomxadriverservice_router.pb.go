// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// AxiomXaDriverServiceRouter is a AxiomXaDriverServiceServer that allows routing named requests to specific AxiomXaDriverServiceClient
type AxiomXaDriverServiceRouter struct {
	UnimplementedAxiomXaDriverServiceServer

	router.Router
}

// compile time check that we implement the interface we need
var _ AxiomXaDriverServiceServer = (*AxiomXaDriverServiceRouter)(nil)

func NewAxiomXaDriverServiceRouter(opts ...router.Option) *AxiomXaDriverServiceRouter {
	return &AxiomXaDriverServiceRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithAxiomXaDriverServiceClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithAxiomXaDriverServiceClientFactory(f func(name string) (AxiomXaDriverServiceClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *AxiomXaDriverServiceRouter) Register(server grpc.ServiceRegistrar) {
	RegisterAxiomXaDriverServiceServer(server, r)
}

// Add extends Router.Add to panic if client is not of type AxiomXaDriverServiceClient.
func (r *AxiomXaDriverServiceRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a AxiomXaDriverServiceClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *AxiomXaDriverServiceRouter) HoldsType(client any) bool {
	_, ok := client.(AxiomXaDriverServiceClient)
	return ok
}

func (r *AxiomXaDriverServiceRouter) AddAxiomXaDriverServiceClient(name string, client AxiomXaDriverServiceClient) AxiomXaDriverServiceClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(AxiomXaDriverServiceClient)
}

func (r *AxiomXaDriverServiceRouter) RemoveAxiomXaDriverServiceClient(name string) AxiomXaDriverServiceClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(AxiomXaDriverServiceClient)
}

func (r *AxiomXaDriverServiceRouter) GetAxiomXaDriverServiceClient(name string) (AxiomXaDriverServiceClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AxiomXaDriverServiceClient), nil
}

func (r *AxiomXaDriverServiceRouter) SaveQRCredential(ctx context.Context, request *SaveQRCredentialRequest) (*SaveQRCredentialResponse, error) {
	child, err := r.GetAxiomXaDriverServiceClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.SaveQRCredential(ctx, request)
}
