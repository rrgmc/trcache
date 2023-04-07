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
	OptName(name string)
}

//troptgen:root
type MetricsOptions[K comparable, V any] interface {
	OptMetrics(metrics Metrics, name string)
}

//troptgen:root name=refresh
type DefaultRefreshOptions[K comparable, V any] interface {
	OptDefaultRefreshFunc(refreshFunc CacheRefreshFunc[K, V])
}

//troptgen:root
type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOptions(options ...GetOption)
	OptCallDefaultSetOptions(options ...SetOption)
	OptCallDefaultDeleteOptions(options ...DeleteOption)
}

//troptgen:root
type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOptions(options ...RefreshOption)
}

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

type RefreshFuncOptions struct {
	Data any
}

//troptgen:refresh
type RefreshOptions[K comparable, V any] interface {
	OptData(data any)
	OptGetOptions(options ...GetOption)
	OptSetOptions(options ...SetOption)
	OptFunc(refreshFunc CacheRefreshFunc[K, V])
}

//go:generate troptgen

//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Cache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name RefreshCache
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Validator
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name Codec
//go:generate go run github.com/vektra/mockery/v2@v2.23.1 --name CacheRefreshFunc
