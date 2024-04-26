package rss_reader

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func (c *Component) ReadFeed(ctx context.Context, link entities.Link) (entities.Feed, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(string(link), ctx)
	if err != nil {
		return entities.Feed{}, err
	}

	parsedFeed := entities.Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		FeedLink:    feed.FeedLink,
		UpdatedAt:   feed.UpdatedParsed,
		Items:       transformItems(feed.Items),
	}

	return parsedFeed, nil
}
