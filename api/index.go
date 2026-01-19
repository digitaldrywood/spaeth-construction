package api

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"spaeth-construction/internal/config"
	"spaeth-construction/internal/ctxkeys"
	"spaeth-construction/internal/email"
	"spaeth-construction/internal/handler"
)

var (
	e    *echo.Echo
	once sync.Once
)

func initEcho() {
	cfg := config.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	emailClient := email.NewClient(cfg.BrevoAPIKey, cfg.ContactEmail)

	e = echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// Inject site config into context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), ctxkeys.SiteConfig, cfg.Site)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	h := handler.New(emailClient)
	h.RegisterRoutes(e)
}

// Handler is the Vercel serverless function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initEcho)
	e.ServeHTTP(w, r)
}
