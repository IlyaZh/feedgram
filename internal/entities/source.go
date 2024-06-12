package entities

import (
	"time"
)

type Source struct {
	Id           int64      `db:"id"`
	URL          string     `db:"url"`
	Title        *string    `db:"title,omitempty"`
	Link         Link       `db:"link"`
	Description  string     `db:"description"`
	IsActive     bool       `db:"is_active"`
	LastPostLink *string    `db:"last_post_link,omitempty"`
	LastPostAt   *time.Time `db:"last_posted_at,omitempty"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at,omitempty"`
	DeletedAt    *time.Time `db:"deleted_at,omitempty"`
}

type UpdateSource struct {
	Id           int64      `db:"id"`
	LastPostLink Link       `db:"last_post_link"`
	LastPostedAt *time.Time `db:"last_posted_at,omitempty"`
	LastSyncAt   *time.Time `db:"last_sync_at,omitempty"`
}

type SourceShort struct {
	Title      string
	CreatedAt  time.Time
	LastPostAt *time.Time
}
