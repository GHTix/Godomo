package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/eclipse/paho.golang/paho"
	probeService "github.com/ghtix/gomodo/internal/probe"
	"github.com/ghtix/gomodo/pkg/linkee"
	"github.com/ghtix/gomodo/pkg/mosquitto"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var (
	tracer      = otel.Tracer("info-service")
	meter       = otel.Meter("info-service")
	viewCounter metric.Int64Counter
)

func init() {
	var err error
	viewCounter, err = meter.Int64Counter("user.views",
		metric.WithDescription("The number of views"),
		metric.WithUnit("{views}"))
	if err != nil {
		panic(err)
	}
}

type LixeeService struct {
	Count       int
	CurrentData *linkee.LinkyLixee
}

func NewLixeeService() *LixeeService {
	return &LixeeService{
		Count: 0,
	}
}

func (service *LixeeService) HelloHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	s := "Hello, World! #" + strconv.Itoa(service.Count)
	slog.Info(s)
	service.Count = service.Count + 1
	w.Write([]byte(s))
}

func (service *LixeeService) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "info")
	defer span.End()

	viewCounter.Add(ctx, 1)

	s := ""
	if service.CurrentData != nil {
		s = service.CurrentData.PromMetrics()
	}
	//slog.Info(s)
	w.Write([]byte(s))
}

func (service *LixeeService) MqttLixeeMessageHandler(publish *paho.Publish) {
	data, err := linkee.New(publish.Payload)
	if err != nil {
		slog.Error("MqttLixeeMessageHandler", slog.String("err", err.Error()))
		return
	}
	slog.Info("MqttLixeeMessageHandler" /*, slog.Any("Payload", data)*/)
	service.CurrentData = data
}

func StartService(mosquittoConfig *mosquitto.Config) {
	notifyInterruptCtx, stopNotifyInterrupt := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopNotifyInterrupt()

	muxService := http.NewServeMux()
	muxProbe := http.NewServeMux()

	probe := probeService.New()
	service := NewLixeeService()

	instrumetedHandleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		muxService.Handle(pattern, handler)
	}

	httpOtelHandler := otelhttp.NewHandler(muxService, "/")

	//muxService.HandleFunc("/hello", service.HelloHandler)
	//muxService.HandleFunc("/metrics", service.MetricsHandler)
	instrumetedHandleFunc("/hello", service.HelloHandler)
	instrumetedHandleFunc("/metrics", service.MetricsHandler)

	muxProbe.HandleFunc("/ready", probe.ProbeReadyHandler)
	muxProbe.HandleFunc("/alive", probe.ProbeAliveHandler)

	server := &http.Server{
		Addr:         ":8787",
		Handler:      httpOtelHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	serverProbe := &http.Server{
		Addr:         ":8788",
		Handler:      muxProbe,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  2 * time.Second,
	}

	// Open Telemetry Trace
	consoleTraceExporter, err := newTraceExporter()
	if err != nil {
		slog.Error("consoleTraceExporter", slog.String("err", err.Error()))
	}

	tracerProvider := newTraceProvider(consoleTraceExporter)
	defer tracerProvider.Shutdown(notifyInterruptCtx)
	otel.SetTracerProvider(tracerProvider)

	// Open Telemetry Metrics
	consoleMetricExporter, err := newMetricExporter()
	if err != nil {
		slog.Error("consoleMetricExporter", slog.String("err", err.Error()))
	}
	meterProvider := newMeterProvider(consoleMetricExporter)

	defer meterProvider.Shutdown(notifyInterruptCtx)
	otel.SetMeterProvider(meterProvider)

	// Open Telemetry Propagator
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// MQTT
	mosquittoClient := mosquitto.New(mosquittoConfig)

	shutdownChan := make(chan bool, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("server.ListenAndServe", slog.String("err", err.Error()))
		}
		time.Sleep(5 * time.Second)
		shutdownChan <- true
	}()

	go func() {
		err := serverProbe.ListenAndServe()
		if err != nil {
			slog.Error("serverProbe.ListenAndServe", slog.String("err", err.Error()))
		}
	}()

	go func() {
		err := mosquittoClient.Start(notifyInterruptCtx, service.MqttLixeeMessageHandler)
		if err != nil {
			slog.Error("mosquittoClient.Start", slog.String("err", err.Error()))
		}
	}()

	<-notifyInterruptCtx.Done()
	stopNotifyInterrupt()

	probe.Ready = false

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 2*time.Second)
	defer shutdownRelease()

	err = mosquittoClient.Stop(shutdownCtx)
	if err != nil {
		slog.Error("mosquittoClient.Stop", slog.String("err", err.Error()))
	}

	slog.Info("server.Shutdown")
	err = server.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("server.ListenAndServe", slog.String("err", err.Error()))
	}

	<-shutdownChan
	slog.Info("ApiBase end")
}
