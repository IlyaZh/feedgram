package metrics_storage

import (
	"context"
	"time"

	"github.com/IlyaZh/feedsgram/internal/logger"
	"go.uber.org/zap"
)

var sourceLastSyncTimesMap = make(map[string]time.Time)

func (c *Component) SourceSyncDone(name string) {
	sourceLastSyncTimesMap[name] = time.Now()
}

func (c *Component) computeLastSyncSourceMapTimes(ctx context.Context, sleep time.Duration) {
	log := logger.GetLogger(ctx)
	for {
		now := time.Now()
		for label, timePoint := range sourceLastSyncTimesMap {
			duration := now.Sub(timePoint).Minutes()
			counter, err := c.metrics.lastSyncSourceMapHours.GetMetricWithLabelValues(label)
			if err != nil {
				log.Error("can't get metric counter", zap.String("metric", "lastSyncSourceMapHours"), zap.String("source", label))
				continue
			}
			counter.Set(duration)
		}
		time.Sleep(sleep)
	}
}
