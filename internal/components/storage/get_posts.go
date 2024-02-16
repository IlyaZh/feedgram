package storage

import (
	"context"
	"log"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/queries"
)

func (c *Component) GetPosts(ctx context.Context, id *int, hasReaded *bool, limit *int) ([]entities.Post, error) {
	settings := c.configs.GetValues().Postgres
	if limit != nil {
		limit = &settings.Limit
	}

	posts := make([]entities.Post, *limit)
	err := c.db.Pool().SelectContext(ctx, &posts, queries.GetPosts, id, hasReaded, *limit)
	if err != nil {
		log.Printf("error occured whle gettings posts. Error: %s", err.Error())
	}
	return posts, err
}
