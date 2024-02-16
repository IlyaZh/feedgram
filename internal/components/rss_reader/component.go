package rss_reader

import (
	"context"
	"github.com/IlyaZh/feedsgram/internal/caches/configs"

	"github.com/IlyaZh/feedsgram/internal/entities"
)

type RssReader interface {
	ReadFeed(ctx context.Context, link entities.Link) (entities.Feed, error)
}

type Component struct {
	config *configs.Cache
}

func NewRssReader(configsCache *configs.Cache) *Component {
	return &Component{
		config: configsCache,
	}
}
