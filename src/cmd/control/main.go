package main

import (
	"github.com/ghtix/gomodo/cmd/services/aqara/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// config, err := config.New("../private/config.yml")
	// if err != nil {
	// 	slog.Error("main", "error loading config", err.Error())
	// }

	router := gin.Default()

	// h := handler.New()

	var v1 = router.Group("/api/v1")
	v1.Use()
	{
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8880")

}
