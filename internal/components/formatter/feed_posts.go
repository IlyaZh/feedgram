package formatter

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

func (c *Component) FeedPosts(ctx context.Context, posts []entities.FeedItem) (message string, err error) {
	config := c.config.GetValues().Formatter

	return
}
