package main

import (
	"fmt"
	"gomodo/internal/config"
	"gomodo/pkg/aqara"
	"gomodo/pkg/bridge"
	"gomodo/pkg/linkee"
	"log/slog"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(mqttOptions *mqtt.ClientOptions) (mqtt.Client, error) {
	slog.Info("connect")
	mqttClient := mqtt.NewClient(mqttOptions)
	token := mqttClient.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {

		slog.Error(err.Error())
		return nil, err
	}
	return mqttClient, nil
}

var messageSubHandler mqtt.MessageHandler = func(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	switch topic {
	case "zigbee2mqtt/LinkyLixee":

		linkyLixee, err := linkee.New(msg.Payload())
		if err == nil {
			slog.Info(topic, "ApparentPower", linkyLixee.ApparentPower, "CurrentSummDelivered", linkyLixee.CurrentSummDelivered)
		} else {
			slog.Warn(topic, "payload", string(msg.Payload()))
		}
	case "zigbee2mqtt/Aqara2":

		aqara, err := aqara.New(msg.Payload())
		if err == nil {
			slog.Debug(topic, "T", aqara.Temperature, "P", aqara.Pressure, "H", aqara.Humidity, "B", aqara.Battery)
		} else {
			slog.Warn(topic, "payload", string(msg.Payload()))
		}
	case "zigbee2mqtt/bridge/devices":
		devices, err := bridge.NewDevices(msg.Payload())
		if err == nil {
			for i, device := range devices {
				slog.Info(topic, "Item", string(i), "FriendlyName", device.FriendlyName, "Enabled", !device.Disabled)
			}
		} else {
			slog.Warn(topic, "payload", string(msg.Payload()))
		}
	case "zigbee2mqtt/bridge/logging":

		log, err := bridge.NewLogging(msg.Payload())
		if err == nil {
			slog.Debug(topic, "Message", log.Message)
		} else {
			slog.Warn(topic, "payload", string(msg.Payload()))
		}
	default:
		slog.Warn("unknown topic", "payload", string(msg.Payload()))
	}
}

func listen(mqttOptions *mqtt.ClientOptions, topic string) {
	slog.Info("listen", "topic", topic)
	mqttOptions.SetClientID(fmt.Sprintf("godomo_%s", topic))
	mqttClient, err := connect(mqttOptions)
	if err != nil {
		slog.Error("listen", "err", err.Error())
	}
	mqttClient.Subscribe(topic, 0, messageSubHandler)
	slog.Info("listen", "end", topic)
}

// func timer() {
// 	for {
// 		slog.Info("HeartBeat")
// 		time.Sleep(10 * time.Second)
// 	}
// }

func main() {
	fmt.Println("Go go godomo")

	config, err := config.New("../private/config.yml")
	if err != nil {
		slog.Error("main", "error loading config", err.Error())
	}

	//https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
	mqttOptions := mqtt.NewClientOptions()
	mqttOptions.AddBroker(config.Mqtt.BrokerUrl)
	mqttOptions.SetUsername(config.Mqtt.UserName)
	mqttOptions.SetPassword(config.Mqtt.Password)

	go listen(mqttOptions, "zigbee2mqtt/bridge/devices")
	go listen(mqttOptions, "zigbee2mqtt/bridge/logging")
	go listen(mqttOptions, "zigbee2mqtt/LinkyLixee")
	go listen(mqttOptions, "zigbee2mqtt/Aqara2")

	// go timer()

	time.Sleep(6000 * time.Second)

}
