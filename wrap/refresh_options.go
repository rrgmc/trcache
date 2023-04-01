package wrap

import (
	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type WrapRefreshOptions[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultRefreshOptions[K, V]
	OptDefaultRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V])
}

// Cache refresh options

// +troptgen refresh
type WrapRefreshRefreshOptions[K comparable, V any] interface {
	trcache.IsRefreshOptions
	trcache.RefreshOptions[K, V]
}

//go:generate troptgen -prefix Wrap
