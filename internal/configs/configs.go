package configs

import (
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"time"
)

type Config struct {
	WebServer WebServer
	Telegram  Telegram
	Postgres  Postgres
	RssReader RssReader
}

type Telegram struct {
	Token          string
	Limit          *int
	Timeout        *int
	AllowedChatIds map[int64]struct{}
}

type WebServer struct {
	Port    int
	Timeout time.Duration
}

type Postgres struct {
	Host               *string `yaml:"host"`
	User               string  `yaml:"user"`
	Password           string  `yaml:"password"`
	Port               *int    `yaml:"port"`
	Database           string  `yaml:"database"`
	SslMode            *string `yaml:"sslmode"`
	MaxOpenConnections *int    `yaml:"max_open_connections"`
	MaxIdleConnections *int    `yaml:"max_idle_connections"`
	Limit              int     `yaml:"limit"`
}

type RssReader struct {
	Timeout time.Duration
}

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
	c.Telegram = Telegram{
		Token:          secdist.Telegram.Token,
		Limit:          raw.Telegram.Limit,
		Timeout:        raw.Telegram.Timeout,
		AllowedChatIds: allowedChatIds,
	}

	c.Postgres = Postgres{
		Host:               &secdist.Postgres.Host,
		User:               secdist.Postgres.User,
		Password:           secdist.Postgres.Password,
		Port:               &secdist.Postgres.Port,
		Database:           secdist.Postgres.Database,
		SslMode:            raw.Postgres.SslMode,
		MaxOpenConnections: raw.Postgres.MaxOpenConnections,
		MaxIdleConnections: raw.Postgres.MaxIdleConnections,
		Limit:              raw.Postgres.Limit,
	}

	c.RssReader = RssReader{
		Timeout: time.Duration(raw.RssReader.Timeout) * time.Second,
	}
	return nil
}

type configRaw struct {
	WebServer webServerRaw `yaml:"web_server"`
	Telegram  telegramRaw  `yaml:"telegram"`
	Postgres  postgresRaw  `yaml:"postgres"`
	RssReader rssReaderRaw `yaml:"rss_reader"`
}

type webServerRaw struct {
	Port    int `yaml:"port"`
	Timeout int `yaml:"timeout"`
}

type telegramRaw struct {
	Token          string  `yaml:"token"`
	Limit          *int    `yaml:"limit"`
	Timeout        *int    `yaml:"timeout"`
	AllowedChatIds []int64 `yaml:"allowed_chats_id"`
}

type postgresRaw struct {
	Host               *string `yaml:"host"`
	User               string  `yaml:"user"`
	Password           string  `yaml:"password"`
	Port               *int    `yaml:"port"`
	Database           string  `yaml:"database"`
	SslMode            *string `yaml:"sslmode"`
	MaxOpenConnections *int    `yaml:"max_open_connections"`
	MaxIdleConnections *int    `yaml:"max_idle_connections"`
	Limit              int     `yaml:"limit"`
}

type rssReaderRaw struct {
	Timeout int `yaml:"timeout"`
}
