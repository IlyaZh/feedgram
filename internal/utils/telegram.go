package utils

import (
	"github.com/IlyaZh/feedsgram/internal/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateTelegramHTMLMessage(chatID int64, text entities.TelegramPost) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		ParseMode:             tgbotapi.ModeHTML,
		Text:                  string(text),
		DisableWebPagePreview: false,
	}
}

func CreateTelegramMarkdownMessage(chatID int64, text entities.TelegramPost) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		ParseMode:             tgbotapi.ModeMarkdown,
		Text:                  string(text),
		DisableWebPagePreview: false,
	}
}
