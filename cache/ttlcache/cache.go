package trttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	options cacheOptions[K, V]
	cache   *ttlcache.Cache[K, V]
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache *ttlcache.Cache[K, V],
	options ...trcache.RootOption) *Cache[K, V] {
	ret := &Cache[K, V]{
		cache: cache,
		options: cacheOptions[K, V]{
			defaultDuration: ttlcache.DefaultTTL,
		},
	}
	_ = trcache.ParseRootOptions(&ret.options, options)
	return ret
}

func NewDefault[K comparable, V any](options ...trcache.RootOption) *Cache[K, V] {
	return New(ttlcache.New[K, V](), options...)
}

func (c *Cache[K, V]) Handle() *ttlcache.Cache[K, V] {
	return c.cache
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptions[K, V]
	_ = trcache.ParseGetOptions(&optns, c.options.fnDefaultGet, options)

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

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	var optns setOptions[K, V]
	_ = trcache.ParseSetOptions(&optns, c.options.fnDefaultSet, options)

	_ = c.cache.Set(key, value, c.options.defaultDuration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	c.cache.Delete(key)
	return nil
}
