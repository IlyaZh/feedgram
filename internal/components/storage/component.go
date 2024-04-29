package storage

import (
	"context"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/db"
	dbDriver "github.com/IlyaZh/feedsgram/internal/db"
	"github.com/IlyaZh/feedsgram/internal/entities"
)

//go:generate mockgen -source component.go -package mocks -destination mock/component.go
type Storage interface {
	UpsertSource(ctx context.Context, source entities.Source) (int64, error)
	GetSources(ctx context.Context, id *int64, isActive *bool, limit *int) ([]entities.Source, bool, error)
	UpdateSources(ctx context.Context, sources []entities.UpdateSource) error
}

// https://github.com/golang/mock

type Component struct {
	configs *configs.Cache
	db      *dbDriver.Db
}

var storage *Component

func NewStorage(config *configs.Cache, db *db.Db) *Component {
	if storage != nil {
		return storage
	}
	storage = &Component{configs: config, db: db}
	return storage
}
