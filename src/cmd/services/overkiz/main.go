package main

import (
	"log/slog"
	"os"

	"github.com/ghtix/gomodo/cmd/services/overkiz/docs"
	"github.com/ghtix/gomodo/cmd/services/overkiz/handler"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/ghtix/gomodo/internal/rest"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	config, err := config.New(os.Getenv("GODOMO_OVERKIZ_CONFIG_PATH"))
	if err != nil {
		slog.Error("main", "error loading config", err.Error())
	}

	c := rest.New(
		config.Overkiz.BaseUrl,
		rest.WithOAuth(rest.NewOAuthData(
			rest.OAuthConfig{
				LoginEndPoint: config.Overkiz.OAuthLoginEndpoint,
				ClientId:      config.Overkiz.OAuthClientId,
				ClientSecret:  config.Overkiz.OAuthClientSecret,
				UserName:      config.Overkiz.UserName,
				Password:      config.Overkiz.Password,
			}),
		),
	)

	router := gin.Default()

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

	router.Run(":8080")

}
