package utils

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/logger"

	"go.uber.org/zap"
)

type Executer interface {
	Period() time.Duration
	Execute(ctx context.Context)
	Finish()
}

var name string = "Periodic"

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
	ctx = logger.CreateSpan(ctx, &name, "Start")
	log := logger.GetLoggerComponent(ctx, name)
	defer func() {
		if r := recover(); r != nil {
			log.Error("Periodic has panicked", zap.Any("panic", r))
		}
	}()

	go c.handler(ctx)
}

func (c *periodic) handler(ctx context.Context) {
	ctx = logger.CreateSpan(ctx, &name, "handler")
	log := logger.GetLoggerComponent(ctx, name)
	for {
		period := c.exec.Period()
		log.Info("Set timer for period", zap.String("period", period.Abs().String()))
		timer := time.NewTimer(period)
		select {
		case <-ctx.Done():
			c.exec.Finish()
			log.Info("Periodic finished through context")
			return
		case <-timer.C:
			c.exec.Execute(ctx)
		}
	}
}
