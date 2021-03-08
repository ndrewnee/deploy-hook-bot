package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Address        string
	Token          string
	AuthToken      string
	TelegramChatID int64
	Debug          bool
}

func Parse() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 9998
	}

	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		log.Printf("[WARN] Env variable TELEGRAM_CHAT_ID is not set: %s", err)
	}

	return Config{
		Address:        ":" + strconv.Itoa(port),
		Token:          os.Getenv("TOKEN"),
		AuthToken:      os.Getenv("AUTH_TOKEN"),
		TelegramChatID: chatID,
		Debug:          os.Getenv("DEBUG") == "true",
	}
}
