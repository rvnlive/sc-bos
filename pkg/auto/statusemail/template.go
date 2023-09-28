package statusemail

import (
	"github.com/vanti-dev/sc-bos/pkg/auto/statusemail/config"
	"github.com/vanti-dev/sc-bos/pkg/gen"
)

type Status struct {
	Sent, Read *gen.StatusLog
	Source     config.Source
}

type Attrs struct {
	WorseLogs  []*Status // logs whose status has got worse since the last send
	BetterLogs []*Status // logs whose status has got better since the last send
	SameLogs   []*Status // logs whose status has not changed since the last send
	BadLogs    []*Status // logs that have a bad status
	AllLogs    []*Status // all logs we have
}
