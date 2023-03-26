package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	options cacheOptions[K, V]
	caches  []trcache.Cache[K, V]
}

func New[K comparable, V any](cache []trcache.Cache[K, V], options ...trcache.CacheOption[K, V]) *Chain[K, V] {
	ret := &Chain[K, V]{
		caches: cache,
	}
	_ = trcache.ParseCacheOptions[K, V](&ret.options, options)
	return ret
}

func NewRefresh[K comparable, V any](cache []trcache.Cache[K, V], options ...trcache.CacheOption[K, V]) trcache.RefreshCache[K, V] {
	// var wopt []wrap.WrapRefreshOption[K, V]
	// if ret.refreshFunc != nil {
	// 	wopt = append(wopt, wrap.WithWrapRefreshFunc[K, V](ret.refreshFunc))
	// }
	return wrap.NewWrapRefreshCache[K, V](New(cache, options...), options...)
}

func (c *Chain[K, V]) Name() string {
	return c.options.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns cacheGetOptions[K, V]
	_ = trcache.ParseCacheGetOptions(&optns, c.options.fnDefaultGet, options)

	var reterr error

	setPrevious := func(cacheIdx int, value V) {
		if c.options.setPreviousOnGet {
			for p := cacheIdx - 1; p >= 0; p++ {
				err := c.caches[p].Set(ctx, key, value, optns.setPreviousOnGetOptions...)
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
	return empty, NewChainError(ChainErrorTypeError, "no cache to get", reterr)
}

func (c *Chain[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	var reterr error

	success := false
	for _, cache := range c.caches {
		if err := cache.Set(ctx, key, value, trcache.AppendCacheSetOptions(c.options.fnDefaultSet, options)...); err != nil {
			reterr = multierr.Append(reterr, err)
		} else {
			success = true
		}
	}

	if reterr != nil {
		errType := ChainErrorTypeError
		if success {
			// at least one was set
			errType = ChainErrorTypeIncomplete
		}
		return NewChainError(errType, "error setting cache", reterr)
	}
	return nil
}

func (c *Chain[K, V]) Delete(ctx context.Context, key K) error {
	var reterr error

	// delete from all
	success := false
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
	return NewChainError(ChainErrorTypeError, "no cache to delete", reterr)
}
