package news_checker

import (
	config "github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/metrics_storage"
	"github.com/IlyaZh/feedsgram/internal/components/rss_reader"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/utils"
)

var name string = "NewsSender"

//go:generate mockgen -source component.go -package mocks -destination mocks/component.go
type NewsChecker interface {
	utils.Executer
}

type Component struct {
	config  config.ConfigsCache
	storage storage.Storage
	reader  rss_reader.RssReader
	out     chan<- []entities.FeedItem
	metrics metrics_storage.MetricsStorage
}

func NewNewsChecker(config config.ConfigsCache, outChannel chan<- []entities.FeedItem, storage storage.Storage, metricsStorage metrics_storage.MetricsStorage) *Component {
	return &Component{
		config:  config,
		storage: storage,
		reader:  rss_reader.NewRssReader(config),
		out:     outChannel,
		metrics: metricsStorage,
	}
}
