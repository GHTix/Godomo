package handler

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/ghtix/gomodo/internal/config"
	mqttClient "github.com/ghtix/gomodo/internal/mqtt"
	"github.com/ghtix/gomodo/pkg/aqara"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Handler struct {
	Config config.Config
	data   *aqara.Aqara
	mutex  sync.RWMutex
}

func (h *Handler) SetData(data []byte) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	var err error
	h.data, err = aqara.New(data)
	if err != nil {
		return fmt.Errorf("handler SetData: %w", err)
	}
	return nil

}

func (h *Handler) GetData() (*aqara.Aqara, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.data, nil

}

func (h *Handler) MessageCallback(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	switch topic {
	case "zigbee2mqtt/Aqara2":

		err := h.SetData(msg.Payload())
		if err != nil {
			slog.Warn(topic, "Set Handler data", err.Error())
		}

	default:
		slog.Warn("unknown topic", "payload", string(msg.Payload()))
	}
}

func New(config config.Config) *Handler {
	handler := Handler{
		Config: config,
	}

	mqttOptions := mqtt.NewClientOptions()
	mqttOptions.AddBroker(config.Mqtt.BrokerUrl)
	mqttOptions.SetUsername(config.Mqtt.UserName)
	mqttOptions.SetPassword(config.Mqtt.Password)

	go mqttClient.Subscribe(mqttOptions, "zigbee2mqtt/Aqara2", handler.MessageCallback)

	return &handler
}
