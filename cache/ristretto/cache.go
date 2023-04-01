package trristretto

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"github.com/dgraph-io/ristretto"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	cache   *ristretto.Cache
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache *ristretto.Cache,
	options ...trcache.RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		cache:   cache,
		options: rootOptionsImpl[K, V]{},
	}
	_ = trcache.ParseRootOptions(&ret.options, options)
	if ret.options.valueCodec == nil {
		return nil, errors.New("value codec is required")
	}
	return ret, nil
}

func NewRefresh[K comparable, V any, RD any](cache *ristretto.Cache,
	options ...trcache.RootOption) (trcache.RefreshCache[K, V, RD], error) {
	c, err := New[K, V](cache, options...)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](c, options...), nil
}

// func NewDefault[K comparable, V any](options ...trcache.RootOption) *Cache[K, V] {
// 	return New(cache.New(), options...)
// }

func (c *Cache[K, V]) Handle() *ristretto.Cache {
	return c.cache
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptionsImpl[K, V]
	_ = trcache.ParseGetOptions(&optns, c.options.callDefaultGetOptions, options)

	value, ok := c.cache.Get(key)
	if !ok {
		var empty V
		return empty, trcache.ErrNotFound
	}

	dec, err := c.options.valueCodec.Unmarshal(ctx, value)
	if err != nil {
		var empty V
		return empty, trcache.CodecError{err}
	}

	if c.options.validator != nil {
		if err := c.options.validator.ValidateGet(ctx, dec); err != nil {
			var empty V
			return empty, err
		}
	}

	return dec, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	optns := setOptionsImpl[K, V]{
		duration: c.options.defaultDuration,
	}
	_ = trcache.ParseSetOptions(&optns, c.options.callDefaultSetOptions, options)

	enc, err := c.options.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}

	if !c.cache.SetWithTTL(key, enc, optns.cost, optns.duration) {
		return errors.New("error setting value")
	}
	if !c.options.eventualConsistency {
		// the default for ristretto is eventual consistency, cache may not be sent instantly
		c.cache.Wait()
	}
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	c.cache.Del(key)
	return nil
}
