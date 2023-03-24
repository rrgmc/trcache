package redis

import (
	"time"

	"github.com/RangelReale/trcache"
)

type Option[K comparable, V any] func(*Cache[K, V])

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.valueCodec = valueCodec
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
