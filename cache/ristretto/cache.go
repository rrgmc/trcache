package trristretto

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
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
	optErr := trcache.ParseOptions(&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	if ret.options.valueCodec == nil {
		return nil, errors.New("value codec is required")
	}
	return ret, nil
}

func (c *Cache[K, V]) Handle() *ristretto.Cache {
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

	value, ok := c.cache.Get(key)
	if !ok {
		var empty V
		return empty, trcache.ErrNotFound
	}

	dec, err := c.options.valueCodec.Decode(ctx, value)
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
	optErr := trcache.ParseOptions(&optns, c.options.callDefaultSetOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	enc, err := c.options.valueCodec.Encode(ctx, value)
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
	optns := deleteOptionsImpl[K, V]{}
	optErr := trcache.ParseOptions(&optns, c.options.callDefaultDeleteOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	c.cache.Del(key)
	return nil
}
