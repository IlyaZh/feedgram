package message_dispatcher

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
)

func (c *Component) dispatch(ctx context.Context) {
	ctx = logger.CreateSpan(ctx, &name, "dispatch")
	ctx = logger.CreateTrace(ctx)
	for message := range c.input {
		switch message.Type {
		case entities.MESSAGE_TYPE_LINK:
			c.handler_link(ctx, *message.Link)
		case entities.MESSAGE_TYPE_COMMAND:
			c.handler_command(ctx, *message.Command)
		}
	}
}
