package wrap

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type wrapRefreshCache[K comparable, V any] struct {
	options wrapRefreshOptionsImpl[K, V]
	cache   trcache.Cache[K, V]
}

func NewWrapRefreshCache[K comparable, V any](cache trcache.Cache[K, V],
	options ...trcache.RootOption) trcache.RefreshCache[K, V] {
	ret := &wrapRefreshCache[K, V]{cache: cache}
	_ = trcache.ParseRootOptions(&ret.options, options)
	return ret
}

func (c *wrapRefreshCache[K, V]) Name() string {
	return c.cache.Name()
}

func (c *wrapRefreshCache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	return c.cache.Get(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	return c.cache.Set(ctx, key, value, options...)
}

func (c *wrapRefreshCache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	return c.cache.Delete(ctx, key, options...)
}

func (c *wrapRefreshCache[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.RefreshOption) (V, error) {
	var optns wrapRefreshRefreshOptionsImpl[K, V]
	_ = trcache.ParseRefreshOptions(&optns, c.options.callDefaultRefreshOptions, options)

	ret, err := c.Get(ctx, key)
	if err == nil {
		if c.options.metricsMetrics != nil {
			c.options.metricsMetrics.Hit(ctx, c.options.metricsName)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if c.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeDecode)
			} else {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if c.options.metricsMetrics != nil {
		c.options.metricsMetrics.Miss(ctx, c.options.metricsName)
	}

	// call refresh
	refreshFn := c.options.defaultRefreshFunc
	if optns.refreshFunc != nil {
		refreshFn = optns.refreshFunc
	}

	if refreshFn == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = refreshFn(ctx, key, trcache.RefreshFuncOptions{
		Data: optns.data,
	})
	if err != nil {
		if c.options.metricsMetrics != nil {
			c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.setOptions...)
	if err != nil {
		if c.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeEncode)
			} else {
				c.options.metricsMetrics.Error(ctx, c.options.metricsName, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
