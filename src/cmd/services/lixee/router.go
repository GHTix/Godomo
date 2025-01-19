package main

import (
	lixeeConfig "github.com/ghtix/gomodo/cmd/services/lixee/config"
	"github.com/ghtix/gomodo/cmd/services/lixee/docs"
	"github.com/ghtix/gomodo/cmd/services/lixee/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine, config lixeeConfig.LixeeServiceConfig) error {
	h := handler.New(config)

	var v1 = router.Group("/api/v1")
	v1.Use()
	{
		v1.GET("availablepower", h.AvailablePower)
		v1.GET("apparentpower", h.ApparentPower)
		v1.GET("currentsummdeliverd", h.CurrentSummDelivered)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return nil
}
