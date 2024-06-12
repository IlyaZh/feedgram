package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// https://github.com/go-telegram-bot-api/telegram-bot-api

type Telegram interface {
	Start(ctx context.Context, output chan<- entities.Message)
	PostMessageHTML(ctx context.Context, message entities.TelegramPost) error
}

var name string = "Telegram"

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
	log := logger.GetLoggerComponent(ctx, name)
	log.Info("Component start")
	settings := c.config.GetValues().Telegram

	u := tgbotapi.NewUpdate(c.offset)
	u.Timeout = *settings.Timeout

	c.updates = c.api.GetUpdatesChan(u)
	go c.handler(ctx, output)
	log.Info("Component has started")
}
