package utils

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
)

type Executer interface {
	Period() time.Duration
	Execute()
	Finish()
}

type periodic struct {
	name string
	exec Executer
	done chan bool
}

func NewPeriodic(name string, executer Executer) *periodic {
	return &periodic{
		name: name,
		exec: executer,
		done: make(chan bool),
	}
}

func (c *periodic) Start(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Warnf("Periodic \"%s\" has panicked: %s", c.name, r)
		}
	}()

	go c.handler(ctx)
}

func (c *periodic) handler(ctx context.Context) {
	for {
		period := c.exec.Period()
		log.Infof("Set timer for period: %s", period.Abs().String())
		timer := time.NewTimer(period)
		select {
		case <-ctx.Done():
			c.exec.Finish()
			log.Infof("Periodic \"%s\" finished through context", c.name)
			return
		case <-timer.C:
			c.exec.Execute()
		}
	}
}
