package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/models"
)

const MsgTemplate = `[Build](https://dashboard.heroku.com/apps/%s/activity/builds/%s)
App: %s
Commit: %s
Status: %s
Published at: %s`

type Server struct {
	mux   *http.ServeMux
	tgbot *bot.Bot
}

func NewServer(tgbot *bot.Bot) *Server {
	server := &Server{
		mux:   http.NewServeMux(),
		tgbot: tgbot,
	}

	server.mux.HandleFunc("/hooks", server.HooksHandler)

	return server
}

func (s *Server) Mux() *http.ServeMux {
	return s.mux
}

func (s *Server) HooksHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Read request body failed: %s", err)
		writeResponse(w, http.StatusBadRequest, models.HookResponse{Error: "Read request body failed"})
		return
	}

	var hook models.Hook

	if err := json.Unmarshal(body, &hook); err != nil {
		log.Printf("[ERROR] Unmarshal request body failed: %s. Body: %s\n", err, body)
		writeResponse(w, http.StatusBadRequest, models.HookResponse{Error: "Unmarshal request body failed"})
		return
	}

	if hook.Action != "update" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	text := fmt.Sprintf(
		MsgTemplate,
		hook.Data.App.Name,
		hook.Data.ID,
		hook.Data.App.Name,
		hook.Data.Slug.Commit,
		hook.Data.Status,
		hook.PublishedAt,
	)

	if _, err := s.tgbot.SendMessage(text); err != nil {
		log.Printf("[ERROR] Send message to telegram failed: %s", err)
		writeResponse(w, http.StatusInternalServerError, models.HookResponse{Error: "Send message to telegram failed"})
		return
	}

	writeResponse(w, http.StatusOK, models.HookResponse{Message: text})
}

func writeResponse(w http.ResponseWriter, status int, response models.HookResponse) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := json.Marshal(response)
	_, _ = w.Write(body)
}
