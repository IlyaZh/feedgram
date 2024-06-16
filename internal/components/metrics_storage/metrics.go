package metrics_storage

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metricsStorage struct {
	lastSyncMinutes        prometheus.Gauge
	lastSyncSourceMapHours *prometheus.GaugeVec
	postedLinksCount       prometheus.Counter
	linkPostErrors         prometheus.Counter
}

func initMetricsStorage() metricsStorage {
	return metricsStorage{
		lastSyncMinutes: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "last_sync_time_minutes",
			Help: "The time interval in minutes between the last time sync was done and now",
		}),
		lastSyncSourceMapHours: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "last_sync_time_minutes_sources_map",
			Help: "The time interval in minutes between the last time some source was done and now",
		}, []string{}),
		postedLinksCount: promauto.NewCounter(prometheus.CounterOpts{
			Name: "posted_to_telegram_links_count",
			Help: "Amount of links posted to telegram",
		}),
		linkPostErrors: promauto.NewCounter(prometheus.CounterOpts{
			Name: "posted_to_telegram_links_error",
			Help: "Amount of error occured while trying to post link to telegram",
		}),
	}
}
