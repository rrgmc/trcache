package redis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type CacheOptions[K comparable, V any] struct {
	trcache.CacheFnDefaultOptions[K, V]
	name            string
	keyCodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
}

// type Option[K comparable, V any] func(*Cache[K, V])

func WithName[K comparable, V any](name string) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheOptions[K, V]:
			opt.name = name
			return true
		}
		return false
	})
}

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheOptions[K, V]:
			opt.keyCodec = keyCodec
			return true
		}
		return false
	})
}

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheOptions[K, V]:
			opt.valueCodec = valueCodec
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheOptions[K, V]:
			opt.validator = validator
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheOptions[K, V]:
			opt.defaultDuration = defaultDuration
			return true
		}
		return false
	})
}
