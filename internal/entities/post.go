package entities

import "time"

type Post struct {
	ID             int64                  `db:"id"`
	SourceId       int64                  `db:"source_id"`
	Title          string                 `db:"title"`
	Body           *string                `db:"body"`
	Link           string                 `db:"link"`
	Author         *string                `db:"author"`
	HasRead        bool                   `db:"has_read"`
	AdditionalInfo map[string]interface{} `db:"additional_info"`
	CreatedAt      time.Time              `db:"created_at"`
	PostedAt       *time.Time             `db:"posted_at"`
	UpdatedAt      *time.Time             `db:"updated_at"`
}
