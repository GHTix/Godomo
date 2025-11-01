package handler

import (
	controlConfig "github.com/ghtix/gomodo/cmd/control/config"
)

type Handler struct {
	Config  controlConfig.ControlConfig
	message chan string
}

func New(config controlConfig.ControlConfig) *Handler {
	ch := make(chan string)
	return &Handler{
		Config:  config,
		message: ch,
	}
}
