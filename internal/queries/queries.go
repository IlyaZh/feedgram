package queries

import _ "embed"

var (
	//go:embed sql/get_sources.sql
	GetSources string

	//go:embed sql/upsert_source.sql
	UpsertSources string
)
