package refresh

import (
	"github.com/RangelReale/trcache"
)

// Option

//troptgen:root
type options[K comparable, V any, RD any] interface {
	trcache.Options[K, V]
	trcache.MetricsOptions[K, V]
	trcache.DefaultRefreshOptions[K, V, RD]
	trcache.CallDefaultRefreshOptions[K, V]
}

// Refresh options

//troptgen:refresh
type refreshOptions[K comparable, V any, RD any] interface {
	trcache.RefreshOptions[K, V, RD]
}

//go:generate troptgen
