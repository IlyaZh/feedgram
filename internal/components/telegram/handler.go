package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
)

func (c *Component) handler(ctx context.Context, output chan<- entities.Message) {
	for update := range c.updates {
		if !c.messageFilter(&update) {
			log.Info("message filtered")
			continue
		}
		var post tgbotapi.Message
		if update.Message != nil {
			post = *update.Message
		} else if update.ChannelPost != nil {
			post = *update.ChannelPost
		}

		links := parseLinks(post.Text)
		if len(links) == 0 {
			log.Info("links not found in message, skip")
			continue
		}

		for _, link := range links {
			output <- entities.NewMessageLink(link)
		}
	}
}
