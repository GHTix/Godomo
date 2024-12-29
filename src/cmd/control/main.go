package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	controlConfig "github.com/ghtix/gomodo/cmd/control/config"
	"github.com/ghtix/gomodo/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("CONTROL")

	configFile := os.Getenv("GODOMO_CONTROL_CONFIG_PATH")
	config, err := config.New[controlConfig.ControlConfig](configFile)
	if err != nil {
		slog.Error("main", "error loading config", err.Error(), "file", configFile)
		return
	} else {
		slog.Info("main", "Config OK from", configFile)
	}

	for _, svc := range config.ControlledServices {
		slog.Info("Service", "Name", svc.Name, "Port", svc.Port)
	}

	router := gin.Default()

	SetupRouter(router)

	slog.Info(config.Service.ListeningAddress)

	server := &http.Server{
		Addr:    config.Service.ListeningAddress,
		Handler: router.Handler(),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Listen", "Err", err.Error())
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
