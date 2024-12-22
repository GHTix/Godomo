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

type OverkizConfig struct {
	BaseUrl            string `mapstructure:"base_url"`
	UserName           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	OAuthLoginEndpoint string `mapstructure:"oauth_login_endpoint"`
	OAuthClientId      string `mapstructure:"oauth_client_id"`
	OAuthClientSecret  string `mapstructure:"oauth_client_secret"`
}

type Config struct {
	Mqtt    MqttConfig    `mapstructure:"mqtt"`
	Overkiz OverkizConfig `mapstructure:"overkiz"`
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
