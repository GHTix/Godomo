package handler

import (
	"net/http"

	overkiz "github.com/ghtix/gomodo/pkg/overkiz/model"
	"github.com/gin-gonic/gin"
)

// ListHeatingSystems lists all HeatingSystem
// @Summary lists all HeatingSystem
// @Description lists all HeatingSystem
// @Tags HeatingSystem
// @Accept  json
// @Produce  json
// @Success 200 {array} overkiz.Device
// @Router /api/v1/heatingsystems [get]
func (h *Handler) ListHeatingSystems(ctx *gin.Context) {

	devices, err := h.Setup(ctx).DevicesFilterByClass("HeatingSystem")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}

	var sensors []overkiz.Device

	for _, device := range *devices {
		sensor, err := h.Setup(ctx).DevicesFilterByClassAndPlace("TemperatureSensor", device.PlaceOID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		sensors = append(sensors, *sensor...)
	}

	*devices = append(*devices, sensors...)

	ctx.JSON(http.StatusOK, devices)
}
