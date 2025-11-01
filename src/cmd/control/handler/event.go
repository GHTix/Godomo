package handler

import (
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Event(ctx *gin.Context) {
	slog.Info("Event enter")
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-h.message; ok {
			ctx.SSEvent("message", msg)
			slog.Info("SSEvent sent")
			return true
		}
		slog.Info("No SSEvent")
		return false
	})
	slog.Info("Event leave")
}
