package storage

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
