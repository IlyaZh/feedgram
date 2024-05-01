package news_checker

import (
	"context"
	"sync"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

func (c *Component) Execute() {
	oneMore := true

	config := c.config.GetValues().NewsChecker
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.Timeout)
	defer cancelFunc()

	posts := make([]entities.FeedItem, 0, 2*config.ChunkSize)

	for oneMore {
		isActive := true
		sources, hasNext, err := c.storage.GetSources(ctx, nil, &isActive, &config.ChunkSize)
		if err != nil {
			log.Errorf("error while getting sources in periodic")
			return
		}
		oneMore = hasNext

		var wg sync.WaitGroup
		rssConfig := c.config.GetValues().RssReader
		responsesChannel := make(chan entities.RssResponse, len(sources)+1)
		// request to every source asynchroniously
		// responses will be stored in responsesChannel
		for _, source := range sources {
			wg.Add(1)

			go func(source entities.Source) {
				defer wg.Done()

				items, err := c.requestFeed(ctx, source)
				if err != nil {
					log.Warnf("getting feed has failed (id=%d, url %s): %s", source.Id, source.URL, err.Error())
					return
				}
				if items == nil {
					log.Infof("source (id=%d, link=%s) has no new posts", source.Id, source.URL)
					return
				}

				for i, item := range *items {
					log.Debugf("Put to channel %d, %s %s", i, item.Link, item.PublishedAt.String())
					responsesChannel <- entities.RssResponse{
						FeedItem: item,
						Id:       source.Id,
					}
				}
			}(source)
		}

		// channel reading goroutine
		var wgRead sync.WaitGroup
		respones := make([]entities.RssResponse, 0, rssConfig.PostsSettings.MaxPostsPerFeed*len(sources))
		wgRead.Add(1)
		go func() {
			defer wgRead.Done()
			for item := range responsesChannel {
				respones = append(respones, item)
			}
		}()

		// wait until requets will ended
		wg.Wait()
		close(responsesChannel)
		// wait until channel will read
		wgRead.Wait()

		// get responses from responsesChannel
		now := time.Now()
		updatedFeeds := make([]entities.UpdateSource, 0, len(respones))
		for _, response := range respones {
			updatedFeeds = append(updatedFeeds, entities.UpdateSource{
				Id:           response.Id,
				LastPostLink: response.Link,
				LastPostedAt: response.PublishedAt,
				LastSyncAt:   &now,
			})
			posts = append(posts, response.FeedItem)
		}

		if len(updatedFeeds) == 0 {
			continue
		}
		err = c.storage.UpdateSources(ctx, updatedFeeds)
		if err != nil {
			log.Errorf("Updating sources failed")
			return
		}
	}
	if len(posts) > 0 {
		c.out <- posts
	}
}
