package refresh

import (
	"github.com/RangelReale/trcache"
)

// Option

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.MetricsOptions[K, V]
	trcache.DefaultRefreshOptions[K, V]
	trcache.CallDefaultRefreshOptions[K, V]
}

// Refresh options

//troptgen:refresh
type refreshOptions[K comparable, V any] interface {
	trcache.RefreshOptions[K, V]
}

//go:generate troptgen
