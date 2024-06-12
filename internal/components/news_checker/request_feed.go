package news_checker

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

func (c *Component) requestFeed(ctx context.Context, source entities.Source) (*[]entities.FeedItem, error) {
	feed, err := c.reader.ReadFeed(ctx, source.Link, source.LastPostAt, source.LastPostLink)
	if err != nil {
		return nil, err
	}

	log.Infof("Got %d posts from feed (id=%d): %s", len(feed.Items), source.Id, source.Link)
	return &feed.Items, err
}
