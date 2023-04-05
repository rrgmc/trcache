package trttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/jellydator/ttlcache/v3"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	cache   *ttlcache.Cache[K, V]
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache *ttlcache.Cache[K, V],
	options ...trcache.RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		cache: cache,
		options: rootOptionsImpl[K, V]{
			defaultDuration: ttlcache.DefaultTTL,
		},
	}
	optErr := trcache.ParseOptions[trcache.RootOption](&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	return ret, nil
}

func NewDefault[K comparable, V any](options ...trcache.RootOption) (*Cache[K, V], error) {
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
	var optns getOptionsImpl[K, V]
	optErr := trcache.ParseOptions[trcache.GetOption](&optns, c.options.callDefaultGetOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
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
	options ...trcache.SetOption) error {
	optns := setOptionsImpl[K, V]{
		duration: c.options.defaultDuration,
	}
	optErr := trcache.ParseOptions[trcache.SetOption](&optns, c.options.callDefaultSetOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	_ = c.cache.Set(key, value, optns.duration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	optns := deleteOptionsImpl[K, V]{}
	optErr := trcache.ParseOptions[trcache.DeleteOption](&optns, c.options.callDefaultDeleteOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	c.cache.Delete(key)
	return nil
}
