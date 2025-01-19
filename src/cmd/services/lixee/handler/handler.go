package handler

import (
	"fmt"
	"log/slog"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	lixeeConfig "github.com/ghtix/gomodo/cmd/services/lixee/config"
	mqttClient "github.com/ghtix/gomodo/internal/mqtt"
	"github.com/ghtix/gomodo/pkg/linkee"
)

type Handler struct {
	Config lixeeConfig.LixeeServiceConfig
	data   *linkee.LinkyLixee
	mutex  sync.RWMutex
}

func New(config lixeeConfig.LixeeServiceConfig) *Handler {

	slog.Info("Lixee", "Mqtt", config.Mqtt.BrokerUrl)

	if len(config.Mqtt.Topics) != 1 {
		slog.Warn("No topics defined")
		return nil
	}
	slog.Info("Lixee", "Topic", config.Mqtt.Topics[0].Name)

	handler := Handler{
		Config: config,
	}

	mqttOptions := mqtt.NewClientOptions()
	mqttOptions.AddBroker(config.Mqtt.BrokerUrl)
	mqttOptions.SetUsername(config.Mqtt.UserName)
	mqttOptions.SetPassword(config.Mqtt.Password)

	go mqttClient.Subscribe(mqttOptions, config.Mqtt.Topics[0].Name, handler.MessageCallback)

	return &handler
}

func (h *Handler) MessageCallback(_ mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	switch topic {
	case h.Config.Mqtt.Topics[0].Name:

		err := h.SetData(msg.Payload())
		if err != nil {
			slog.Warn(topic, "Set Handler data", err.Error())
		}

	default:
		slog.Warn("unknown topic", "payload", string(msg.Payload()))
	}
}

func (h *Handler) SetData(data []byte) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	var err error
	h.data, err = linkee.New(data)
	slog.Debug("Lixee", "GetData", h.data)
	if err != nil {
		return fmt.Errorf("handler SetData: %w", err)
	}
	return nil
}

func (h *Handler) GetData() (*linkee.LinkyLixee, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	slog.Debug("Lixee", "GetData", h.data)
	return h.data, nil

}
