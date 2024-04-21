package message_dispatcher

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/transformer"
	"github.com/labstack/gommon/log"
)

func (c *Component) handler_link(ctx context.Context, link entities.Link) {
	feed, err := c.rss_reader.ReadFeed(ctx, link)
	if err != nil {
		log.Errorf("Error in handler_link: %s", err.Error())
		panic(err)
	}

	id, err := c.storage.UpsertSource(ctx, transformer.Feed2Source(feed))
	if err != nil {
		log.Errorf("Error while upserting source: %s", err.Error())
		panic(err)
	}

	config := c.configs.GetValues().RssReader

	startTimepoint := time.Now().Add(-config.PostsSettings.NewFeeds.DaysInPast)
	sources, err := c.storage.GetSource(ctx, &id, nil, nil)
	if err != nil {
		panic(err)
	}
	if len(sources) != 1 {
		log.Errorf("Returned unexpected number of results. Expexted 1, returned %d", len(sources))
		panic(ErrUnexpectedResult)
	}
	source := sources[0]
	if source.LastPostedAt != nil {
		startTimepoint = *source.LastPostedAt
	}
	if len(feed.Items) == 0 {
		log.Infof("No posts in feed: id = %d, link = %s", id, feed.Link)
		return
	}
	posts := make([]entities.FeedItem, 0, len(feed.Items))
	for _, post := range feed.Items {
		if post.PublishedAt == nil {
			posts = append(posts, post)
			break
		}
		if post.PublishedAt.After(startTimepoint) {
			posts = append(posts, post)
		}
	}
	if len(posts) == 0 && config.PostsSettings.NewFeeds.AtLeastOncePost {
		posts = append(posts, feed.Items[0])
	}

}
