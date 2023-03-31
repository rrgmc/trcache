package trcache

import "time"

//
// Root options
//

// Root Call Default options

// +troptgen root
type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOptions(options ...GetOption)
	OptCallDefaultSetOptions(options ...SetOption)
	OptCallDefaultDeleteOptions(options ...DeleteOption)
}

// +troptgen root
type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOptions(options ...RefreshOption)
}

//
// Get options
//

// +troptgen get
type GetOptions[K comparable, V any] interface {
	OptCustomOptions(customOptions []any)
}

//
// Set options
//

// +troptgen set
type SetOptions[K comparable, V any] interface {
	OptDuration(duration time.Duration)
}

// +troptgen delete
type DeleteOptions[K comparable, V any] interface {
}

//
// Refresh options
//

type RefreshFuncOptions struct {
	Data any
}

// +troptgen refresh
type RefreshOptions[K comparable, V any] interface {
	OptData(data any)
	OptSetOptions(options ...SetOption)
	OptRefreshFunc(refreshFunc CacheRefreshFunc[K, V])
}

//go:generate troptgen
