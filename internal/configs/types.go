package configs

import "time"

type Config struct {
	Telegram  Telegram
	Mysql     Mysql
	RssReader RssReader
}

type configRaw struct {
	Telegram  telegramRaw  `yaml:"telegram"`
	Mysql     mysqlRaw     `yaml:"mysql"`
	RssReader rssReaderRaw `yaml:"rss_reader"`
}

type Telegram struct {
	Token          string
	UseWebhook     bool
	Limit          *int
	Timeout        *int
	AllowedChatIds map[int64]struct{}
}

type telegramRaw struct {
	Token          string  `yaml:"token"`
	UseWebhook     *bool   `yaml:"use_webhool"`
	Limit          *int    `yaml:"limit"`
	Timeout        *int    `yaml:"timeout"`
	AllowedChatIds []int64 `yaml:"allowed_chats_id"`
}

type Mysql struct {
	Host               *string
	User               string
	Password           string
	Port               *int
	Database           string
	MaxOpenConnections *int
	MaxIdleConnections *int
	Limit              int
}

type mysqlRaw struct {
	Host               *string `yaml:"host"`
	User               string  `yaml:"user"`
	Password           string  `yaml:"password"`
	Port               *int    `yaml:"port"`
	Database           string  `yaml:"database"`
	MaxOpenConnections *int    `yaml:"max_open_connections"`
	MaxIdleConnections *int    `yaml:"max_idle_connections"`
	Limit              int     `yaml:"limit"`
}

type RssReader struct {
	Timeout       time.Duration
	BufferSize    int
	PostsSettings rssReaderPostsSettings
}

type rssReaderRaw struct {
	Timeout      int                       `yaml:"timeout"`
	BufferSize   *int                      `yaml:"buffer_size,omitempty"`
	PostSettings rssReaderPostsSettingsRaw `yaml:"posts_settings"`
}

type rssReaderPostsSettingsRaw struct {
	NewFeeds rssReaderPostsSettingsNewFeedsRaw `json:"new_feeds"`
}

type rssReaderPostsSettings struct {
	NewFeeds rssReaderPostsSettingsNewFeeds
}

type rssReaderPostsSettingsNewFeedsRaw struct {
	DaysInPast      int  `yaml:"days_in_past"`
	AtLeastOncePost bool `yaml:"at_least_once_post"`
}

type rssReaderPostsSettingsNewFeeds struct {
	DaysInPast      time.Duration
	AtLeastOncePost bool
}
