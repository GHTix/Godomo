package handler

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Test(ctx *gin.Context) {

	rollers, _ := h.Setup(ctx).DevicesFilterByClass("RollerShutter")
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	ctx.Abort()
	// 	return
	// }
	// ctx.JSON(http.StatusOK, rollers)

	tmpl := template.Must(template.ParseFiles("./template/heating_system.html"))
	err := tmpl.Execute(ctx.Writer, gin.H{
		"rollers": rollers,
	})
	if err != nil {
		panic(err)
	}
}
