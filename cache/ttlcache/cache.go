package ttlcache

import (
	"context"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	cache           *ttlcache.Cache[K, V]
	name            string
	defaultOptions  trcache.DefaultOptions[K, V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
}

func New[K comparable, V any](cache *ttlcache.Cache[K, V], option ...Option[K, V]) *Cache[K, V] {
	ret := &Cache[K, V]{
		cache:           cache,
		defaultDuration: ttlcache.DefaultTTL,
	}
	for _, opt := range option {
		opt(ret)
	}
	return ret
}

func (c *Cache[K, V]) Name() string {
	return c.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns CacheGetOptions[K, V]
	trcache.ParseCacheGetOptions([]any{&optns, &optns.CacheGetOptions},
		trcache.AppendCacheGetOptions(c.defaultOptions.Get, options)...)

	var ttlopt []ttlcache.Option[K, V]
	if !optns.Touch {
		ttlopt = append(ttlopt, ttlcache.WithDisableTouchOnHit[K, V]())
	}

	item := c.cache.Get(key, ttlopt...)
	if item == nil {
		var empty V
		return empty, trcache.ErrNotFound
	}

	if c.validator != nil {
		if err := c.validator.ValidateGet(ctx, item.Value()); err != nil {
			var empty V
			return empty, err
		}
	}

	return item.Value(), nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	var optns trcache.CacheSetOptions[K, V]
	trcache.ParseCacheSetOptions([]any{&optns},
		trcache.AppendCacheSetOptions(c.defaultOptions.Set, options)...)

	_ = c.cache.Set(key, value, c.defaultDuration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	c.cache.Delete(key)
	return nil
}
