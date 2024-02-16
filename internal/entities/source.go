package entities

import "time"

type Source struct {
	Id          int64      `db:"id"`
	URL         string     `db:"url"`
	Title       *string    `db:"title"`
	Link        string     `db:"link"`
	Description string     `db:"description"`
	IsActive    bool       `db:"is_active"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
