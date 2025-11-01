package handler

import (
	"html/template"
	"path"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Dashboard(ctx *gin.Context) {

	templateFile := path.Join(h.Config.Template.TemplatePath, "dashboard.html")

	tmpl := template.Must(template.ParseFiles(templateFile))

	err := tmpl.Execute(ctx.Writer, gin.H{})
	if err != nil {
		panic(err)
	}
}
