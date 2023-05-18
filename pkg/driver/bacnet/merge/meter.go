package merge

import (
	"context"
	"encoding/json"

	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/smart-core-os/sc-golang/pkg/resource"
	"github.com/vanti-dev/gobacnet"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/config"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/known"
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/gentrait/meter"
	"github.com/vanti-dev/sc-bos/pkg/gentrait/statuspb"
	"github.com/vanti-dev/sc-bos/pkg/node"
	"github.com/vanti-dev/sc-bos/pkg/task"
)

type meterConfig struct {
	config.Trait
	Usage *config.ValueSource `json:"usage,omitempty"`
}

func readMeterConfig(raw []byte) (cfg meterConfig, err error) {
	err = json.Unmarshal(raw, &cfg)
	return
}

type meterTrait struct {
	client   *gobacnet.Client
	known    known.Context
	statuses *statuspb.Map
	logger   *zap.Logger

	model *meter.Model
	*meter.ModelServer
	config   meterConfig
	pollTask *task.Intermittent
}

func newMeter(client *gobacnet.Client, devices known.Context, statuses *statuspb.Map, config config.RawTrait, logger *zap.Logger) (*meterTrait, error) {
	cfg, err := readMeterConfig(config.Raw)
	if err != nil {
		return nil, err
	}
	model := meter.NewModel(resource.WithNoDuplicates())
	t := &meterTrait{
		client:      client,
		known:       devices,
		statuses:    statuses,
		logger:      logger,
		model:       model,
		ModelServer: meter.NewModelServer(model),
		config:      cfg,
	}
	t.pollTask = task.NewIntermittent(t.startPoll)
	return t, nil
}

func (t *meterTrait) AnnounceSelf(a node.Announcer) node.Undo {
	return a.Announce(t.config.Name, node.HasTrait(meter.TraitName, node.WithClients(gen.WrapMeterApi(t))))
}

func (t *meterTrait) GetMeterReading(ctx context.Context, request *gen.GetMeterReadingRequest) (*gen.MeterReading, error) {
	_, err := t.pollPeer(ctx)
	if err != nil {
		return nil, err
	}
	return t.ModelServer.GetMeterReading(ctx, request)
}

func (t *meterTrait) PullMeterReadings(request *gen.PullMeterReadingsRequest, server gen.MeterApi_PullMeterReadingsServer) error {
	err := t.pollTask.Attach(server.Context())
	if err != nil {
		return err
	}
	return t.ModelServer.PullMeterReadings(request, server)
}

func (t *meterTrait) startPoll(init context.Context) (stop task.StopFn, err error) {
	return startPoll(init, "meter", t.config.PollPeriodDuration(), t.config.PollTimeoutDuration(), t.logger, func(ctx context.Context) error {
		_, err := t.pollPeer(ctx)
		return err
	})
}

func (t *meterTrait) pollPeer(ctx context.Context) (*gen.MeterReading, error) {
	responses := readProperties(ctx, t.client, t.known, *t.config.Usage)
	var errs []error
	usage, err := float32Value(responses[0])
	if err != nil {
		errs = append(errs, ErrReadProperty{Prop: "usage", Cause: err})
	}
	updatePollErrorStatus(t.statuses, t.config.Name, 1, errs...)
	if len(errs) > 0 {
		return nil, multierr.Combine(errs...)
	}
	data := &gen.MeterReading{
		Usage: usage,
	}
	return t.model.UpdateMeterReading(data)
}
