package handler

import (
	"context"
	"log/slog"

	"github.com/ghtix/gomodo/internal/rest"
	overkiz "github.com/ghtix/gomodo/pkg/overkiz/model"
)

type Handler struct {
	Client rest.RestClient
	setup  *overkiz.Setup
}

func New(client rest.RestClient) *Handler {
	return &Handler{
		Client: client,
		setup:  nil,
	}
}

func (h *Handler) Setup(ctx context.Context) *overkiz.Setup {
	if h.setup != nil {
		return h.setup
	}
	resp, err := h.Client.Get(ctx, "setup")
	if err != nil {
		slog.Error("goverkiz-cli", "error getting setup", err.Error())
	}

	setup, err := rest.JsonResponse[overkiz.Setup](resp)
	if err != nil {
		slog.Error("goverkiz-cli", "error json unmarshaling setup", err.Error())
	}
	return setup
}
