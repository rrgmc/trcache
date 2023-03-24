package redis

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

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.keycodec = keyCodec
	}
}

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
