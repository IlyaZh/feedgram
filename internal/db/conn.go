package db

import (
	"context"
	"fmt"

	"github.com/IlyaZh/feedsgram/internal/caches/configs"
	"go.uber.org/zap"

	"github.com/IlyaZh/feedsgram/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// http://jmoiron.github.io/sqlx/

const (
	defaultPort               = 5432
	defaultHost               = "localhost"
	defaultSSlMode            = "disable"
	defaultMaxOpenConnections = 5
	defaultMaxIdleConnections = 2
)

type Db struct {
	dbx *sqlx.DB
}

var db *Db

func CreateInstance(ctx context.Context, config *configs.Cache) *Db {
	ctx = logger.CreateSpan(ctx, nil, "CreateInstance")
	log := logger.GetLogger(ctx)
	if db != nil && db.dbx != nil {
		return db
	}
	db = &Db{}

	settings := config.GetValues().Mysql
	port := defaultPort
	if settings.Port != nil {
		port = *settings.Port
	}
	host := defaultHost
	if settings.Host != nil {
		host = *settings.Host
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", settings.User, settings.Password, host, port, settings.Database)
	log.Info("Trying to connect to db", zap.String("datebase", settings.Database), zap.String("host", host))
	db.dbx = sqlx.MustConnect("mysql", connectionString)
	log.Info("DB connect is OK", zap.String("datebase", settings.Database), zap.String("host", host))

	maxOpenConnections := defaultMaxOpenConnections
	if settings.MaxOpenConnections != nil {
		maxOpenConnections = *settings.MaxOpenConnections
	}
	db.dbx.SetMaxOpenConns(maxOpenConnections)
	maxIdleConnections := defaultMaxIdleConnections
	if settings.MaxIdleConnections != nil {
		maxIdleConnections = *settings.MaxIdleConnections
	}
	db.dbx.SetMaxIdleConns(maxIdleConnections)
	return db
}

func (p *Db) Pool() *sqlx.DB {
	return p.dbx
}
