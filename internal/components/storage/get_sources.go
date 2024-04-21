package storage

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
)

func (c *Component) GetSource(ctx context.Context, id *int64, isActive *bool, limit *int) ([]entities.Source, error) {
	settings := c.configs.GetValues().Mysql
	if limit != nil {
		limit = &settings.Limit
	}

	sources := make([]entities.Source, *limit)
	err := c.db.Pool().SelectContext(ctx, &sources, queries.GetSources, id, isActive, *limit)
	if err != nil {
		log.Printf("error occured whle gettings source. Error: %s", err.Error())
	}
	return sources, err
}
