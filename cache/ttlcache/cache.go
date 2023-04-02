package trttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	cache   *ttlcache.Cache[K, V]
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache *ttlcache.Cache[K, V],
	options ...RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		cache: cache,
		options: rootOptionsImpl[K, V]{
			defaultDuration: ttlcache.DefaultTTL,
		},
	}
	optErr := trcache.ParseRootOptions(&ret.options, options)
	if optErr != nil && !ret.options.ignoreOptionNotSupported {
		return nil, optErr
	}
	return ret, nil
}

func NewRefresh[K comparable, V any, RD any](cache *ttlcache.Cache[K, V],
	options ...RootOption) (trcache.RefreshCache[K, V, RD], error) {
	c, err := New(cache, options...)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](c, options...)
}

func NewDefault[K comparable, V any](options ...RootOption) (*Cache[K, V], error) {
	return New(ttlcache.New[K, V](), options...)
}

func NewDefaultRefresh[K comparable, V any, RD any](options ...RootOption) (trcache.RefreshCache[K, V, RD], error) {
	c, err := NewDefault[K, V](options...)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](c, options...)
}

func (c *Cache[K, V]) Handle() *ttlcache.Cache[K, V] {
	return c.cache
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K,
	options ...GetOption) (V, error) {
	var optns getOptionsImpl[K, V]
	optErr := trcache.ParseGetOptions(&optns, c.options.callDefaultGetOptions, options)
	if optErr != nil && !optns.ignoreOptionNotSupported {
		var empty V
		return empty, optErr
	}

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
	options ...SetOption) error {
	optns := setOptionsImpl[K, V]{
		duration: c.options.defaultDuration,
	}
	optErr := trcache.ParseSetOptions(&optns, c.options.callDefaultSetOptions, options)
	if optErr != nil && !optns.ignoreOptionNotSupported {
		return optErr
	}

	_ = c.cache.Set(key, value, optns.duration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...DeleteOption) error {
	optns := deleteOptionsImpl[K, V]{}
	optErr := trcache.ParseDeleteOptions(&optns, c.options.callDefaultDeleteOptions, options)
	if optErr != nil && !optns.ignoreOptionNotSupported {
		return optErr
	}

	c.cache.Delete(key)
	return nil
}
