// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// ElectricHistoryRouter is a ElectricHistoryServer that allows routing named requests to specific ElectricHistoryClient
type ElectricHistoryRouter struct {
	UnimplementedElectricHistoryServer

	router.Router
}

// compile time check that we implement the interface we need
var _ ElectricHistoryServer = (*ElectricHistoryRouter)(nil)

func NewElectricHistoryRouter(opts ...router.Option) *ElectricHistoryRouter {
	return &ElectricHistoryRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithElectricHistoryClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithElectricHistoryClientFactory(f func(name string) (ElectricHistoryClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *ElectricHistoryRouter) Register(server *grpc.Server) {
	RegisterElectricHistoryServer(server, r)
}

// Add extends Router.Add to panic if client is not of type ElectricHistoryClient.
func (r *ElectricHistoryRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a ElectricHistoryClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *ElectricHistoryRouter) HoldsType(client any) bool {
	_, ok := client.(ElectricHistoryClient)
	return ok
}

func (r *ElectricHistoryRouter) AddElectricHistoryClient(name string, client ElectricHistoryClient) ElectricHistoryClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(ElectricHistoryClient)
}

func (r *ElectricHistoryRouter) RemoveElectricHistoryClient(name string) ElectricHistoryClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(ElectricHistoryClient)
}

func (r *ElectricHistoryRouter) GetElectricHistoryClient(name string) (ElectricHistoryClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ElectricHistoryClient), nil
}

func (r *ElectricHistoryRouter) ListElectricDemandHistory(ctx context.Context, request *ListElectricDemandHistoryRequest) (*ListElectricDemandHistoryResponse, error) {
	child, err := r.GetElectricHistoryClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.ListElectricDemandHistory(ctx, request)
}
