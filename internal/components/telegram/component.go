package telegram

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// https://github.com/go-telegram-bot-api/telegram-bot-api

type Telegram interface {
	Start(ctx context.Context, output chan<- entities.Message)
	PostMessageHTML(ctx context.Context, message string) error
}

type Component struct {
	token   string
	offset  int
	config  configs.ConfigsCache
	api     TelegramAPI
	updates tgbotapi.UpdatesChannel
}

func NewTelegram(config configs.ConfigsCache, tgApi TelegramAPI) *Component {
	return &Component{
		token:  config.GetValues().Telegram.Token,
		config: config,
		api:    tgApi,
		offset: 0}
}

func (c *Component) Start(ctx context.Context, output chan<- entities.Message) {
	log.Println("Telegram component start")
	settings := c.config.GetValues().Telegram

	u := tgbotapi.NewUpdate(c.offset)
	u.Timeout = *settings.Timeout

	c.updates = c.api.GetUpdatesChan(u)
	go c.handler(ctx, output)
	log.Println("Telegram component has started")
}
