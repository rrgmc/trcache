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

// Cache get options

type CacheGetOption func(options *cacheGetOptions)

type cacheGetOptions struct {
	touch bool
}

func WithCacheGetTouch(touch bool) CacheGetOption {
	return func(options *cacheGetOptions) {
		options.touch = touch
	}
}
