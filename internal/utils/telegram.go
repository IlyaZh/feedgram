package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateTelegramHTMLMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		ParseMode:             tgbotapi.ModeHTML,
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

func CreateTelegramMarkdownMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		ParseMode:             tgbotapi.ModeMarkdown,
		Text:                  text,
		DisableWebPagePreview: false,
	}
}
