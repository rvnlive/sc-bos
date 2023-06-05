package config

import (
	"github.com/vanti-dev/sc-bos/pkg/zone"
)

type Root struct {
	zone.Config

	// 'All' meaning all devices that are mentioned elsewhere in the config.
	StatusLogAll bool     `json:"statusLogAll,omitempty"`
	StatusLogs   []string `json:"statusLogs,omitempty"`
}
