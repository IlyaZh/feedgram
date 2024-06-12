package message_sender

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/telegram"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"go.uber.org/zap"
)

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type MeesageSender interface {
	Start(ctx context.Context)
}

const name string = "MessageSender"

type Component struct {
	config    configs.ConfigsCache
	telegram  telegram.Telegram
	feedsChan <-chan []entities.FeedItem
	postsChan <-chan entities.TelegramPost
}

func NewMeesageSender(config configs.ConfigsCache, telegram telegram.Telegram, feedsChan <-chan []entities.FeedItem, postsChan <-chan entities.TelegramPost) *Component {
	return &Component{
		config:    config,
		telegram:  telegram,
		feedsChan: feedsChan,
		postsChan: postsChan,
	}
}

func (c *Component) Start(ctx context.Context) {
	log := logger.GetLoggerComponent(ctx, name)

	go func(log *zap.Logger) { // receive feeds
		for posts := range c.feedsChan {
			message, err := c.formatFeedPosts(ctx, posts)
			if err != nil {
				log.Error("error while formatting feed posts", zap.Error(err))
				continue
			}

			err = c.telegram.PostMessageHTML(ctx, message)
			if err != nil {
				log.Error("error while sending feed digest to telegram", zap.Error(err))
			}
		}
	}(log)

	go func(log *zap.Logger) { // recevie posts
		for post := range c.postsChan {
			err := c.telegram.PostMessageHTML(ctx, post)
			if err != nil {
				log.Error("error while sending post to telegram", zap.Error(err))
			}
		}
	}(log)
}
