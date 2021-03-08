package main

import (
	"log"
	"net/http"

	"github.com/ndrewnee/deploy-hook-bot/api"
	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/config"
)

func main() {
	config := config.ParseConfig()

	tgbot, err := bot.New(bot.Options{Config: config})
	if err != nil {
		log.Fatal("Init telegram bot failed: ", err)
	}

	server := api.NewServer(tgbot)

	if err := http.ListenAndServe(config.Address, server.Mux()); err != nil {
		log.Printf("[ERROR] Listen and serve failed: %s", err)
	}
}
