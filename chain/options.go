package chain

import "github.com/RangelReale/trcache"

type Option[K comparable, V any] func(*chainOptions[K, V])

type chainOptions[K comparable, V any] struct {
	name        string
	refreshFunc trcache.CacheRefreshFunc[K, V]
}

func WithName[K comparable, V any](name string) Option[K, V] {
	return func(c *chainOptions[K, V]) {
		c.name = name
	}
}

func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) Option[K, V] {
	return func(c *chainOptions[K, V]) {
		c.refreshFunc = refreshFunc
	}
}
