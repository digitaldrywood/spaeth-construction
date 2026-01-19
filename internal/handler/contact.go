package handler

import (
	"log/slog"
	"net/http"

	"spaeth-construction/templates/pages"

	"github.com/labstack/echo/v4"
)

type ContactForm struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Phone   string `form:"phone"`
	Service string `form:"service"`
	Message string `form:"message"`
}

func (h *Handler) ContactSubmit(c echo.Context) error {
	var form ContactForm
	if err := c.Bind(&form); err != nil {
		slog.Error("failed to bind contact form", "error", err)
		return pages.Contact(false, "Invalid form submission. Please try again.").Render(c.Request().Context(), c.Response().Writer)
	}

	// Validate required fields
	if form.Name == "" || form.Email == "" || form.Message == "" {
		return pages.Contact(false, "Please fill in all required fields.").Render(c.Request().Context(), c.Response().Writer)
	}

	// Send email via Brevo if configured
	if h.email != nil && h.email.APIKey() != "" {
		err := h.email.SendContactNotification(
			form.Name,
			form.Email,
			form.Phone,
			form.Service,
			form.Message,
		)
		if err != nil {
			slog.Error("failed to send contact email", "error", err)
			// Still show success to user, but log the error
		}
	} else {
		slog.Info("contact form submitted (email not configured)",
			"name", form.Name,
			"email", form.Email,
			"phone", form.Phone,
			"service", form.Service,
		)
	}

	// Return success page
	c.Response().Header().Set("HX-Push-Url", "/contact?success=true")
	return pages.Contact(true, "").Render(c.Request().Context(), c.Response().Writer)
}

// ContactFormPartial returns just the form for HTMX requests
func (h *Handler) ContactFormPartial(c echo.Context) error {
	// Check if this is an HTMX request
	if c.Request().Header.Get("HX-Request") == "true" {
		return c.HTML(http.StatusOK, `<div class="text-green-400 text-center p-8">Thank you for your message! We'll be in touch soon.</div>`)
	}
	return h.Contact(c)
}
