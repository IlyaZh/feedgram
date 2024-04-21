package message_dispatcher

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

func (c *Component) dispatch(ctx context.Context) {
	for message := range c.input {
		switch message.Type {
		case entities.MESSAGE_TYPE_LINK:
			c.handler_link(ctx, *message.Link)
		}
	}
}
