package rss_reader

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
)

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type RssReader interface {
	ReadFeed(ctx context.Context, link entities.Link, newerThan *time.Time, lastPostLink, userAgent *string) (entities.Feed, error)
}

var name string = "RssReader"

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
