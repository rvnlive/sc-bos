package history

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/smart-core-os/sc-api/go/traits"
	"github.com/vanti-dev/sc-bos/internal/util/pull"
)

func (a *automation) collectElectricDemandChanges(ctx context.Context, name string, payloads chan<- []byte) {
	var client traits.ElectricApiClient
	if err := a.clients.Client(&client); err != nil {
		a.logger.Error("collection aborted", zap.Error(err))
		return
	}

	pullFn := func(ctx context.Context, changes chan<- []byte) error {
		stream, err := client.PullDemand(ctx, &traits.PullDemandRequest{Name: name, UpdatesOnly: true})
		if err != nil {
			return err
		}
		for {
			msg, err := stream.Recv()
			if err != nil {
				return err
			}
			for _, change := range msg.Changes {
				payload, err := proto.Marshal(change.Demand)
				if err != nil {
					return err
				}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case changes <- payload:
				}
			}
		}
	}
	pollFn := func(ctx context.Context, changes chan<- []byte) error {
		demand, err := client.GetDemand(ctx, &traits.GetDemandRequest{Name: name})
		if err != nil {
			return err
		}
		payload, err := proto.Marshal(demand)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case changes <- payload:
		}
		return nil
	}

	err := pull.Changes(ctx, pull.NewFetcher(pullFn, pollFn), payloads, pull.WithLogger(a.logger))
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}
	if err != nil {
		a.logger.Warn("collection aborted", zap.Error(err))
	}
}
