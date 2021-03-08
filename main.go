package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/config"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config := config.ParseConfig()

	tgbot, err := bot.New(bot.Options{Config: config})
	if err != nil {
		log.Fatal("Init telegram bot failed: ", err)
	}

	updates, err := tgbot.GetUpdatesChan()
	if err != nil {
		log.Fatal("Get updates chan failed: ", err)
	}

	for range updates {
	}
}
