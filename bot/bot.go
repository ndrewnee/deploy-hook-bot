package bot

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ndrewnee/deploy-hook-bot/config"
)

type (
	Bot struct {
		config config.Config
		botAPI *tgbotapi.BotAPI
	}

	Options struct {
		Config config.Config
		BotAPI *tgbotapi.BotAPI
	}
)

func New(options ...Options) (*Bot, error) {
	var opts Options

	if len(options) > 0 {
		opts = options[0]
	}

	if opts.Config == (config.Config{}) {
		opts.Config = config.ParseConfig()
	}

	if opts.BotAPI == nil {
		botAPI, err := tgbotapi.NewBotAPI(opts.Config.Token)
		if err != nil {
			return nil, err
		}

		botAPI.Debug = opts.Config.Debug
		opts.BotAPI = botAPI
	}

	log.Printf("Authorized on account %s", opts.BotAPI.Self.UserName)

	return &Bot{
		botAPI: opts.BotAPI,
		config: opts.Config,
	}, nil
}

func (b *Bot) GetUpdatesChan() (tgbotapi.UpdatesChannel, error) {
	if b.config.Webhook {
		webhook := tgbotapi.NewWebhook(b.config.WebhookHost + "/" + b.botAPI.Token)

		if _, err := b.botAPI.SetWebhook(webhook); err != nil {
			return nil, fmt.Errorf("set webhook failed: %s", err)
		}

		info, err := b.botAPI.GetWebhookInfo()
		if err != nil {
			return nil, fmt.Errorf("get webhook info failed: %s", err)
		}

		if info.LastErrorDate != 0 {
			log.Printf("[ERROR] Telegram callback failed: %s", info.LastErrorMessage)
		}

		updates := b.botAPI.ListenForWebhook("/" + b.botAPI.Token)

		go func() {
			if err := http.ListenAndServe(b.config.Address, nil); err != nil {
				log.Printf("[ERROR] Listen and serve failed: %s", err)
			}
		}()

		return updates, nil
	}

	response, err := b.botAPI.RemoveWebhook()
	if err != nil {
		return nil, fmt.Errorf("removed webhook failed: %s", err)
	}

	if !response.Ok {
		return nil, fmt.Errorf("remove webhook response contains error: %s", response.Description)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.botAPI.GetUpdatesChan(u)
	if err != nil {
		return nil, fmt.Errorf("get updates chan failed: %s", err)
	}

	return updates, nil
}
