package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/entities"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/IlyaZh/feedsgram/internal/queries"

	"go.uber.org/zap"
)

func (c *Component) UpdateSources(ctx context.Context, sources []entities.UpdateSource) error {
	ctx = logger.CreateSpan(ctx, &name, "UpdateSources")
	log := logger.GetLoggerComponent(ctx, name)
	stmt, err := c.db.Pool().PrepareNamedContext(ctx, queries.UpdateSource)
	if err != nil {
		log.Error("Error while prepare statement in update sources method", zap.Error(err))
		return err
	}
	for _, source := range sources {
		_, err := stmt.ExecContext(ctx, source)
		if err != nil {
			log.Error("Error while execute statement in update sources method", zap.Error(err))
		}
	}
	return nil
}
