package ttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

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
