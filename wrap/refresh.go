package wrap

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type wrapRefreshCache[K comparable, V any] struct {
	options wrapRefreshCacheOptions[K, V]
	cache   trcache.Cache[K, V]
}

func NewWrapRefreshCache[K comparable, V any](cache trcache.Cache[K, V],
	options ...trcache.CacheOption[K, V]) trcache.RefreshCache[K, V] {
	ret := &wrapRefreshCache[K, V]{cache: cache}
	_ = trcache.ParseCacheOptions[K, V](&ret.options, options)
	return ret
}

func (c *wrapRefreshCache[K, V]) Name() string {
	return c.cache.Name()
}

func (c *wrapRefreshCache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.CacheGetOption[K, V]) (V, error) {
	return c.cache.Get(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.CacheSetOption[K, V]) error {
	return c.cache.Set(ctx, key, value, options...)
}

func (c *wrapRefreshCache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.CacheDeleteOption[K, V]) error {
	return c.cache.Delete(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.CacheRefreshOption[K, V]) (V, error) {
	var optns wrapRefreshCacheRefreshOptions[K, V]
	_ = trcache.ParseCacheRefreshOptions[K, V](&optns, c.options.fnDefaultRefresh, options)

	ret, err := c.Get(ctx, key)
	if err == nil {
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		var empty V
		return empty, err
	}

	refreshFn := c.options.refreshFunc
	if optns.refreshFn != nil {
		refreshFn = optns.refreshFn
	}

	if refreshFn == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = refreshFn(ctx, key, trcache.CacheRefreshFuncOptions{
		Data: optns.data,
	})
	if err != nil {
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.cacheSetOpt...)
	if err != nil {
		var empty V
		return empty, err
	}

	return ret, nil
}
