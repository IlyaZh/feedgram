package configs

import "time"

type Config struct {
	Telegram    Telegram
	Mysql       Mysql
	RssReader   RssReader
	NewsChecker NewsChecker
	Formatter   Formatter
}

type configRaw struct {
	Telegram     telegramRaw    `yaml:"telegram"`
	Mysql        mysqlRaw       `yaml:"mysql"`
	RssReader    rssReaderRaw   `yaml:"rss_reader"`
	NewsChecker  newsCheckerRaw `yaml:"news_checker"`
	FormatterRaw formatterRaw   `yaml:"formatter"`
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
	MaxPostsPerFeed int                               `yaml:"max_post_per_feed"`
	NewFeeds        rssReaderPostsSettingsNewFeedsRaw `yaml:"new_feeds"`
}

type rssReaderPostsSettings struct {
	MaxPostsPerFeed int
	NewFeeds        rssReaderPostsSettingsNewFeeds
}

type rssReaderPostsSettingsNewFeedsRaw struct {
	AtLeastOncePost bool `yaml:"at_least_once_post"`
}

type rssReaderPostsSettingsNewFeeds struct {
	AtLeastOncePost bool
}

type newsCheckerRaw struct {
	PeriodMin  int64 `yaml:"period_min"`
	BufferSize *int  `yaml:"buffer_size,omitempty"`
	TimeoutMs  int64 `yaml:"timeout_ms"`
	ChunkSize  int   `yaml:"chunk_size"`
}

type NewsChecker struct {
	Period     time.Duration
	BufferSize int
	Timeout    time.Duration
	ChunkSize  int
}

type formatterRaw map[string]formatterItemRaw

type Formatter map[string]FormatterItem

type formatterItemRaw struct {
	Header string  `yaml:"header"`
	Loop   string  `yaml:"loop"`
	Footer *string `yaml:"footer,omitempty"`
}

type FormatterItem struct {
	Header string
	Loop   string
	Footer *string
}
