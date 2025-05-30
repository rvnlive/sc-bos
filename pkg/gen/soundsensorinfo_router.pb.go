// Code generated by protoc-gen-router. DO NOT EDIT.

package gen

import (
	context "context"
	fmt "fmt"
	router "github.com/smart-core-os/sc-golang/pkg/router"
	grpc "google.golang.org/grpc"
)

// SoundSensorInfoRouter is a SoundSensorInfoServer that allows routing named requests to specific SoundSensorInfoClient
type SoundSensorInfoRouter struct {
	UnimplementedSoundSensorInfoServer

	router.Router
}

// compile time check that we implement the interface we need
var _ SoundSensorInfoServer = (*SoundSensorInfoRouter)(nil)

func NewSoundSensorInfoRouter(opts ...router.Option) *SoundSensorInfoRouter {
	return &SoundSensorInfoRouter{
		Router: router.NewRouter(opts...),
	}
}

// WithSoundSensorInfoClientFactory instructs the router to create a new
// client the first time Get is called for that name.
func WithSoundSensorInfoClientFactory(f func(name string) (SoundSensorInfoClient, error)) router.Option {
	return router.WithFactory(func(name string) (any, error) {
		return f(name)
	})
}

func (r *SoundSensorInfoRouter) Register(server grpc.ServiceRegistrar) {
	RegisterSoundSensorInfoServer(server, r)
}

// Add extends Router.Add to panic if client is not of type SoundSensorInfoClient.
func (r *SoundSensorInfoRouter) Add(name string, client any) any {
	if !r.HoldsType(client) {
		panic(fmt.Sprintf("not correct type: client of type %T is not a SoundSensorInfoClient", client))
	}
	return r.Router.Add(name, client)
}

func (r *SoundSensorInfoRouter) HoldsType(client any) bool {
	_, ok := client.(SoundSensorInfoClient)
	return ok
}

func (r *SoundSensorInfoRouter) AddSoundSensorInfoClient(name string, client SoundSensorInfoClient) SoundSensorInfoClient {
	res := r.Add(name, client)
	if res == nil {
		return nil
	}
	return res.(SoundSensorInfoClient)
}

func (r *SoundSensorInfoRouter) RemoveSoundSensorInfoClient(name string) SoundSensorInfoClient {
	res := r.Remove(name)
	if res == nil {
		return nil
	}
	return res.(SoundSensorInfoClient)
}

func (r *SoundSensorInfoRouter) GetSoundSensorInfoClient(name string) (SoundSensorInfoClient, error) {
	res, err := r.Get(name)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SoundSensorInfoClient), nil
}

func (r *SoundSensorInfoRouter) DescribeSoundLevel(ctx context.Context, request *DescribeSoundLevelRequest) (*SoundLevelSupport, error) {
	child, err := r.GetSoundSensorInfoClient(request.Name)
	if err != nil {
		return nil, err
	}

	return child.DescribeSoundLevel(ctx, request)
}
