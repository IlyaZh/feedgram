package telegram

import (
	"log"

	"github.com/IlyaZh/feedsgram/internal/models/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	token string
	// limit   int
	// timeout int
	offset  int
	configs *config.Cache
	api     *tgbotapi.BotAPI
	isDebug bool
}

func NewTelegram(configs *config.Cache, isDebug bool) Telegram {
	token := configs.GetValues().Components.Telegram.Token
	return Telegram{token: token, configs: configs, offset: 0, isDebug: isDebug}
}

func (t *Telegram) start() {
	var err error
	settings := t.configs.GetValues().Components.Telegram
	t.api, err = tgbotapi.NewBotAPI(t.token)
	if err != nil {
		panic(err.Error())
	}
	t.api.Debug = t.isDebug

	u := tgbotapi.NewUpdate(t.offset)
	u.Timeout = *settings.Timeout

	updates := t.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func (t *Telegram) Start() {
	go t.start()
}
