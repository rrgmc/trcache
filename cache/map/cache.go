package trmap

import (
	"context"

	"github.com/rrgmc/trcache"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	cache   map[K]V
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache map[K]V,
	options ...trcache.RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		cache:   cache,
		options: rootOptionsImpl[K, V]{},
	}
	optErr := trcache.ParseOptions(&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	return ret, nil
}

func NewDefault[K comparable, V any](options ...trcache.RootOption) (*Cache[K, V], error) {
	return New[K, V](map[K]V{}, options...)
}

func (c *Cache[K, V]) Handle() map[K]V {
	return c.cache
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptionsImpl[K, V]
	optErr := trcache.ParseOptions(&optns, c.options.callDefaultGetOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	item, ok := c.cache[key]
	if !ok {
		var empty V
		return empty, trcache.ErrNotFound
	}

	if c.options.validator != nil {
		if err := c.options.validator.ValidateGet(ctx, item); err != nil {
			var empty V
			return empty, err
		}
	}

	return item, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	optns := setOptionsImpl[K, V]{}
	optErr := trcache.ParseOptions(&optns, c.options.callDefaultSetOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	c.cache[key] = value
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	optns := deleteOptionsImpl[K, V]{}
	optErr := trcache.ParseOptions(&optns, c.options.callDefaultDeleteOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	delete(c.cache, key)
	return nil
}
