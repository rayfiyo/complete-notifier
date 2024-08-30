package main

import (
	"flag"
	"log"

	"github.com/rayfiyo/complete-notifier/config"
	"github.com/rayfiyo/complete-notifier/webhook"
)

func main() {
	flag.Parse()
	message := flag.Arg(0)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	err = webhook.Send(cfg.WebhookURL, message)
	if err != nil {
		log.Fatalf("Failed to send webhook: %v", err)
	}

	log.Println("Webhook sent successfully")
}
