package rss_reader

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
)

type RssReader interface {
	ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink *string) (entities.Feed, error)
}

type Component struct {
	config configs.ConfigsCache
}

func NewRssReader(configsCache configs.ConfigsCache) *Component {
	return &Component{
		config: configsCache,
	}
}
