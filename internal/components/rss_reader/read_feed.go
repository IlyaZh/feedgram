package rss_reader

import (
	"context"
	"sort"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
)

func (c *Component) ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink *string) (entities.Feed, error) {
	ctx = logger.CreateSpan(ctx, &name, "ReadFeed")
	log := logger.GetLoggerComponent(ctx, name)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(string(link), ctx)
	if err != nil {
		return entities.Feed{}, err
	}

	log.Debug("post is newer than another", zap.Timep("newer_than", newerThan), zap.Stringp("last_post_link", lastPostLink))

	sort.Slice(feed.Items, func(i, j int) bool {
		if feed.Items[i] == nil && feed.Items[j] == nil {
			return true
		}

		if feed.Items[i] == nil {
			return false
		}

		if feed.Items[j] == nil {
			return true
		}

		return feed.Items[i].PublishedParsed.After(*feed.Items[j].PublishedParsed)
	})
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
		items = append(items, transformItem(c.sanitizer, item))
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
