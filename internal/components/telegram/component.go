package telegram

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// https://github.com/go-telegram-bot-api/telegram-bot-api

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type Telegram interface {
	Start(ctx context.Context, output chan<- entities.Message)
	PostMessageHTML(ctx context.Context, message string) error
}

type Component struct {
	token   string
	offset  int
	config  *configs.Cache
	api     *tgbotapi.BotAPI
	isDebug bool
	updates tgbotapi.UpdatesChannel
}

func NewTelegram(config *configs.Cache, isDebug bool) Telegram {
	return &Component{
		token:   config.GetValues().Telegram.Token,
		config:  config,
		offset:  0,
		isDebug: isDebug}
}

func (c *Component) Start(ctx context.Context, output chan<- entities.Message) {
	log.Println("Telegram component start")
	var err error
	settings := c.config.GetValues().Telegram
	c.api, err = tgbotapi.NewBotAPI(c.token)
	if err != nil {
		panic(err.Error())
	}
	c.api.Debug = c.isDebug

	u := tgbotapi.NewUpdate(c.offset)
	u.Timeout = *settings.Timeout

	c.updates = c.api.GetUpdatesChan(u)
	go c.handler(ctx, output)
	log.Println("Telegram component has started")
}
