package rss_reader

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func (c *Component) ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink *string) (entities.Feed, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(string(link), ctx)
	if err != nil {
		return entities.Feed{}, err
	}

	items := make([]entities.FeedItem, 0, len(feed.Items))
	config := c.config.GetValues().RssReader

	for _, item := range feed.Items {
		if item == nil {
			continue
		}
		hasOnePost := len(items) != 0

		if newerThan != nil && item.PublishedParsed == nil && item.PublishedParsed.Before(*newerThan) && (hasOnePost && config.PostsSettings.NewFeeds.AtLeastOncePost) {
			break
		}
		if lastPostLink != nil && item.Link == *lastPostLink {
			break
		}
		if lastPostLink == nil && newerThan == nil && config.PostsSettings.NewFeeds.AtLeastOncePost && hasOnePost {
			break
		}
		items = append(items, transformItem(item))
	}

	parsedFeed := entities.Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        entities.Link(feed.Link),
		FeedLink:    feed.FeedLink,
		UpdatedAt:   feed.UpdatedParsed,
		Items:       items,
	}

	return parsedFeed, nil
}
