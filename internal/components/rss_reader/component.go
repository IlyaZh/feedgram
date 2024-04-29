package rss_reader

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
)

type RssReader interface {
	ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink *string) (entities.Feed, error)
}

type Component struct {
	config    configs.ConfigsCache
	sanitizer utils.Sanitizer
}

func NewRssReader(configsCache configs.ConfigsCache) *Component {
	return &Component{
		config:    configsCache,
		sanitizer: utils.NewSanitizer(),
	}
}
