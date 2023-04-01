package chain

import (
	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(name string)
	OptRefreshFunc(refreshFunc trcache.CacheRefreshFunc[K, V])
	OptSetPreviousOnGet(setPreviousOnGet bool)
}

// Cache get options

// +troptgen get
type getOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptSetOptions(options ...trcache.SetOption)
	OptGetStrategy(getStrategy GetStrategy[K, V])
}

// Cache set options

// +troptgen set
type setOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
	OptSetStrategy(setStrategy SetStrategy[K, V])
}

// Cache delete options

// +troptgen delete
type deleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
	OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V])
}

//go:generate troptgen
