package config

import "github.com/ghtix/gomodo/internal/config"

type ControlledServiceConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type ControlConfig struct {
	Service            config.ServiceConfig      `mapstructure:"service"`
	ControlledServices []ControlledServiceConfig `mapstructure:"controlled_services"`
}
