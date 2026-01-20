package handler

import (
	"spaeth-construction/pkg/email"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	email *email.Client
}

func New(emailClient *email.Client) *Handler {
	return &Handler{
		email: emailClient,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	// Static files - serve from public directory
	e.Static("/images", "public/images")
	e.Static("/static", "public/static")
	e.File("/favicon.ico", "public/favicon.ico")

	// Health check
	e.GET("/health", h.Health)

	// Public pages
	e.GET("/", h.Home)
	e.GET("/about", h.About)
	e.GET("/services", h.Services)
	e.GET("/projects", h.Projects)
	e.GET("/contact", h.Contact)

	// Contact form submission
	e.POST("/contact", h.ContactSubmit)
}
