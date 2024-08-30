package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebhookURL string
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

	return &Config{
		WebhookURL: webhookURL,
	}, nil
}
