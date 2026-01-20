package handler

import (
	"net/http"

	"spaeth-construction/templates/pages"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) Home(c echo.Context) error {
	return pages.Home().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) About(c echo.Context) error {
	return pages.About().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Services(c echo.Context) error {
	return pages.Services().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Projects(c echo.Context) error {
	return pages.Projects().Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) Contact(c echo.Context) error {
	return pages.Contact(false, "").Render(c.Request().Context(), c.Response().Writer)
}
