package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
)

func (c *Component) messageFilter(msg *tgbotapi.Update) bool {
	if msg == nil || (msg.Message == nil && msg.ChannelPost == nil) {
		return false
	}

	var post tgbotapi.Message
	if msg.Message != nil {
		post = *msg.Message
	} else if msg.ChannelPost != nil {
		post = *msg.ChannelPost
	} else {
		log.Error("undefined type of post, skip")
		return false
	}
	settings := c.config.GetValues().Telegram

	if len(settings.AllowedChatIds) > 0 {
		if post.Chat == nil {
			log.Info("Chat is nil")
			return false
		}
		if _, allowed := settings.AllowedChatIds[post.Chat.ID]; !allowed {
			log.Warnf("message from not allowed chat: id = %d", post.Chat.ID)
			return false
		}
	}

	if post.Text == "" {
		log.Info("Message has no text")
		return false
	}

	return true
}
