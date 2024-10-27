package config

import (
	"github.com/spf13/viper"
)

type MqttConfigTopic struct {
	Name string `mapstructure:"name"`
}

type MqttConfig struct {
	BrokerUrl string            `mapstructure:"broker_url"`
	UserName  string            `mapstructure:"username"`
	Password  string            `mapstructure:"password"`
	Topics    []MqttConfigTopic `mapstructure:"topics"`
}

type Config struct {
	Mqtt MqttConfig `mapstructure:"mqtt"`
}

func New(configFilePath string) (*Config, error) {
	config := Config{}

	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
