package main

import (
	"log/slog"
	"os"

	"github.com/ghtix/gomodo/cmd/services/aqara/docs"
	"github.com/ghtix/gomodo/cmd/services/aqara/handler"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config, err := config.New(os.Getenv("GODOMO_AQARA_CONFIG_PATH"))
	if err != nil {
		slog.Error("main", "error loading config", err.Error())
	}

	h := handler.New(*config)

	router := gin.Default()

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

	router.Run(":8080")
}
