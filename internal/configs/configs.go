package configs

import (
	"time"

	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	WebServer WebServer
	Telegram  Telegram
	Mysql     Mysql
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

type Mysql struct {
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

	c.RssReader = RssReader{
		Timeout: time.Duration(raw.RssReader.Timeout) * time.Second,
	}
	return nil
}

type configRaw struct {
	WebServer webServerRaw `yaml:"web_server"`
	Telegram  telegramRaw  `yaml:"telegram"`
	Mysql     mysqlRaw     `yaml:"mysql"`
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

type rssReaderRaw struct {
	Timeout int `yaml:"timeout"`
}
