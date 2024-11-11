package main

import (
	"log/slog"

	"github.com/ghtix/gomodo/cmd/services/overkiz/docs"
	"github.com/ghtix/gomodo/cmd/services/overkiz/handler"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/ghtix/gomodo/internal/rest"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.New("../private/config.yml")
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
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8888")

}
