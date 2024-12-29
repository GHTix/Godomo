package mqttClient

type MqttConfigTopic struct {
	Name string `mapstructure:"name"`
}

type MqttConfig struct {
	BrokerUrl string            `mapstructure:"broker_url"`
	UserName  string            `mapstructure:"username"`
	Password  string            `mapstructure:"password"`
	Topics    []MqttConfigTopic `mapstructure:"topics"`
}
