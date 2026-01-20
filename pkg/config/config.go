package config

import (
	"log/slog"
	"os"
)

type SiteConfig struct {
	Name           string
	URL            string
	DefaultOGImage string
}

type Config struct {
	Port         string
	Env          string
	Site         SiteConfig
	BrevoAPIKey  string
	ContactEmail string
}

func Load() *Config {
	cfg := &Config{
		Port: getEnvOrDefault("PORT", "3000"),
		Env:  getEnvOrDefault("ENV", "development"),
		Site: SiteConfig{
			Name:           getEnvOrDefault("SITE_NAME", "Cloverbelt Construction"),
			URL:            getEnvOrDefault("SITE_URL", "http://localhost:3000"),
			DefaultOGImage: getEnvOrDefault("DEFAULT_OG_IMAGE", "/images/construction-1.jpg"),
		},
		BrevoAPIKey:  os.Getenv("BREVO_API_KEY"),
		ContactEmail: getEnvOrDefault("CONTACT_EMAIL", "info@cloverbeltconstructionwi.com"),
	}

	if cfg.BrevoAPIKey == "" && cfg.IsProduction() {
		slog.Warn("BREVO_API_KEY not set, contact form emails will not be sent")
	}

	return cfg
}

func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

func (c *Config) IsProduction() bool {
	return c.Env == "production"
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
