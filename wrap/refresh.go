package wrap

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type wrapRefreshCache[K comparable, V any] struct {
	options wrapRefreshCacheOptions[K, V]
	cache   trcache.Cache[K, V]
	// refreshFunc           trcache.CacheRefreshFunc[K, V]
	// defaultRefreshOptions trcache.CacheFnDefaultRefreshOptions[K, V]
}

func NewWrapRefreshCache[K comparable, V any](cache trcache.Cache[K, V], options ...trcache.CacheOption[K, V]) trcache.RefreshCache[K, V] {
	ret := &wrapRefreshCache[K, V]{cache: cache}
	trcache.ParseCacheOptions[K, V](&ret.options, options)
	return ret
}

func (c *wrapRefreshCache[K, V]) Name() string {
	return c.cache.Name()
}

func (c *wrapRefreshCache[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	return c.cache.Get(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	return c.cache.Set(ctx, key, value, options...)
}

func (c *wrapRefreshCache[K, V]) Delete(ctx context.Context, key K) error {
	return c.cache.Delete(ctx, key)
}

func (c *wrapRefreshCache[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.CacheRefreshOption[K, V]) (V, error) {
	var optns trcache.CacheRefreshOptions[K, V]
	trcache.ParseCacheRefreshOptions[K, V]([]any{&optns},
		trcache.AppendCacheRefreshOptions(c.options.fnDefaultRefresh, options)...)

	ret, err := c.Get(ctx, key)
	if err == nil {
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		var empty V
		return empty, err
	}

	refreshFn := c.options.refreshFunc
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
