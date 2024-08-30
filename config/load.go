package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func NewWebhookURL(w string) *Config {
	return &Config{
		WebhookURL: w,
	}
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		return nil, fmt.Errorf("DISCORD_WEBHOOK_URL is not set in .env file")
	}

	return NewWebhookURL(webhookURL), nil
}
