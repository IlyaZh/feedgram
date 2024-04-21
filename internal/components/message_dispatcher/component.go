package message_dispatcher

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/components/rss_reader"
	"github.com/IlyaZh/feedsgram/internal/components/storage"
	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/labstack/gommon/log"
)

type MessageDispatcher interface {
	Start()
}

type Component struct {
	configs    *configs.Cache
	storage    storage.Storage
	rss_reader rss_reader.RssReader
	input      <-chan entities.Message
}

var dispatcher *Component

func NewMessageDispatcher(config *configs.Cache, storage storage.Storage, input <-chan entities.Message) *Component {
	if dispatcher == nil {
		dispatcher = &Component{
			configs:    config,
			storage:    storage,
			rss_reader: rss_reader.NewRssReader(config),
			input:      input,
		}
	}
	return dispatcher
}

func (c *Component) Start(ctx context.Context) {
	log.Info("Dispatcher start")
	go c.dispatch(ctx)
	log.Info("Dispatcher has started")
}
