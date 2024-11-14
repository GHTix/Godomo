package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sensor get the last data of the sensor
// @Summary get the last data of the sensor
// @Description get the last data of the sensor
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Success 200 {object} aqara.Aqara
// @Router /api/v1/sensor [get]
func (h *Handler) Sensor(ctx *gin.Context) {

	data, err := h.GetData()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}

	if data == nil {
		ctx.JSON(http.StatusNoContent, nil)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, data)
}
