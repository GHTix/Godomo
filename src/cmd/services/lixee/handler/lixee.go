package handler

import (
	"net/http"

	"github.com/ghtix/gomodo/pkg/common"
	"github.com/gin-gonic/gin"
)

// AvailablePower get the available power
// @Summary get the vailable power
// @Description get the vailable power
// @Tags Lixee
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/availablepower [get]
func (h *Handler) AvailablePower(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("available power", float64(data.AvailablePower), "kVA"))
}

// ApparentPower get the apparent power
// @Summary get the apparent power
// @Description get the apparent power
// @Tags Lixee
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/apparentpower [get]
func (h *Handler) ApparentPower(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("apparent power", float64(data.ApparentPower), "VA"))
}

// CurrentSummDelivered get the apparent power
// @Summary get the apparent power
// @Description get the apparent power
// @Tags Lixee
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Sensor
// @Router /api/v1/currentsummdeliverd [get]
func (h *Handler) CurrentSummDelivered(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, common.NewSensor("current summ deliverd", float64(data.CurrentSummDelivered), "kWh"))
}
