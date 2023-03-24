package ttlcache

import "github.com/RangelReale/trcache"

type Option[K comparable, V any] func(*Cache[K, V])

func WithValidator[K comparable, V any](validator trcache.Validator[V]) Option[K, V] {
	return func(c *Cache[K, V]) {
		c.validator = validator
	}
}
