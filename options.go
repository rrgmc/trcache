package trcache

import "time"

//
// Root options
//

//troptgen:root
type Options[K comparable, V any] interface {
}

//troptgen:root
type NameOptions[K comparable, V any] interface {
	// OptName sets the cache name.
	OptName(name string)
}

//troptgen:root
type MetricsOptions[K comparable, V any] interface {
	// OptMetrics sets the [Metrics] instance to call for metrics.
	OptMetrics(metrics Metrics, name string)
}

//troptgen:root name=refresh
type DefaultRefreshOptions[K comparable, V any] interface {
	// OptDefaultRefreshFunc sets the default refresh function to be used for [RefreshCache.GetOrRefresh].
	OptDefaultRefreshFunc(refreshFunc CacheRefreshFunc[K, V])
}

//troptgen:root
type CallDefaultOptions[K comparable, V any] interface {
	// OptCallDefaultGetOptions sets the default options to be appended to all [Cache.Get] calls.
	OptCallDefaultGetOptions(options ...GetOption)
	// OptCallDefaultSetOptions sets the default options to be appended to all [Cache.Set] calls.
	OptCallDefaultSetOptions(options ...SetOption)
	// OptCallDefaultDeleteOptions sets the default options to be appended to all [Cache.Delete] calls.
	OptCallDefaultDeleteOptions(options ...DeleteOption)
}

//troptgen:root
type CallDefaultRefreshOptions[K comparable, V any] interface {
	// OptCallDefaultRefreshOptions sets the default options to be appended to all [RefreshCache.GetOrRefresh] calls.
	OptCallDefaultRefreshOptions(options ...RefreshOption)
}

//
// Get options
//

//troptgen:get
type GetOptions[K comparable, V any] interface {
}

//
// Set options
//

//troptgen:set
type SetOptions[K comparable, V any] interface {
	// OptDuration sets the cache duration (TTL) to be used for this set call instead of the default.
	OptDuration(duration time.Duration)
}

//
// Delete options
//

//troptgen:delete
type DeleteOptions[K comparable, V any] interface {
}

//
// Refresh options
//

// RefreshFuncOptions is the options sent to the refresh function.
type RefreshFuncOptions struct {
	// Data is the custom data sent with [WithRefreshData].
	Data any
}

//troptgen:refresh
type RefreshOptions[K comparable, V any] interface {
	// OptData sets a custom data to be sent to the refresh function.
	OptData(data any)
	// OptFunc sets the refresh function for this call, possibly overriding a default one.
	OptFunc(refreshFunc CacheRefreshFunc[K, V])
	// OptGetOptions sets the options to be appended to all [Cache.Get] calls inside the refresh function.
	OptGetOptions(options ...GetOption)
	// OptSetOptions sets the options to be appended to all [Cache.Set] calls inside the refresh function.
	OptSetOptions(options ...SetOption)
}

//go:generate troptgen

//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Cache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name RefreshCache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Validator
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Codec
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name CacheRefreshFunc
