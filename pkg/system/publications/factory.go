package publications

import (
	"context"
	"errors"
	"fmt"

	"github.com/smart-core-os/sc-golang/pkg/trait"
	"github.com/smart-core-os/sc-golang/pkg/trait/publication"
	"github.com/vanti-dev/sc-bos/internal/util/pgxutil"
	"github.com/vanti-dev/sc-bos/pkg/node"
	"github.com/vanti-dev/sc-bos/pkg/system"
	"github.com/vanti-dev/sc-bos/pkg/system/publications/config"
	"github.com/vanti-dev/sc-bos/pkg/system/publications/pgxpublications"
	"github.com/vanti-dev/sc-bos/pkg/task/service"
	"go.uber.org/zap"
)

var Factory factory

type factory struct{}

func (_ factory) New(services system.Services) service.Lifecycle {
	return NewSystem(services)
}

func NewSystem(services system.Services) *System {
	s := &System{
		logger:    services.Logger.Named("publications"),
		name:      services.Node.Name(),
		announcer: services.Node,
	}
	s.Service = service.New(service.MonoApply(s.applyConfig))
	return s
}

type System struct {
	*service.Service[config.Root]
	logger *zap.Logger

	name      string
	announcer node.Announcer
}

func (s *System) applyConfig(ctx context.Context, cfg config.Root) error {
	if cfg.Storage == nil {
		return errors.New("no storage")
	}
	if cfg.Storage.Type != "postgres" {
		return fmt.Errorf("unsuported storage type %s, want one of [postgres]", cfg.Storage.Type)
	}

	pool, err := pgxutil.Connect(ctx, cfg.Storage.ConnectConfig)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	server, err := pgxpublications.NewServerFromPool(ctx, pool, pgxpublications.WithLogger(s.logger))
	if err != nil {
		return fmt.Errorf("init: %w", err)
	}

	// Note, ctx in cancelled each time config is updated (and on stop) because we use MonoApply in NewSystem
	announcer := node.AnnounceContext(ctx, s.announcer)
	announcer.Announce(s.name, node.HasTrait(trait.Publication, node.WithClients(publication.WrapApi(server))))

	return nil
}