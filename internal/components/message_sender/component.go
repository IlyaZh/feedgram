package message_sender

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/telegram"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

type MeesageSender interface {
	Start(ctx context.Context)
}

type Component struct {
	config   configs.ConfigsCache
	telegram telegram.Telegram
	input    <-chan []entities.FeedItem
}

func NewMeesageSender(config configs.ConfigsCache, telegram telegram.Telegram, input <-chan []entities.FeedItem) MeesageSender {
	return &Component{
		config:   config,
		telegram: telegram,
		input:    input,
	}
}

func (c *Component) Start(ctx context.Context) {
	go func() {
		for posts := range c.input {
			message, err := c.formatFeedPosts(ctx, posts)
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
}
