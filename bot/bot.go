package bot

import (
	"fmt"
	"log"

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

func (b *Bot) SendMessage(text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(b.config.TelegramChatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	sent, err := b.botAPI.Send(msg)
	if err != nil {
		errMsg := msg
		errMsg.Text = "Oops, something went wrong!"
		_, _ = b.botAPI.Send(errMsg)

		return tgbotapi.Message{}, fmt.Errorf("send message failed: %s. Text: \n%s", err, msg.Text)
	}

	return sent, nil
}
