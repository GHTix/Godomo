package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var ColorReset = "\033[0m"
var ColorRed = "\033[31m"
var ColorGreen = "\033[32m"
var ColorYelow = "\033[33m"
var ColorBlue = "\033[34m"

const (
	// ANSI modes
	ansiEsc          = '\u001b'
	ansiReset        = "\u001b[0m"
	ansiFaint        = "\u001b[2m"
	ansiResetFaint   = "\u001b[22m"
	ansiBrightRed    = "\u001b[91m"
	ansiBrightGreen  = "\u001b[92m"
	ansiBrightYellow = "\u001b[93m"
	ansiBrightBlue   = "\u001b[94m"
)

func ColorString(color string, text string) string {
	return string(color) + text + string(ansiReset)
}

func HttpGet(url string) {
	res, err := http.Get(url)
	if err != nil {
		slog.Error(ColorString(ansiBrightRed, "http.Get"), slog.String("url", url), slog.String("error", err.Error()))
		return
	}
	code := res.StatusCode
	status := res.Status
	switch {
	case code >= http.StatusContinue && code < http.StatusMultipleChoices: // 100 to 299
		status = ColorString(ansiBrightGreen, res.Status)
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest: // 300 to 399
		status = ColorString(ansiBrightBlue, res.Status)
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError: // 400 to 499
		status = ColorString(ansiBrightYellow, res.Status)
	case code >= http.StatusInternalServerError: // 500 to
		status = ColorString(ansiBrightRed, res.Status)
	}
	data, err := io.ReadAll(res.Body)
	if err == nil {
		slog.Info(status, slog.String("url", url), slog.Any("body", data))
	} else {
		slog.Info(status, slog.String("url", url))
	}

}

func Loop(url string, sleepTime time.Duration) {
	for {
		go HttpGet(url)
		time.Sleep(sleepTime)
	}
}

func main() {
	notifyInterruptCtx, stopNotifyInterrupt := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopNotifyInterrupt()

	go Loop("http://localhost:8787/metrics", 15*time.Second)
	go Loop("http://localhost:8788/ready", 1*time.Second)

	<-notifyInterruptCtx.Done()
	stopNotifyInterrupt()
}
