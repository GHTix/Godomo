package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyMessage struct {
	Message string
}

// @Description Notify
// @Tag Notifier
// @Accept  json
// @Param data body handler.NotifyMessage true "The message"
// @Success 202
// @Failure 400
// @Router /api/v1/notify [post]
func (h *Handler) Notify(ctx *gin.Context) {

	var msg NotifyMessage
	if err := ctx.ShouldBind(&msg); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}
	slog.Info(msg.Message)
	h.message <- msg.Message
	ctx.JSON(http.StatusAccepted, nil)
}
