package chain

import (
	"github.com/RangelReale/trcache"
)

// Option

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.NameOptions[K, V]
	trcache.CallDefaultOptions[K, V]
	// OptGetStrategy sets the [GetStrategy] to use for the chain operation. The default is
	// [GetStrategyGetFirstSetPrevious].
	OptGetStrategy(getStrategy GetStrategy[K, V])
	// OptSetStrategy sets the [SetStrategy] to use for the chain operation. The default is
	// [SetStrategySetAll].
	OptSetStrategy(setStrategy SetStrategy[K, V])
	// OptDeleteStrategy sets the [DeleteStrategy] to use for the chain operation. The default is
	// [DeleteStrategyDeleteAll].
	OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V])
}

//troptgen:root name=rootRefreshOptions
type rootRefreshOptions[K comparable, V any] interface {
	trcache.MetricsOptions[K, V]
	trcache.DefaultRefreshOptions[K, V]
	trcache.CallDefaultRefreshOptions[K, V]
}

// Cache get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
	// OptSetOptions adds options to the [Cache.Set] call done after one of the [Cache.Get] function calls succeeds.
	OptSetOptions(options ...trcache.SetOption)
	optGetInfo(info *getInfo)
}

// Cache set options

//troptgen:set
type setOptions[K comparable, V any] interface {
	trcache.SetOptions[K, V]
}

// Cache delete options

//troptgen:delete
type deleteOptions[K comparable, V any] interface {
	trcache.DeleteOptions[K, V]
}

// Refresh options

//troptgen:refresh
type refreshOptions[K comparable, V any] interface {
	trcache.RefreshOptions[K, V]
}

//go:generate troptgen
