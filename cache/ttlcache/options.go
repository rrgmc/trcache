package ttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type CacheOptions[K comparable, V any] struct {
	trcache.CacheFnDefaultOptions[K, V]
	name            string
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
