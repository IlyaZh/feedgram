package message_sender

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/telegram"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type MeesageSender interface {
	Start(ctx context.Context)
}

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
	go func() { // receive feeds
		for posts := range c.feedsChan {
			message, err := c.formatFeedPosts(posts)
			if err != nil {
				log.Errorf("error while formatting feed posts: %s", err.Error())
				continue
			}

			err = c.telegram.PostMessageHTML(ctx, message)
			if err != nil {
				log.Errorf("error while sending feed digest to telegram: %s", err.Error())
			}
		}
	}()

	go func() { // recevie posts
		for post := range c.postsChan {
			err := c.telegram.PostMessageHTML(ctx, post)
			if err != nil {
				log.Errorf("error while sending post to telegram: %s", err.Error())
			}
		}
	}()
}
