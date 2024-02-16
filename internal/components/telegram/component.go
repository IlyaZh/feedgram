package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/rss_reader"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// https://github.com/go-telegram-bot-api/telegram-bot-api

type Component struct {
	token     string
	offset    int
	config    *configs.Cache
	api       *tgbotapi.BotAPI
	isDebug   bool
	updates   tgbotapi.UpdatesChannel
	storage   storage.Storage
	rssReader rss_reader.RssReader
}

func NewTelegram(config *configs.Cache, storage *storage.Component, rssReader *rss_reader.Component, isDebug bool) Component {
	return Component{
		token:   config.GetValues().Telegram.Token,
		config:  config,
		offset:  0,
		isDebug: isDebug,
		storage: storage}
}

func (c *Component) Start() {
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
	go c.handler(context.TODO())
}
