package ttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	cache *ttlcache.Cache[K, V]
}

func NewCache[K comparable, V any](cache *ttlcache.Cache[K, V]) *Cache[K, V] {
	return &Cache[K, V]{
		cache: cache,
	}
}

func (c *Cache[K, V]) Get(ctx context.Context, key K) (V, error) {
	item := c.cache.Get(key)
	if item == nil {
		var empty V
		return empty, trcache.ErrNotFound
	}
	return item.Value(), nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption) error {
	_ = c.cache.Set(key, value, ttlcache.DefaultTTL)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	c.cache.Delete(key)
	return nil
}

func (c *Cache[K, V]) Clear(ctx context.Context, key K) error {
	c.cache.DeleteAll()
	return nil
}
