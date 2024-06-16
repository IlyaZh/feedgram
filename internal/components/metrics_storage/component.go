package metrics_storage

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/IlyaZh/feedsgram/internal/configs"
	"github.com/IlyaZh/feedsgram/internal/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var name = "MetricsStorage"

type MetricsStorage interface {
	Start(ctx context.Context, config configs.MetricsStorage)
	LinkPosted(amount uint)
	LinkPostError(amount uint)
	SourceSyncDone(name string)
	SyncDone()
}

type Component struct {
	metrics metricsStorage
}

var impl *Component

func NewMetricsStorage() MetricsStorage {
	if impl == nil {
		impl = &Component{
			metrics: initMetricsStorage(),
		}
	}
	return impl
}

func (c *Component) Start(ctx context.Context, config configs.MetricsStorage) {
	ctx = logger.CreateSpan(ctx, &name, "Start")
	log := logger.GetLoggerComponent(ctx, name)
	log.Info("Component starting")
	go func(ctx context.Context, log *zap.Logger) {
		http.Handle(metricsHandler, promhttp.Handler())
		port := config["port"]
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		log.Info(fmt.Sprintf("Metrics available on handler \"%s\" and port = %d", metricsHandler, port))
	}(ctx, log)

	computePeriodRaw, ok := config["compute_period_sec"]
	if !ok || computePeriodRaw == 0 {
		computePeriodRaw = defaultComputePeriod
	}
	computePeriod := time.Second * time.Duration(computePeriodRaw.(int))

	go c.computeLastSyncTime(ctx, computePeriod)
	go c.computeLastSyncSourceMapTimes(ctx, computePeriod)

	log.Info("Component has started")
}
