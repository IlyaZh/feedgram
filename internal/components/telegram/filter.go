package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
)

func (c *Component) messageFilter(msg *tgbotapi.Update) bool {
	if msg == nil || msg.ChannelPost == nil {
		return false
	}

	post := *msg.ChannelPost
	settings := c.config.GetValues().Telegram

	if len(settings.AllowedChatIds) > 0 {
		if post.Chat == nil {
			return false
		}
		if _, allowed := settings.AllowedChatIds[post.Chat.ID]; !allowed {
			log.Warnf("message from not allowed chat: id = %d", post.Chat.ID)
			return false
		}
	}

	if post.Text == "" {
		return false
	}

	return true
}
