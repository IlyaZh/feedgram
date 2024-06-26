package configs

import "time"

type Config struct {
	Telegram       Telegram
	Mysql          Mysql
	RssReader      RssReader
	NewsChecker    NewsChecker
	Formatter      Formatter
	MetricsStorage MetricsStorage
}

type configRaw struct {
	Telegram       telegramRaw       `yaml:"telegram"`
	Mysql          mysqlRaw          `yaml:"mysql"`
	RssReader      rssReaderRaw      `yaml:"rss_reader"`
	NewsChecker    newsCheckerRaw    `yaml:"news_checker"`
	Formatter      formatterRaw      `yaml:"formatter"`
	MetricsStorage metricsStorageRaw `yaml:"metrics_storage"`
}

type Telegram struct {
	Token            string
	BotID            int64
	UseWebhook       bool
	Limit            *int
	Timeout          *int
	AllowedChatIds   map[int64]struct{}
	ChatForFeed      int64
	MessageWhenStart bool
}

type telegramRaw struct {
	Token            string  `yaml:"token"`
	BotID            int64   `yaml:"bot_id"`
	UseWebhook       *bool   `yaml:"use_webhool"`
	Limit            *int    `yaml:"limit"`
	Timeout          *int    `yaml:"timeout"`
	AllowedChatIds   []int64 `yaml:"allowed_chats_id"`
	ChatForFeed      int64   `yaml:"chat_for_feed"`
	MessageWhenStart *bool   `yaml:"message_when_start"`
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
	UserAgent     *string
}

type rssReaderRaw struct {
	Timeout      int                       `yaml:"timeout"`
	BufferSize   *int                      `yaml:"buffer_size,omitempty"`
	PostSettings rssReaderPostsSettingsRaw `yaml:"posts_settings"`
	UserAgent    *string                   `yaml:"user_agent`
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

type MetricsStorage map[string]interface{}

type metricsStorageRaw map[string]interface{}
