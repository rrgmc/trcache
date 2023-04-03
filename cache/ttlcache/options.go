package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.NameOptions[K, V]
	trcache.CallDefaultOptions[K, V]
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
}

// Cache get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
	OptTouch(touch bool)
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

//go:generate troptgen
