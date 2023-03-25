package ttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type Option[K comparable, V any] func(*Cache[K, V])

func WithName[K comparable, V any](name string) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.name = name
	}
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.validator = validator
	}
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.defaultDuration = defaultDuration
	}
}

func WithDefaultGetOptions[K comparable, V any](options ...trcache.CacheGetOption[K, V]) Option[K, V] {
	return func(o *Cache[K, V]) {
		trcache.WithDefaultGetOptions[K, V](options...)(&o.defaultOptions)
	}
}

func WithDefaultSetOptions[K comparable, V any](options ...trcache.CacheSetOption[K, V]) Option[K, V] {
	return func(o *Cache[K, V]) {
		trcache.WithDefaultSetOptions[K, V](options...)(&o.defaultOptions)
	}
}

// Cache get options

type CacheGetOptions[K comparable, V any] struct {
	trcache.CacheGetOptions[K, V]
	Touch bool
}

// Cache get options: declarations

func WithCacheGetTouch[K comparable, V any](touch bool) trcache.CacheGetOption[K, V] {
	return trcache.CacheGetOptionFunc(func(options any) bool {
		switch opt := options.(type) {
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
