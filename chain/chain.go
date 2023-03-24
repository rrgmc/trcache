package chain

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	caches      []trcache.Cache[K, V]
	refreshFunc trcache.CacheRefreshFunc[K, V]
}

func NewChain[K comparable, V any](cache []trcache.Cache[K, V], options ...ChainOption[K, V]) *Chain[K, V] {
	ret := &Chain[K, V]{
		caches: cache,
	}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

func (c *Chain[K, V]) Get(ctx context.Context, key K) (V, error) {
	var reterr error

	for _, cache := range c.caches {
		if ret, err := cache.Get(ctx, key); err == nil {
			return ret, nil
		} else {
			reterr = multierr.Append(reterr, err)
		}
	}

	var empty V
	return empty, trcache.NewChainError("no cache to get", reterr)
}

func (c *Chain[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption) error {
	var reterr error

	for _, cache := range c.caches {
		if err := cache.Set(ctx, key, value, options...); err == nil {
			return nil
		} else {
			reterr = multierr.Append(reterr, err)
		}
	}

	return trcache.NewChainError("no cache to get", reterr)
}

func (c *Chain[K, V]) Delete(ctx context.Context, key K) error {
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
	return trcache.NewChainError("no cache to get", reterr)
}

func (c *Chain[K, V]) Clear(ctx context.Context) error {
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
	return trcache.NewChainError("no cache to get", reterr)
}

func (c *Chain[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.CacheRefreshOption[K, V]) (V, error) {
	var optns trcache.CacheRefreshOptions[K, V]
	for _, opt := range options {
		opt(&optns)
	}

	ret, err := c.Get(ctx, key)
	if err == nil {
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		var empty V
		return empty, err
	}

	refreshFn := c.refreshFunc
	if optns.RefreshFn != nil {
		refreshFn = optns.RefreshFn
	}

	if refreshFn == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = refreshFn(ctx, key, optns.CacheRefreshFuncOptions)
	if err != nil {
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.CacheSetOpt...)
	if err != nil {
		var empty V
		return empty, err
	}

	return ret, nil
}

type ChainOption[K comparable, V any] func(*Chain[K, V])

func WithChainRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) ChainOption[K, V] {
	return func(c *Chain[K, V]) {
		c.refreshFunc = refreshFunc
	}
}
