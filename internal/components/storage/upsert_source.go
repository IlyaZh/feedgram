package storage

import (
	"context"

	"github.com/labstack/gommon/log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
)

func (c *Component) UpsertSource(ctx context.Context, source entities.Source) (id int64, err error) {
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpsertSources)
	if err != nil {
		log.Printf("Error while prepare statement in upsert source method. Error: %s", err.Error())
		return
	}

	result, err := stmt.ExecContext(ctx, source)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}
