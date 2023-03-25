package ttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	options cacheOptions[K, V]
	cache   *ttlcache.Cache[K, V]
}

func New[K comparable, V any](cache *ttlcache.Cache[K, V], options ...trcache.CacheOption[K, V]) *Cache[K, V] {
	ret := &Cache[K, V]{
		cache: cache,
		options: cacheOptions[K, V]{
			defaultDuration: ttlcache.DefaultTTL,
		},
	}
	trcache.ParseCacheOptions[K, V](&ret.options, options)
	return ret
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns cacheGetOptions[K, V]
	trcache.ParseCacheGetOptions(&optns, c.options.fnDefaultGet, options)

	var ttlopt []ttlcache.Option[K, V]
	if !optns.touch {
		ttlopt = append(ttlopt, ttlcache.WithDisableTouchOnHit[K, V]())
	}

	item := c.cache.Get(key, ttlopt...)
	if item == nil {
		var empty V
		return empty, trcache.ErrNotFound
	}

	if c.options.validator != nil {
		if err := c.options.validator.ValidateGet(ctx, item.Value()); err != nil {
			var empty V
			return empty, err
		}
	}

	return item.Value(), nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	var optns cacheSetOptions[K, V]
	trcache.ParseCacheSetOptions(&optns, c.options.fnDefaultSet, options)

	_ = c.cache.Set(key, value, c.options.defaultDuration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	c.cache.Delete(key)
	return nil
}
