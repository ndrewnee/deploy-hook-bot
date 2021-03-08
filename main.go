package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/config"
	"github.com/ndrewnee/deploy-hook-bot/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	config := config.ParseConfig()

	tgbot, err := bot.New(bot.Options{Config: config})
	if err != nil {
		log.Fatal("Init telegram bot failed: ", err)
	}

	http.HandleFunc("/hooks", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)

		log.Printf("Request body:\n%s", bytes)

		var hook models.Hook
		json.Unmarshal(bytes, &hook)

		tgbot.SendMessage(hook.Data.Description)
	})

	if err := http.ListenAndServe(config.Address, nil); err != nil {
		log.Printf("[ERROR] Listen and serve failed: %s", err)
	}
}
