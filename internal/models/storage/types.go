package taskstorage

import (
	"database/sql"
	"time"
)

type Source struct {
	Id          int64          `db:"id"`
	Url         string         `db:"url"`
	Title       sql.NullString `db:"title"`
	Link        string         `db:"link"`
	Description string         `db:"description"`
	IsActive    bool           `db:"is_acive"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
}

type Post struct {
	Id             int64                  `db:"id"`
	SourceId       int64                  `db:"source_id"`
	Title          string                 `db:"title"`
	Body           sql.NullString         `db:"body"`
	Link           string                 `db:"link"`
	Author         sql.NullString         `db:"author"`
	HasReaded      bool                   `db:"has_readed"`
	AdditionalInfo map[string]interface{} `db:"additional_info"`
	PostedAt       sql.NullTime           `db:"posted_at"`
	CreatedAt      time.Time              `db:"created_at"`
}

type Channel struct {
	ChanngelTgId int64        `db:"channel_tg_id"`
	AddedByTgId  int64        `db:"added_by_tg_id"`
	Sources      []int64      `db:"sources"`
	IsActive     bool         `db:"is_acive"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}
