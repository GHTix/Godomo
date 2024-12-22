package handler

import (
	"net/http"

	"github.com/ghtix/gomodo/pkg/common"
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

// Temperature get the last temperature mesure
// @Summary get the last temperature mesure
// @Description get the last temperature mesure
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/temperature [get]
func (h *Handler) Temperature(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("temperature", data.Temperature, "°C"))
}

// Pressure get the last pressure mesure
// @Summary get the last pressure mesure
// @Description get the last pressure mesure
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/pressure [get]
func (h *Handler) Pressure(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("pressure", data.Pressure, "hPa"))
}

// Humidity get the last humidity mesure
// @Summary get the last humidity mesure
// @Description get the last humidity mesure
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/humidity [get]
func (h *Handler) Humidity(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("humidity", data.Humidity, "%"))
}

// Battery get the last battery mesure
// @Summary get the last battery mesure
// @Description get the last battery mesure
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/battery [get]
func (h *Handler) Battery(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("battrey", data.Humidity, "%"))
}
