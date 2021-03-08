package config

import (
	"os"
	"strconv"
)

type Config struct {
	Address     string
	Token       string
	WebhookHost string
	Webhook     bool
	Debug       bool
}

func ParseConfig() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 9999
	}

	webhookHost := os.Getenv("WEBHOOK_HOST")
	if webhookHost == "" {
		webhookHost = "https://deploy-hook-bot.herokuapp.com"
	}

	return Config{
		Address:     ":" + strconv.Itoa(port),
		WebhookHost: webhookHost,
		Token:       os.Getenv("TOKEN"),
		Webhook:     os.Getenv("WEBHOOK") == "true",
		Debug:       os.Getenv("DEBUG") == "true",
	}
}
