package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Component) handler(ctx context.Context, output chan<- entities.Message) {
	for update := range c.updates {
		ctx = logger.CreateSpan(ctx, &name, "handler")
		log := logger.GetLoggerComponent(ctx, name)
		if !c.messageFilter(ctx, &update) {
			log.Info("message filtered")
			continue
		}
		var post tgbotapi.Message
		if update.Message != nil {
			post = *update.Message
		} else if update.ChannelPost != nil {
			post = *update.ChannelPost
		}

		if post.IsCommand() {
			output <- entities.NewMessageCommand(entities.Command(post.Command()))
		} else {
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
}
