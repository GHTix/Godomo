package mosquitto

import "net/url"

// const (
// 	DefaultMosquittoUrl      = "mqtt://192.168.1.21:1883"
// 	DefaultMosquittoUser     = "homeassistant"
// 	DefaultMosquittoClientId = "GodomoLixee"
// 	DefaultMosquittoTopic    = "zigbee2mqtt/LinkyLixee"
// )

type Config struct {
	Url      url.URL
	User     string
	Pwd      []byte
	Topic    string
	ClientId string
}

func NewConfig(url url.URL, user string, pwd []byte, topic string, clientId string) *Config {
	return &Config{
		Url:      url,
		User:     user,
		Pwd:      pwd,
		Topic:    topic,
		ClientId: clientId,
	}
}
