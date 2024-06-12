package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/queries"

	"go.uber.org/zap"
)

func (c *Component) UpsertSource(ctx context.Context, source entities.Source) (id int64, err error) {
	ctx = logger.CreateSpan(ctx, &name, "UpsetSource")
	log := logger.GetLoggerComponent(ctx, name)
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpsertSources)
	if err != nil {
		log.Error("Error wile prepare statement in upsert source method", zap.Error(err))
		return
	}

	result, err := stmt.ExecContext(ctx, source)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}
