package main

import (
	aqaraConfig "github.com/ghtix/gomodo/cmd/services/aqara/config"
	"github.com/ghtix/gomodo/cmd/services/aqara/docs"
	"github.com/ghtix/gomodo/cmd/services/aqara/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine, config aqaraConfig.AqaraServiceConfig) error {
	h := handler.New(config)

	var v1 = router.Group("/api/v1")
	v1.Use()
	{
		v1.GET("sensor", h.Sensor)
		v1.GET("temperature", h.Temperature)
		v1.GET("pressure", h.Pressure)
		v1.GET("humidity", h.Humidity)
		v1.GET("battery", h.Battery)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return nil
}
