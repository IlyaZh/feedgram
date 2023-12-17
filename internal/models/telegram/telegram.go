package telegram

import (
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	token   string
	limit   int
	timeout int
	offset  int
	api     *tgbotapi.BotAPI
	isDebug bool
}

const (
	DEFAULT_TIMEOUT = 60
	DEFAULT_LIMIT   = 100
)

func NewTelegram(settings entities.ConfigTelegram, isDebug bool) Telegram {
	limit := DEFAULT_LIMIT
	if settings.Limit != nil {
		limit = *settings.Limit
	}

	timeout := DEFAULT_TIMEOUT
	if settings.Timeout != nil {
		timeout = *settings.Timeout
	}

	return Telegram{token: settings.Token, limit: limit, timeout: timeout, offset: 0, isDebug: isDebug}
}

func (t *Telegram) Start() {
	var err error
	t.api, err = tgbotapi.NewBotAPI(t.token)
	if err != nil {
		panic(err.Error())
	}
	t.api.Debug = t.isDebug

	u := tgbotapi.NewUpdate(t.offset)
	u.Timeout = t.timeout

	updates := t.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func (c *Telegram) GetLimit() int {
	return c.limit
}

func (c *Telegram) SetLimit(limit int) {
	if limit < 1 {
		log.Panicf("Trying to set limit value is less than 1: actual value = %d", limit)
	}
	c.limit = limit
}

func (c *Telegram) GetTimeout() int {
	return c.timeout
}

func (c *Telegram) SetTimeout(timeout int) {
	if timeout < 1 {
		log.Panicf("Trying to set timeout value is less than 1: actual value = %d", timeout)
	}
	c.timeout = timeout
}
