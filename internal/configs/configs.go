package configs

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/consts"
	"github.com/IlyaZh/feedsgram/internal/logger"

	"gopkg.in/yaml.v2"
)

var name string = "Config"

func (c *Config) Scan(ctx context.Context, v []byte, secdist SecDist) error {
	ctx = logger.CreateSpan(ctx, &name, "Scan")
	log := logger.GetLoggerComponent(ctx, name)
	var raw configRaw
	err := yaml.Unmarshal(v, &raw)
	if err != nil {
		log.Warn("config parse has failed")
		return err
	}

	allowedChatIds := make(map[int64]struct{})
	for _, id := range raw.Telegram.AllowedChatIds {
		allowedChatIds[id] = struct{}{}
	}

	useWebhook := false
	if raw.Telegram.UseWebhook != nil {
		useWebhook = *raw.Telegram.UseWebhook
	}

	messageWhenStart := false
	if raw.Telegram.MessageWhenStart != nil {
		messageWhenStart = *raw.Telegram.MessageWhenStart
	}
	c.Telegram = Telegram{
		Token:            secdist.Telegram.Token,
		BotID:            raw.Telegram.BotID,
		UseWebhook:       useWebhook,
		Limit:            raw.Telegram.Limit,
		Timeout:          raw.Telegram.Timeout,
		AllowedChatIds:   allowedChatIds,
		ChatForFeed:      raw.Telegram.ChatForFeed,
		MessageWhenStart: messageWhenStart,
	}

	c.Mysql = Mysql{
		Host:               &secdist.Mysql.Host,
		User:               secdist.Mysql.User,
		Password:           secdist.Mysql.Password,
		Port:               &secdist.Mysql.Port,
		Database:           secdist.Mysql.Database,
		MaxOpenConnections: raw.Mysql.MaxOpenConnections,
		MaxIdleConnections: raw.Mysql.MaxIdleConnections,
		Limit:              raw.Mysql.Limit,
	}

	rssReaderBufferSize := consts.ChannelDefaultBufferSize
	if raw.RssReader.BufferSize != nil {
		rssReaderBufferSize = *raw.RssReader.BufferSize
	}

	c.RssReader = RssReader{
		Timeout:    time.Duration(raw.RssReader.Timeout) * time.Second,
		BufferSize: rssReaderBufferSize,
		PostsSettings: rssReaderPostsSettings{
			MaxPostsPerFeed: raw.RssReader.PostSettings.MaxPostsPerFeed,
			NewFeeds: rssReaderPostsSettingsNewFeeds{
				AtLeastOncePost: raw.RssReader.PostSettings.NewFeeds.AtLeastOncePost,
			},
		},
		UserAgent: raw.RssReader.UserAgent,
	}

	newsCheckerBufferSize := consts.ChannelDefaultBufferSize
	if raw.NewsChecker.BufferSize != nil {
		newsCheckerBufferSize = *raw.NewsChecker.BufferSize
	}
	c.NewsChecker = NewsChecker{
		Period:     time.Duration(raw.NewsChecker.PeriodMin) * time.Minute,
		BufferSize: newsCheckerBufferSize,
		Timeout:    time.Duration(raw.NewsChecker.TimeoutMs) * time.Millisecond,
		ChunkSize:  raw.NewsChecker.ChunkSize,
	}

	c.Formatter = make(Formatter)
	for k, v := range raw.Formatter {
		c.Formatter[k] = FormatterItem{
			Header: v.Header,
			Loop:   v.Loop,
			Footer: v.Footer,
		}
	}

	c.MetricsStorage = MetricsStorage(raw.MetricsStorage)

	return nil
}
