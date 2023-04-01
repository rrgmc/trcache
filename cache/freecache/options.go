package trfreecache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.Options[K, V]
	trcache.CallDefaultOptions[K, V]
	OptKeyCodec(keyCodec trcache.KeyCodec[K])
	OptValueCodec(valueCodec trcache.Codec[V])
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
}

// Cache get options

// +troptgen get
type getOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
}

// Cache set options

// +troptgen set
type setOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
}

// Cache delete options

// +troptgen delete
type deleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
}

//go:generate troptgen
