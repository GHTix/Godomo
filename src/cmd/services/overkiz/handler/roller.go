package handler

import (
	"io"
	"log/slog"
	"net/http"

	overkiz "github.com/ghtix/gomodo/pkg/overkiz/model"
	"github.com/gin-gonic/gin"
)

// ListRollers lists all RollerShutter
// @Summary lists all RollerShutter
// @Description lists all RollerShutter
// @Tags Roller
// @Accept  json
// @Produce  json
// @Success 200 {array} overkiz.Device
// @Router /api/v1/rollers [get]
func (h *Handler) ListRollers(ctx *gin.Context) {

	rollers, err := h.Setup(ctx).DevicesFilterByClass("RollerShutter")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, rollers)
}

// GetRollerByOID returns device filterd by oid
// @Summary returns device filterd by oid
// @Description returns device filterd by oid
// @Tags Roller
// @Accept  json
// @Produce  json
// @Param oid path string true "oid of the device"
// @Success 200 {object} overkiz.Device
// @Router /api/v1/rollers/{oid} [get]
func (h *Handler) GetRollerByOID(ctx *gin.Context) {
	oid := ctx.Param("oid")

	device, err := h.Setup(ctx).DeviceFilterByOid(oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	if device == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	}
	ctx.JSON(http.StatusOK, device)

}

// OpenRollerByOID opens the roller
// @Summary opens the roller
// @Description opens the roller
// @Tags Roller
// @Accept  json
// @Produce  json
// @Param oid path string true "oid of the device"
// @Success 200 {object} overkiz.Device
// @Router /api/v1/rollers/{oid}/open [post]
func (h *Handler) OpenRollerByOID(ctx *gin.Context) {
	oid := ctx.Param("oid")

	device, err := h.Setup(ctx).DeviceFilterByOid(oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	if device == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		ctx.Abort()
		return
	}

	command := overkiz.NewSingleCommand(device.Label, device.DeviceURL, "open", nil)

	resp, err := h.Client.PostJson(ctx, "exec/apply", command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("goverkiz-cli", "error reading response exec current", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, data)
}

// CloseRollerByOID closes the roller
// @Summary closes the roller
// @Description closes the roller
// @Tags Roller
// @Accept  json
// @Produce  json
// @Param oid path string true "oid of the device"
// @Success 200 {object} overkiz.Device
// @Router /api/v1/rollers/{oid}/close [post]
func (h *Handler) CloseRollerByOID(ctx *gin.Context) {
	oid := ctx.Param("oid")

	device, err := h.Setup(ctx).DeviceFilterByOid(oid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	if device == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		ctx.Abort()
		return
	}

	command := overkiz.NewSingleCommand(device.Label, device.DeviceURL, "close", nil)

	resp, err := h.Client.PostJson(ctx, "exec/apply", command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("goverkiz-cli", "error reading response exec current", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, data)
}
