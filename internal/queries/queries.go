package queries

import _ "embed"

var (
	//go:embed sql/get_sources.sql
	GetSources string

	//go:embed sql/upsert_source.sql
	UpsertSources string

	//go:embed sql/get_posts.sql
	GetPosts string

	//go:embed sql/upsert_posts.sql
	UpsertPosts string
)
