package configs

import (
	"time"

	"github.com/IlyaZh/feedsgram/internal/consts"

	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
)

func (c *Config) Scan(v []byte, secdist SecDist) error {
	var raw configRaw
	err := yaml.Unmarshal(v, &raw)
	if err != nil {
		log.Warn("config parse has failed")
		return err
	}

	c.WebServer = WebServer{
		Port:    raw.WebServer.Port,
		Timeout: time.Duration(raw.WebServer.Timeout) * time.Second,
	}

	allowedChatIds := make(map[int64]struct{})
	for _, id := range raw.Telegram.AllowedChatIds {
		allowedChatIds[id] = struct{}{}
	}

	useWebhook := false
	if raw.Telegram.UseWebhook != nil {
		useWebhook = *raw.Telegram.UseWebhook
	}

	c.Telegram = Telegram{
		Token:          secdist.Telegram.Token,
		UseWebhook:     useWebhook,
		Limit:          raw.Telegram.Limit,
		Timeout:        raw.Telegram.Timeout,
		AllowedChatIds: allowedChatIds,
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
			NewFeeds: rssReaderPostsSettingsNewFeeds{
				DaysInPast:      time.Duration(raw.RssReader.PostSettings.NewFeeds.DaysInPast*24) * time.Hour,
				AtLeastOncePost: raw.RssReader.PostSettings.NewFeeds.AtLeastOncePost,
			},
		},
	}
	return nil
}
