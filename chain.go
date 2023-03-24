package trcache

import (
	"context"

	"go.uber.org/multierr"
)

type ChainCache[K comparable, V any] struct {
	caches []Cache[K, V]
}

func NewChainCache[K comparable, V any](cache ...Cache[K, V]) *ChainCache[K, V] {
	return &ChainCache[K, V]{
		caches: cache,
	}
}

func (c *ChainCache[K, V]) Get(ctx context.Context, key K) (V, error) {
	var reterr error

	for _, cache := range c.caches {
		if ret, err := cache.Get(ctx, key); err == nil {
			return ret, nil
		} else {
			reterr = multierr.Append(reterr, err)
		}
	}

	var empty V
	return empty, NewChainError("no cache to get", reterr)
}

func (c *ChainCache[K, V]) Set(ctx context.Context, key K, value V, options ...CacheSetOption) error {
	var reterr error

	for _, cache := range c.caches {
		if err := cache.Set(ctx, key, value, options...); err == nil {
			return nil
		} else {
			reterr = multierr.Append(reterr, err)
		}
	}

	return NewChainError("no cache to get", reterr)
}

func (c *ChainCache[K, V]) Delete(ctx context.Context, key K) error {
	var reterr error
	success := false

	// delete from all
	for _, cache := range c.caches {
		if err := cache.Delete(ctx, key); err != nil {
			reterr = multierr.Append(reterr, err)
		} else {
			success = true
		}
	}

	if success || reterr == nil {
		return nil
	}
	return NewChainError("no cache to get", reterr)
}

func (c *ChainCache[K, V]) Clear(ctx context.Context) error {
	var reterr error
	success := false

	// delete from all
	for _, cache := range c.caches {
		if err := cache.Clear(ctx); err != nil {
			reterr = multierr.Append(reterr, err)
		} else {
			success = true
		}
	}

	if success || reterr == nil {
		return nil
	}
	return NewChainError("no cache to get", reterr)
}
