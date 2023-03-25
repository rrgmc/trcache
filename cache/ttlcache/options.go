package ttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// type CacheOptions[K comparable, V any] struct {
// 	trcache.CacheFnDefaultOptions[K, V]
// 	name            string
// 	validator       trcache.Validator[V]
// 	defaultDuration time.Duration
// }

type CacheOptions[K comparable, V any] interface {
	trcache.IsCacheOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsCacheOptionImpl
	name            string
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

// Cache get options

type CacheGetOptions[K comparable, V any] interface {
	trcache.IsCacheGetOption
	trcache.CacheGetOptions[K, V]
	OptTouch(bool)
}

type cacheGetOptions[K comparable, V any] struct {
	trcache.IsCacheGetOptionImpl
	customOptions []any
	touch         bool
}

var _ CacheGetOptions[string, string] = &cacheGetOptions[string, string]{}

func (c *cacheGetOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *cacheGetOptions[K, V]) OptTouch(b bool) {
	c.touch = b
}

// Cache get options: declarations

func WithCacheGetTouch[K comparable, V any](touch bool) trcache.CacheGetOption[K, V] {
	return trcache.CacheGetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptTouch(touch)
			return true
		}
		return false
	})
}

// Cache set options

type CacheSetOptions[K comparable, V any] interface {
	trcache.IsCacheSetOption
	trcache.CacheSetOptions[K, V]
}

type cacheSetOptions[K comparable, V any] struct {
	trcache.IsCacheSetOptionImpl
	duration time.Duration
}

var _ CacheSetOptions[string, string] = &cacheSetOptions[string, string]{}

func (c *cacheSetOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}
