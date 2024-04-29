package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
	"github.com/labstack/gommon/log"
)

func (c *Component) UpdateSources(ctx context.Context, sources []entities.UpdateSource) error {
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpdateSource)
	if err != nil {
		log.Errorf("Error while prepare statement in update sources method. Error: %s", err.Error())
		return err
	}
	for _, source := range sources {
		_, err := stmt.ExecContext(ctx, source)
		if err != nil {
			log.Errorf("Error while execute statement in update sources method. Error: %s", err.Error())
		}
	}
	return nil
}
