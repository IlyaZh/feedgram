package news_checker

import (
	"context"
	"sync"
	"time"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"

	"go.uber.org/zap"
)

func (c *Component) Execute(ctx context.Context) {
	ctx = logger.CreateSpan(ctx, &name, "Execute")
	log := logger.GetLoggerComponent(ctx, name)
	oneMore := true

	config := c.config.GetValues().NewsChecker
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.Timeout)
	defer cancelFunc()

	posts := make([]entities.FeedItem, 0, 2*config.ChunkSize)

	for oneMore {
		isActive := true
		sources, hasNext, err := c.storage.GetSources(ctx, nil, &isActive, &config.ChunkSize)
		if err != nil {
			log.Error("error while getting sources in periodic", zap.Error(err))
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
					log.Warn("getting feed process has failed", zap.Int64("source_id", source.Id), zap.String("source_url", source.URL), zap.Error(err))
					return
				}
				if items == nil {
					log.Info("there are no new posts in source", zap.Int64("source_id", source.Id), zap.String("source_url", source.URL))
					return
				}

				for i, item := range *items {
					log.Debug("Put post to channel", zap.Int("idx", i), zap.String("item_link", string(item.Link)), zap.Timep("item_published_at", item.PublishedAt))
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
			log.Error("Updating sources failed", zap.Error(err))
			return
		}
	}
	if len(posts) > 0 {
		c.out <- posts
	}
}
