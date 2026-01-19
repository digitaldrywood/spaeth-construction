package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"spaeth-construction/internal/config"
	"spaeth-construction/internal/email"
	"spaeth-construction/internal/handler"
	"spaeth-construction/internal/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.Load()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	middleware.Setup(e, cfg)

	emailClient := email.NewClient(cfg.BrevoAPIKey, cfg.ContactEmail)
	h := handler.New(emailClient)
	h.RegisterRoutes(e)

	go func() {
		addr := ":" + cfg.Port
		slog.Info("starting server", "url", "http://localhost:"+cfg.Port, "env", cfg.Env)
		if err := e.Start(addr); err != nil {
			slog.Info("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", "error", err)
	}

	slog.Info("server stopped")
}
