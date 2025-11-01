package main

import (
	controlConfig "github.com/ghtix/gomodo/cmd/control/config"
	"github.com/ghtix/gomodo/cmd/control/docs"
	"github.com/ghtix/gomodo/cmd/control/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine, config controlConfig.ControlConfig) error {

	h := handler.New(config)

	var apiV1 = router.Group("/api/v1")
	var sseV1 = router.Group("/sse/v1")
	var renderV1 = router.Group("/render/v1")

	apiV1.Use()
	{
		apiV1.POST("notify", h.Notify)
	}
	sseV1.Use()
	{
		sseV1.GET("events", h.Event)
	}
	renderV1.Use()
	{
		renderV1.GET("dashboard", h.Dashboard)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return nil
}
