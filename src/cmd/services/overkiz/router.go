package main

import (
	"github.com/ghtix/gomodo/cmd/services/overkiz/docs"
	"github.com/ghtix/gomodo/cmd/services/overkiz/handler"
	"github.com/ghtix/gomodo/internal/rest"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine, c *rest.RestClient) error {

	h := handler.New(*c)
	var v1 = router.Group("/api/v1")
	v1.Use()
	{
		v1.GET("rollers", h.ListRollers)
		v1.GET("rollers/:oid", h.GetRollerByOID)
		v1.POST("rollers/:oid/open", h.OpenRollerByOID)
		v1.POST("rollers/:oid/close", h.CloseRollerByOID)

		v1.GET("heatingsystems", h.ListHeatingSystems)
	}

	router.GET("test", h.Test)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return nil
}
