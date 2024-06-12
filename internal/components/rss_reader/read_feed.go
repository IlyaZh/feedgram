package rss_reader

import (
	"context"
	"sort"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
	"google.golang.org/appengine/log"
)

func (c *Component) ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink *string) (entities.Feed, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(string(link), ctx)
	if err != nil {
		return entities.Feed{}, err
	}

	dNewerThan := "nil"
	dLastPostLink := "nil"
	if newerThan != nil {
		dNewerThan = newerThan.Format(time.DateTime)
	}
	if lastPostLink != nil {
		dLastPostLink = *lastPostLink
	}
	log.Debugf("newerThan = %s, lastPostLink = %s", dNewerThan, dLastPostLink)

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
