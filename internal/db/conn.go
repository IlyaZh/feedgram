package db

import (
	"fmt"
	"github.com/IlyaZh/feedsgram/internal/caches/configs"

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

func CreateInstance(config *configs.Cache) *Db {
	if db != nil && db.dbx != nil {
		return db
	}
	db = &Db{}

	settings := config.GetValues().Postgres
	sslMode := defaultSSlMode
	if settings.SslMode != nil {
		sslMode = *settings.SslMode
	}
	port := defaultPort
	if settings.Port != nil {
		port = *settings.Port
	}
	host := defaultHost
	if settings.Host != nil {
		host = *settings.Host
	}

	var err error
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", settings.User, settings.Password, host, port, sslMode)
	db.dbx, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		panic(err)
	}

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
