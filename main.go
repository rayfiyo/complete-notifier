package main

import (
	"flag"
	"log"
	"time"

	"github.com/rayfiyo/complete-notifier/config"
	"github.com/rayfiyo/complete-notifier/webhook"
)

func main() {
	url := flag.String("url", "", "DISCORD_WEBHOOK_URL(API). Can also be written in .env (see README.md for details).")
	flag.Parse()
	message := flag.Arg(0)

	var cfg *config.Config
	if *url != "" {
		cfg = config.NewWebhookURL(*url)
	} else {
		var err error
		cfg, err = config.Load()
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

	t := time.Now().Format(" 2006-01-02 15:04:05 (Z07:00) ")

	if err := webhook.Send(cfg.WebhookURL, message+t); err != nil {
		log.Fatalf("Failed to send webhook: %v", err)
	}
}
