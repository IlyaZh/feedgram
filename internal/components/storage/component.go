package storage

import (
	"context"
	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"github.com/IlyaZh/feedsgram/internal/db"
	dbDriver "github.com/IlyaZh/feedsgram/internal/db"
	"github.com/IlyaZh/feedsgram/internal/entities"
)

type Storage interface {
	UpsertSource(ctx context.Context, source entities.Source) error
	GetSource(ctx context.Context, id *int, isActive *bool, limit *int) ([]entities.Source, error)
	UpsertPost(ctx context.Context, post entities.Post) error
	GetPosts(ctx context.Context, id *int, hasReaded *bool, limit *int) ([]entities.Post, error)
}

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
