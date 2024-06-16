package news_checker

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"

	"go.uber.org/zap"
)

func (c *Component) requestFeed(ctx context.Context, source entities.Source, userAgent *string) (*[]entities.FeedItem, error) {
	feed, err := c.reader.ReadFeed(ctx, source.Link, source.LastPostAt, source.LastPostLink, userAgent)
	if err != nil {
		return nil, err
	}

	ctx = logger.CreateSpan(ctx, &name, "requestFeed")
	log := logger.GetLoggerComponent(ctx, name)
	log.Info("got photos from feed", zap.Int("count", len(feed.Items)), zap.Int64("source_id", source.Id), zap.String("source_link", string(source.Link)))
	return &feed.Items, err
}
