package config

import (
	"github.com/ghtix/gomodo/internal/config"
	"github.com/ghtix/gomodo/pkg/overkiz"
)

type OverkizServiceConfig struct {
	Service config.ServiceConfig  `mapstructure:"service"`
	Overkiz overkiz.OverkizConfig `mapstructure:"overkiz"`
}
