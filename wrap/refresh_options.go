package wrap

import (
	"github.com/RangelReale/trcache"
)

// Option

//troptgen:root
type wrapRefreshOptions[K comparable, V any, RD any] interface {
	trcache.Options[K, V]
	trcache.MetricsOptions[K, V]
	trcache.DefaultRefreshOptions[K, V, RD]
	trcache.CallDefaultRefreshOptions[K, V]
}

// Cache refresh options

//troptgen:refresh
type wrapRefreshRefreshOptions[K comparable, V any, RD any] interface {
	trcache.RefreshOptions[K, V, RD]
}

//go:generate troptgen -prefix wrap
