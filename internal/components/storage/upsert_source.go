package storage

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
)

func (c *Component) UpsertSource(ctx context.Context, source entities.Source) error {
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpsertSources)
	if err != nil {
		log.Printf("Error while prepare statement in upsert source method. Error: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, source)
	return err
}
