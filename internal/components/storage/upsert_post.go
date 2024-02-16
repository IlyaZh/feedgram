package storage

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
)

func (c *Component) UpsertPost(ctx context.Context, post entities.Post) error {
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpsertPosts)
	if err != nil {
		log.Printf("Error while prepare statement in upsert posts method. Error: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, post)
	return err
}
