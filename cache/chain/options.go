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
}

// Cache get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
	OptSetOptions(options ...trcache.SetOption)
	OptGetStrategy(getStrategy GetStrategy[K, V])
}

// Cache set options

//troptgen:set
type setOptions[K comparable, V any] interface {
	trcache.SetOptions[K, V]
	OptSetStrategy(setStrategy SetStrategy[K, V])
}

// Cache delete options

//troptgen:delete
type deleteOptions[K comparable, V any] interface {
	trcache.DeleteOptions[K, V]
	OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V])
}

//go:generate troptgen
