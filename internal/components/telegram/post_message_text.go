package telegram

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
	"google.golang.org/appengine/log"
)

func (c *Component) PostMessageHTML(ctx context.Context, message entities.TelegramPost) error {
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
