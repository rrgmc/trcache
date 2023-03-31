package trcache

import "time"

//
// Root options
//

// Root Call Default options

// +troptgen root
type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOptions(...GetOption)
	OptCallDefaultSetOptions(...SetOption)
	OptCallDefaultDeleteOptions(...DeleteOption)
}

// +troptgen root
type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOptions(...RefreshOption)
}

//
// Get options
//

// +troptgen get
type GetOptions[K comparable, V any] interface {
	OptCustomOptions([]any)
}

//
// Set options
//

// +troptgen set
type SetOptions[K comparable, V any] interface {
	OptDuration(time.Duration)
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
	OptData(any)
	OptSetOptions([]SetOption)
	OptRefreshFunc(CacheRefreshFunc[K, V])
}

//go:generate troptgen
