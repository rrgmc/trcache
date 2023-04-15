package chain

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type ChainRefresh[K comparable, V any] struct {
	*Chain[K, V]
	rootRefreshOptions rootRefreshOptionsImpl[K, V]
}

var _ trcache.RefreshCache[string, string] = &ChainRefresh[string, string]{}

func NewRefresh[K comparable, V any](cache []trcache.Cache[K, V],
	options ...trcache.RootOption) (*ChainRefresh[K, V], error) {
	checker := trcache.NewOptionChecker(options)

	c, err := New[K, V](cache, trcache.ForwardOptionsChecker(checker)...)
	if err != nil {
		return nil, err
	}

	ret := &ChainRefresh[K, V]{
		Chain: c,
	}

	optErr := trcache.ParseOptionsChecker(checker, &ret.rootRefreshOptions)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}

	if err = checker.CheckCacheError(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *ChainRefresh[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.RefreshOption) (V, error) {
	optns := refreshOptionsImpl[K, V]{
		funcx: c.rootRefreshOptions.defaultRefreshFunc,
	}
	optErr := trcache.ParseOptions(&optns, c.rootRefreshOptions.callDefaultRefreshOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	refreshGetInfo := &getInfo{}

	ret, err := c.Get(ctx, key, trcache.ConcatOptions(optns.getOptions, []trcache.GetOption{
		withGetGetInfo[K, V](refreshGetInfo),
	})...)
	if err == nil {
		if c.rootRefreshOptions.metricsMetrics != nil {
			c.rootRefreshOptions.metricsMetrics.Hit(ctx, refreshGetInfo.cacheName,
				c.rootRefreshOptions.metricsName, key)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if c.rootRefreshOptions.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.rootRefreshOptions.metricsMetrics.Error(ctx, refreshGetInfo.cacheName,
					c.rootRefreshOptions.metricsName, key, trcache.MetricsErrorTypeDecode)
			} else {
				c.rootRefreshOptions.metricsMetrics.Error(ctx, refreshGetInfo.cacheName,
					c.rootRefreshOptions.metricsName, key, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if c.rootRefreshOptions.metricsMetrics != nil {
		c.rootRefreshOptions.metricsMetrics.Miss(ctx, c.rootRefreshOptions.metricsName, "", key)
	}

	// call refresh
	if optns.funcx == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = optns.funcx(ctx, key, trcache.RefreshFuncOptions{
		Data: optns.data,
	})
	if err != nil {
		if c.rootRefreshOptions.metricsMetrics != nil {
			c.rootRefreshOptions.metricsMetrics.Error(ctx, c.rootRefreshOptions.metricsName, "",
				key, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.setOptions...)
	if err != nil {
		if c.rootRefreshOptions.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				c.rootRefreshOptions.metricsMetrics.Error(ctx, c.rootRefreshOptions.metricsName, "",
					key, trcache.MetricsErrorTypeEncode)
			} else {
				c.rootRefreshOptions.metricsMetrics.Error(ctx, c.rootRefreshOptions.metricsName, "",
					key, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
