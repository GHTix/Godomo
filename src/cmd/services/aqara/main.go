package main

import (
	"log/slog"

	"github.com/ghtix/gomodo/cmd/services/aqara/docs"
	"github.com/ghtix/gomodo/cmd/services/aqara/handler"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config, err := config.New("../private/config.yml")
	if err != nil {
		slog.Error("main", "error loading config", err.Error())
	}

	h := handler.New(*config)

	router := gin.Default()

	var v1 = router.Group("/api/v1")
	v1.Use()
	{
		v1.GET("sensor", h.Sensor)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8889")
}
