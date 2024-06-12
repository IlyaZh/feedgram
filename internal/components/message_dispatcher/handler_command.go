package message_dispatcher

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"go.uber.org/zap"
)

type sourceMessage struct {
	Title      string
	LastPostAt *time.Time
	CreatedAt  time.Time
}

func (c *Component) handler_command(ctx context.Context, command entities.Command) {
	log := logger.GetLogger(ctx)

	switch string(command) {
	case CommandList:
		var id int64 = 0
		isActive := true
		limit := 10
		var hasNext bool = true
		var err error
		var sources []entities.Source

		for hasNext {
			sources, hasNext, err = c.storage.GetSources(ctx, &id, &isActive, &limit)
			if err != nil {
				log.Error("Error while upserting source", zap.Error(err))
			}
			messages := make([]sourceMessage, 0, limit)
			for _, source := range sources {
				msg := sourceMessage{
					Title:      *source.Title,
					LastPostAt: source.LastPostAt,
					CreatedAt:  source.CreatedAt,
				}
				if len(msg.Title) == 0 {
					msg.Title = string(source.Link)
				}
				messages = append(messages, msg)
			}
		}
	}

}
