package mqttClient

import (
	"fmt"
	"log/slog"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe(mqttOptions *mqtt.ClientOptions, topic string, callback mqtt.MessageHandler) error {

	mqttOptions.SetClientID(fmt.Sprintf("godomo_%s", topic))
	mqttClient := mqtt.NewClient(mqttOptions)
	tokenConnect := mqttClient.Connect()
	for !tokenConnect.WaitTimeout(3 * time.Second) {
	}
	if err := tokenConnect.Error(); err != nil {
		slog.Error("Subscribe", "Connect", err.Error(), "topic", topic)
		return fmt.Errorf("Subscribe Connect: %w", err)
	}

	tokenSubscribe := mqttClient.Subscribe(topic, 0, callback)
	for !tokenSubscribe.WaitTimeout(3 * time.Second) {
	}
	if err := tokenSubscribe.Error(); err != nil {
		slog.Error("Subscribe", "Connect", err.Error(), "topic", topic)
		return fmt.Errorf("Subscribe Subscribe: %w", err)
	}

	slog.Info("Subscribe", "topic", topic)
	return nil
}
