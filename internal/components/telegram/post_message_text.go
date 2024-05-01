package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/utils"
	"github.com/labstack/gommon/log"
)

func (c *Component) PostMessageHTML(ctx context.Context, message string) error {
	config := c.config.GetValues().Telegram

	log.Debugf("trying to send message: %s", message)
	_, err := c.api.Send(utils.CreateTelegramHTMLMessage(config.ChatForFeed, message))
	if err != nil {
		log.Errorf("telegram send error: %s", err.Error())
		return err
	}
	log.Infof("telegram message has succesfully sent")

	return err
}
