package telegram

import (
	"log"

	"github.com/IlyaZh/feedgram/internal/entities"
	"google.golang.org/appengine/log"
)

type Telegram struct {
	token   string
	limit   int
	timeout int
	offset  int
}

const (
	DEFAULT_TIMEOUT = 60
	DEFAULT_LIMIT   = 100
)

func NewTelegram(settings entities.Telegram) Telegram {
	limit := DEFAULT_LIMIT
	if settings.Limit != nil {
		limit = *settings.Limit
	}

	timeout := DEFAULT_TIMEOUT
	if settings.Timeout != nil {
		timeout = *settings.Timeout
	}

	return Telegram{token: settings.Token, limit: limit, timeout: timeout, offset: 0}
}

func (c *Telegram) GetLimit() int {

}

func (c *Telegram) SetLimit(limit int) {
	if limit < 1 {
		log.Panicf("Trying to set limit value is less than 1: actual value = %d", limit)
	}
	c.limit = limit
}

func (c *Telegram) GetTimeout() int {
	return c.timeout
}

func (c *Telegram) SetTimeout(timeout int) {
	if timeout < 1 {
		log.Panicf("Trying to set timeout value is less than 1: actual value = %d", timeout)
	}
	c.timeout = timeout
}
