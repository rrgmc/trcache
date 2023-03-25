package redis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// type CacheOptions[K comparable, V any] struct {
// 	trcache.CacheFnDefaultOptions[K, V]
// 	name            string
// 	keyCodec        trcache.KeyCodec[K]
// 	valueCodec      trcache.Codec[V]
// 	validator       trcache.Validator[V]
// 	defaultDuration time.Duration
// }

type CacheOptions[K comparable, V any] interface {
	trcache.IsCacheOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptKeyCodec(trcache.KeyCodec[K])
	OptValueCodec(trcache.Codec[V])
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsCacheOptionImpl
	name            string
	keyCodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
	fnDefaultGet    []trcache.CacheGetOption[K, V]
	fnDefaultSet    []trcache.CacheSetOption[K, V]
}

var _ CacheOptions[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptFnDefaultGet(i []trcache.CacheGetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultSet(i []trcache.CacheSetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptKeyCodec(t trcache.KeyCodec[K]) {
	c.keyCodec = t
}

func (c *cacheOptions[K, V]) OptValueCodec(t trcache.Codec[V]) {
	c.valueCodec = t
}

func (c *cacheOptions[K, V]) OptValidator(t trcache.Validator[V]) {
	c.validator = t
}

func (c *cacheOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	c.defaultDuration = duration
}

// type Option[K comparable, V any] func(*Cache[K, V])

func WithName[K comparable, V any](name string) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptDefaultDuration(defaultDuration)
			return true
		}
		return false
	})
}
