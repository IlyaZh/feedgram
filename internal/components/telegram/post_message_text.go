package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"go.uber.org/zap"
)

func (c *Component) PostMessageHTML(ctx context.Context, message entities.TelegramPost) error {
	ctx = logger.CreateSpan(ctx, &name, "PostMessageHTML")
	log := logger.GetLoggerComponent(ctx, name)
	config := c.config.GetValues().Telegram

	log.Debug("trying to send message", zap.String("message", string(message)))
	_, err := c.api.Send(utils.CreateTelegramHTMLMessage(config.ChatForFeed, message))
	if err != nil {
		log.Error("telegram send error: %s", zap.Error(err))
		c.metrics.LinkPostError(1)
		return err
	}
	c.metrics.LinkPosted(1)
	log.Info("telegram message has succesfully sent")

	return err
}
