package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	ListeningAddress string `mapstructure:"listening"`
}

func New[T any](configFilePath string) (*T, error) {
	var config T

	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config file %s. %w", configFilePath, err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file %s. %w", configFilePath, err)
	}

	return &config, nil
}
