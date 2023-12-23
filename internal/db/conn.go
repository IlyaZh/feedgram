package db

import (
	"fmt"

	"databases/sqlx"

	"github.com/IlyaZh/feedsgram/internal/models/config"
)

// http://jmoiron.github.io/sqlx/

const (
	DEFAULT_PORT                 = 5432
	DEFAULT_HOST                 = "localhost"
	DEFAULT_SSL_MODE             = "disable"
	DEFAULT_MAX_OPEN_CONNECTIONS = 5
	DEFAULT_MAX_IDLE_CONNECTIONS = 2
)

type Db struct {
	dbx *sqlx.DB
}

var ptr *Db

func CreateInstance(config *config.Cache) *Db {
	if ptr != nil && ptr.dbx != nil {
		return ptr
	}
	ptr = &Db{}

	settings := config.GetValues().Components.Postgres
	sslMode := DEFAULT_SSL_MODE
	if settings.SslMode != nil {
		sslMode = *settings.SslMode
	}
	port := DEFAULT_PORT
	if settings.Port != nil {
		port = *settings.Port
	}
	host := DEFAULT_HOST
	if settings.Host != nil {
		host = *settings.Host
	}

	var err error
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", settings.User, settings.Password, host, port, sslMode)
	ptr.dbx, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	max_open_connections := DEFAULT_MAX_OPEN_CONNECTIONS
	if settings.MaxOpenConnections != nil {
		max_open_connections = *settings.MaxOpenConnections
	}
	ptr.dbx.SetMaxOpenConns(max_open_connections)
	max_idle_connections := DEFAULT_MAX_IDLE_CONNECTIONS
	if settings.MaxIdleConnections != nil {
		max_idle_connections = *settings.MaxIdleConnections
	}
	ptr.dbx.SetMaxIdleConns(max_idle_connections)
}

func (p *Db) Pool() *sqlx.DB {
	return p.dbx
}
