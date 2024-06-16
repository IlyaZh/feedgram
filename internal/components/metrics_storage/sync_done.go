package metrics_storage

import (
	"context"
	"time"
)

var lastSyncTime *time.Time

func (c *Component) SyncDone() {
	now := time.Now()
	lastSyncTime = &now
}

func (c *Component) computeLastSyncTime(ctx context.Context, delay time.Duration) {
	for {
		now := time.Now()
		if lastSyncTime == nil {
			lastSyncTime = &now
		}
		c.metrics.lastSyncMinutes.Set(float64(now.Sub(*lastSyncTime).Minutes()))
		time.Sleep(delay)
	}
}
