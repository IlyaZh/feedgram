package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
	"github.com/labstack/gommon/log"
)

func (c *Component) UpdateSources(ctx context.Context, sources []entities.UpdateSource) error {
	_, err := c.db.Pool().NamedExecContext(ctx, queries.UpdateSources, sources)
	if err != nil {
		log.Errorf("Error wile prepare statement in update sources method. Error: %s", err.Error())
		return err
	}
	return nil
}
