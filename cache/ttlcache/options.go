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

// func WithDefaultGetOptions[K comparable, V any](options ...trcache.CacheGetOption[K, V]) Option[K, V] {
// 	return func(o *Cache[K, V]) {
// 		trcache.WithDefaultGetOptions[K, V](options...)(&o.defaultOptions)
// 	}
// }
//
// func WithDefaultSetOptions[K comparable, V any](options ...trcache.CacheSetOption[K, V]) Option[K, V] {
// 	return func(o *Cache[K, V]) {
// 		trcache.WithDefaultSetOptions[K, V](options...)(&o.defaultOptions)
// 	}
// }

// Cache get options

type CacheGetOptions[K comparable, V any] struct {
	trcache.CacheGetOptions[K, V]
	Touch bool
}

// Cache get options: declarations

func WithCacheGetTouch[K comparable, V any](touch bool) trcache.CacheGetOption[K, V] {
	return trcache.CacheGetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheGetOptions[K, V]:
			opt.Touch = touch
			return true
		}
		return false
	})
}

// // Cache get options: implementations
//
// type withCacheGetTouch[K comparable, V any] struct {
// 	touch bool
// }
//
// func (o withCacheGetTouch[K, V]) ApplyCacheGetOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheGetOptions[K, V]:
// 		opt.Touch = o.touch
// 		return true
// 	}
// 	return false
// }
