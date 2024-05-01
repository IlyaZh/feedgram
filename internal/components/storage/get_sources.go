package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
	"github.com/labstack/gommon/log"
)

func (c *Component) GetSources(ctx context.Context, id *int64, isActive *bool, limit *int) ([]entities.Source, bool, error) {
	settings := c.configs.GetValues().Mysql
	if limit == nil {
		limit = &settings.Limit
	}

	sources := make([]entities.Source, 0, *limit)
	err := c.db.Pool().SelectContext(ctx, &sources, queries.GetSources, id, isActive, *limit+1)
	if err != nil {
		log.Errorf("error occured whle gettings source. Error: %s", err.Error())
	}
	hasNext := len(sources) == *limit+1
	if !hasNext {
		return sources, hasNext, err
	}
	return sources[:len(sources)-1], hasNext, err
}
