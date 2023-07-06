package merge

import (
	"context"
	"errors"
	"fmt"

	"github.com/vanti-dev/gobacnet"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/comm"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/config"
	"github.com/vanti-dev/sc-bos/pkg/driver/bacnet/known"
)

func ptr[T any](v T, err error) (*T, error) {
	return &v, err
}

func valuesEquivalent(a, b any) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func readPropertyFloat32(ctx context.Context, client *gobacnet.Client, known known.Context, value config.ValueSource) (float32, error) {
	data, err := comm.ReadProperty(ctx, client, known, value)
	if err != nil {
		return 0, err
	}
	return comm.Float32Value(data)
}

var (
	ErrTraitNotSupported = errors.New("trait not supported")
)
