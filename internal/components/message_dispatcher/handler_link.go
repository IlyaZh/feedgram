package message_dispatcher

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/transformer"

	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

func (c *Component) handler_link(ctx context.Context, link entities.Link) {
	log := logger.GetLogger(ctx)

	now := time.Now()
	feed, err := c.rss_reader.ReadFeed(ctx, link, &now, nil)
	if err != nil {
		if err == gofeed.ErrFeedTypeNotDetected {
			log.Info("Feed type is not detected. Skip")
			return
		} else {
			log.Error("Error in handler_link", zap.Error(err))
		}
	}

	feed.Link = link // because sometimes there is a wrong (maybe old) URL at feed
	_, err = c.storage.UpsertSource(ctx, transformer.Feed2Source(feed))
	if err != nil {
		log.Error("Error while upserting source", zap.Error(err))
	}

}
