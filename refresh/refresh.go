package refresh

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

type Helper[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
}

func NewHelper[K comparable, V any](options ...trcache.RootOption) (*Helper[K, V], error) {
	ret := &Helper[K, V]{}
	optErr := trcache.ParseOptions[trcache.RootOption](&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	return ret, nil
}

func (r *Helper[K, V]) GetOrRefresh(ctx context.Context, c trcache.Cache[K, V], key K,
	options ...trcache.RefreshOption) (V, error) {
	optns := refreshOptionsImpl[K, V]{
		funcx: r.options.defaultRefreshFunc,
	}
	optErr := trcache.ParseOptions[trcache.RefreshOption](&optns, r.options.callDefaultRefreshOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	ret, err := c.Get(ctx, key, optns.getOptions...)
	if err == nil {
		if r.options.metricsMetrics != nil {
			r.options.metricsMetrics.Hit(ctx, r.options.metricsName)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if r.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				r.options.metricsMetrics.Error(ctx, r.options.metricsName, trcache.MetricsErrorTypeDecode)
			} else {
				r.options.metricsMetrics.Error(ctx, r.options.metricsName, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if r.options.metricsMetrics != nil {
		r.options.metricsMetrics.Miss(ctx, r.options.metricsName)
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
		if r.options.metricsMetrics != nil {
			r.options.metricsMetrics.Error(ctx, r.options.metricsName, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.setOptions...)
	if err != nil {
		if r.options.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				r.options.metricsMetrics.Error(ctx, r.options.metricsName, trcache.MetricsErrorTypeEncode)
			} else {
				r.options.metricsMetrics.Error(ctx, r.options.metricsName, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
