package refresh

import (
	"github.com/RangelReale/trcache"
)

// Refresh options

//troptgen:refresh
type refreshOptions[K comparable, V any, RD any] interface {
	trcache.RefreshOptions[K, V, RD]
}

type defaultRefreshOptions[K comparable, V any, RD any] struct {
	callDefaultRefreshOptions []trcache.RefreshOption
	defaultRefreshFunc        trcache.CacheRefreshFunc[K, V, RD]
	metricsMetrics            trcache.Metrics
	metricsName               string
}

//go:generate troptgen
