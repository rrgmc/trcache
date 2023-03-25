package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	caches           []trcache.Cache[K, V]
	name             string
	setPreviousOnGet bool
}

func New[K comparable, V any](cache []trcache.Cache[K, V], options ...Option[K, V]) trcache.RefreshCache[K, V] {
	var optns chainOptions[K, V]
	for _, opt := range options {
		opt(&optns)
	}
	var wopt []wrap.WrapRefreshOption[K, V]
	if optns.refreshFunc != nil {
		wopt = append(wopt, wrap.WithWrapRefreshFunc[K, V](optns.refreshFunc))
	}

	return wrap.NewWrapRefreshCache[K, V](&Chain[K, V]{
		caches:           cache,
		name:             optns.name,
		setPreviousOnGet: optns.setPreviousOnGet,
	})
}

func (c *Chain[K, V]) Name() string {
	return c.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns CacheGetOptions[K, V]
	trcache.ParseCacheGetOptions([]any{&optns, &optns.CacheGetOptions}, options...)

	var reterr error

	setPrevious := func(cacheIdx int, value V) {
		if c.setPreviousOnGet {
			for p := cacheIdx - 1; p >= 0; p++ {
				err := c.caches[p].Set(ctx, key, value, optns.SetPreviousOnGetOptions...)
				if err != nil {
					// do nothing
				}
			}
		}
	}

	for cacheIdx, cache := range c.caches {
		if ret, err := cache.Get(ctx, key, options...); err == nil {
			setPrevious(cacheIdx, ret)
			return ret, nil
		} else {
			reterr = multierr.Append(reterr, err)
		}
	}

	var empty V
	return empty, trcache.NewChainError("no cache to get", reterr)
}

func (c *Chain[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
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
