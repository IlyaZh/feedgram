package queries

import _ "embed"

var (
	//go:embed sql/get_sources.sql
	GetSources string

	//go:embed sql/upsert_source.sql
	UpsertSources string

	//go:embed sql/update_source.sql
	UpdateSource string
)
