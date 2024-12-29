package config

import (
	"github.com/ghtix/gomodo/internal/config"
	mqttClient "github.com/ghtix/gomodo/internal/mqtt"
)

type AqaraServiceConfig struct {
	Service config.ServiceConfig  `mapstructure:"service"`
	Mqtt    mqttClient.MqttConfig `mapstructure:"mqtt"`
}