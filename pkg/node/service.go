package node

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/smart-core-os/sc-golang/pkg/wrap"
	"github.com/vanti-dev/sc-bos/internal/router"
)

type Service struct {
	routerService *router.Service
	conn          grpc.ClientConnInterface
}

func (s *Service) NameRoutable() bool {
	return s.routerService.KeyRoutable()
}

// RemoteService returns a Service that routes requests to a ClientConn.
func RemoteService(desc protoreflect.ServiceDescriptor, conn grpc.ClientConnInterface) *Service {
	return &Service{
		routerService: maybeNameRoutedService(desc),
		conn:          conn,
	}
}

// LocalService returns a Service that routes requests to a local server implementation srv.
//
// The service described by desc must exist in the global protobuf registry, or an error will be returned.
func LocalService(desc grpc.ServiceDesc, srv any) (*Service, error) {
	reflectDesc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(desc.ServiceName))
	if err != nil {
		return nil, err
	}
	reflectSrvDesc, ok := reflectDesc.(protoreflect.ServiceDescriptor)
	if !ok {
		return nil, err
	}
	conn := wrap.ServerToClient(desc, srv)

	return &Service{
		routerService: maybeNameRoutedService(reflectSrvDesc),
		conn:          conn,
	}, nil
}

// AnnounceService will make the service available on the node's router.
//
// If the gRPC service is already registered by a device, then requests that can't be routed to a device will be routed
// to the provided Service implementation.
// Otherwise, all requests for the service will be routed to the provided Service implementation.
func (n *Node) AnnounceService(srv *Service) (Undo, error) {
	serviceName := srv.routerService.Name()
	// service must be present in the router first
	err := n.SupportService(srv)
	if err != nil {
		return NilUndo, err
	}

	err = n.router.AddRoute(serviceName, "", srv.conn)
	if err != nil {
		return NilUndo, err
	}
	return func() {
		_ = n.router.DeleteRoute(serviceName, "")
	}, nil
}

// SupportService enables the node's router to handle requests for the service.
//
// No routes are added, so by default all requests will fail.
func (n *Node) SupportService(srv *Service) error {
	serviceName := srv.routerService.Name()
	if n.router.GetService(serviceName) == nil {
		err := n.router.AddService(srv.routerService)
		if err != nil {
			return err
		}
	}
	return nil
}

func maybeNameRoutedService(desc protoreflect.ServiceDescriptor) *router.Service {
	srv, err := router.NewRoutedService(desc, "name")
	if err == nil {
		return srv
	}
	return router.NewUnroutedService(desc)
}
