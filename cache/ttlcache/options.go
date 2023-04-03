package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Root options

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.NameOptions[K, V]
	trcache.CallDefaultOptions[K, V]
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
}

//troptgen:root name=rootRefreshOptions
type optionsRefresh[K comparable, V any, RD any] interface {
	trcache.DefaultRefreshOptions[K, V, RD]
	trcache.MetricsOptions[K, V]
	trcache.CallDefaultRefreshOptions[K, V]
}

// Get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
	OptTouch(touch bool)
}

// Set options

//troptgen:set
type setOptions[K comparable, V any] interface {
	trcache.SetOptions[K, V]
}

// Delete options

//troptgen:delete
type deleteOptions[K comparable, V any] interface {
	trcache.DeleteOptions[K, V]
}

// Refresh options

//troptgen:refresh
type refreshOptions[K comparable, V any, RD any] interface {
	trcache.RefreshOptions[K, V, RD]
}

//go:generate troptgen
