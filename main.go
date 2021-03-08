package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/config"
	"github.com/ndrewnee/deploy-hook-bot/models"
)

const msgTemplate = `[Build](https://dashboard.heroku.com/apps/%s/activity/builds/%s)
App: %s
Commit: %s
Status: %s
Published at: %s`

func main() {
	rand.Seed(time.Now().UnixNano())
	config := config.ParseConfig()

	tgbot, err := bot.New(bot.Options{Config: config})
	if err != nil {
		log.Fatal("Init telegram bot failed: ", err)
	}

	http.HandleFunc("/hooks", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("[ERROR] Read request body failed: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var hook models.HookAPIBuild

		if err := json.Unmarshal(body, &hook); err != nil {
			log.Printf("[ERROR] Unmarshal request body failed: %s. Body: %s\n", err, body)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		text := fmt.Sprintf(
			msgTemplate,
			hook.Data.App.Name,
			hook.Data.ID,
			hook.Data.App.Name,
			hook.Data.Slug.Commit,
			hook.Data.Status,
			hook.PublishedAt,
		)

		if _, err := tgbot.SendMessage(text); err != nil {
			log.Printf("[ERROR] Send message to telegram failed: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	if err := http.ListenAndServe(config.Address, nil); err != nil {
		log.Printf("[ERROR] Listen and serve failed: %s", err)
	}
}
