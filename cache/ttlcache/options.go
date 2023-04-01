package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type Options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(name string)
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
}

// Cache get options

// +troptgen get
type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptTouch(touch bool)
}

// Cache set options

// +troptgen set
type SetOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
}

// Cache delete options

// +troptgen delete
type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
}

//go:generate troptgen
