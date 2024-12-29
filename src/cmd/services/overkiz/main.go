package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	overkizConfig "github.com/ghtix/gomodo/cmd/services/overkiz/config"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/ghtix/gomodo/internal/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("OVERKIZ")

	configFile := os.Getenv("GODOMO_OVERKIZ_CONFIG_PATH")
	config, err := config.New[overkizConfig.OverkizServiceConfig](configFile)
	if err != nil {
		slog.Error("main", "error loading config", err.Error(), "file", configFile)
		return
	} else {
		slog.Info("main", "Config OK from", configFile)
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

	SetupRouter(router, c)

	server := &http.Server{
		Addr:    config.Service.ListeningAddress,
		Handler: router.Handler(),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Listen")
		}
	}()

	<-ctx.Done()
	stop()

	// Cleanup
	slog.Info("Clean up..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error while shutting down Server. Initiating force shutdown...", "err", err)
	} else {
		slog.Info("Server exiting")
	}

}
