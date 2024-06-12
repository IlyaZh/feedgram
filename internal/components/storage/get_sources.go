package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/queries"

	"go.uber.org/zap"
)

func (c *Component) GetSources(ctx context.Context, id *int64, isActive *bool, limit *int) ([]entities.Source, bool, error) {
	ctx = logger.CreateSpan(ctx, &name, "GetSources")
	log := logger.GetLoggerComponent(ctx, name)
	settings := c.configs.GetValues().Mysql
	if limit == nil {
		limit = &settings.Limit
	}

	sources := make([]entities.Source, 0, *limit)
	err := c.db.Pool().SelectContext(ctx, &sources, queries.GetSources, id, isActive, *limit+1)
	if err != nil {
		log.Error("error occured whle gettings source", zap.Error(err))
	}
	hasNext := len(sources) == *limit+1
	if !hasNext {
		return sources, hasNext, err
	}
	return sources[:len(sources)-1], hasNext, err
}
