package message_dispatcher

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/transformer"
	"github.com/labstack/gommon/log"
	"github.com/mmcdole/gofeed"
)

func (c *Component) handler_link(ctx context.Context, link entities.Link) {
	now := time.Now()
	feed, err := c.rss_reader.ReadFeed(ctx, link, &now, nil)
	if err != nil {
		if err == gofeed.ErrFeedTypeNotDetected {
			log.Infof("Feed type is not detected. Skip")
			return
		} else {
			log.Errorf("Error in handler_link: %s", err.Error())
			panic(err)
		}
	}

	feed.Link = link // because sometimes there is a wrong (maybe old) URL at feed
	_, err = c.storage.UpsertSource(ctx, transformer.Feed2Source(feed))
	if err != nil {
		log.Errorf("Error while upserting source: %s", err.Error())
		panic(err)
	}

}
