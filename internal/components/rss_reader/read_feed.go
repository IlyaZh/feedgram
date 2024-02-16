package rss_reader

import (
	"context"
	"fmt"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/mmcdole/gofeed"
)

func (c *Component) ReadFeed(ctx context.Context, link entities.Link) (entities.Feed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	link = entities.Link("http://feeds.twit.tv/twit.xml")
	feed, _ := fp.ParseURLWithContext(string(link), ctx)
	fmt.Printf("%v", feed)

	return entities.Feed{}, nil
}
