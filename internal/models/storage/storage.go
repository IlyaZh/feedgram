package taskstorage

import (
	"database/sql"
	"log"

	"github.com/IlyaZh/feedsgram/internal/db"
	dbDriver "github.com/IlyaZh/feedsgram/internal/db"
	configModel "github.com/IlyaZh/feedsgram/internal/models/config"
)

const (
	DEFAULT_LIMIT      = 100
	QUERY_GET_CHANNELS = `SELECT 
		channel_tg_id, added_by_tg_id, sources, title, is_activated, created_at, updated_at, deleted_at 
		FROM channels
		WHERE (channel_tg_id > $1 OR $1 IS NULL)
		AND (is_active = $2 OR $2 IS NULL)
		LIMIT $3`
)

type Storage struct {
	configs *configModel.Cache
	db      *dbDriver.Db
}

var ptr *Storage

func NewStorage(config *configModel.Cache, db *db.Db) *Storage {
	if ptr != nil {
		return ptr
	}
	ptr = &Storage{configs: config, db: db}
	return ptr
}

type SourcesCursor int64
type QuerySourcesParams struct {
	cursor     *SourcesCursor
	OnlyActive bool
}

type ChannelsCursor int64
type QueryChannelsParams struct {
	cursor     *ChannelsCursor
	OnlyActive bool
}

func (s *Storage) GetChannels(params QueryChannelsParams) ([]Channel, *ChannelsCursor, error) {
	limit := DEFAULT_LIMIT
	limitFromConfig := s.configs.GetValues().Components.Postgres.Limit
	if limitFromConfig != nil {
		limit = *limitFromConfig
	}
	var args map[string]interface{"limit": limit}
	if params.cursor != nil {
		args["cursor"] = *params.cursor
	}
	if params.OnlyActive {
		args["is_active"] = true
	}

	rows, err := s.db.Pool().NamedExec(QUERY_GET_CHANNELS, args)
	var channels []Channel
	if err != nil {
		log.Printf("GetChannels query failed, err: %v\n", err)
		return channels, nil, err
	}
	if rows.RowsAffected() == 0 {
		return channels, nil, nil
	}
	for rows.Next() {
		var c Channel
		err = rows.StructScan(&c)
		if err == nil {
			channels = append(channels, c)
		}
	}
	rows.Close()
	return channels, (*ChannelsCursor)(&channels[len(channels)-1].ChanngelTgId), nil
}
