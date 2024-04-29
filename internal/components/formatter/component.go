package formatter

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/entities"
)

type Formater interface {
	FeedPosts(ctx context.Context, posts []entities.FeedItem) (string, error)
}

type Component struct {
	config configs.ConfigsCache
}

func NewFormatter(config configs.ConfigsCache) Formater {
	return &Component{
		config: config,
	}
}
