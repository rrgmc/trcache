package trmap

import (
	"github.com/RangelReale/trcache"
)

// Root options

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.NameOptions[K, V]
	trcache.CallDefaultOptions[K, V]
	OptValidator(validator trcache.Validator[V])
}

// Get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
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

//go:generate troptgen
