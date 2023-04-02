package refresh

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

func GetOrRefresh[K comparable, V any, RD any](ctx context.Context, c trcache.Cache[K, V], key K,
	defaultOptions defaultRefreshOptions[K, V, RD], options ...trcache.RefreshOption) (V, error) {
	optns := refreshOptionsImpl[K, V, RD]{
		funcx: defaultOptions.defaultRefreshFunc,
	}
	optErr := trcache.ParseRefreshOptions(&optns, defaultOptions.callDefaultRefreshOptions, options)
	if optErr != nil && !optns.ignoreOptionNotSupported {
		var empty V
		return empty, optErr
	}

	ret, err := c.Get(ctx, key, optns.getOptions...)
	if err == nil {
		if defaultOptions.metricsMetrics != nil {
			defaultOptions.metricsMetrics.Hit(ctx, defaultOptions.metricsName)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if defaultOptions.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				defaultOptions.metricsMetrics.Error(ctx, defaultOptions.metricsName, trcache.MetricsErrorTypeDecode)
			} else {
				defaultOptions.metricsMetrics.Error(ctx, defaultOptions.metricsName, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if defaultOptions.metricsMetrics != nil {
		defaultOptions.metricsMetrics.Miss(ctx, defaultOptions.metricsName)
	}

	// call refresh
	if optns.funcx == nil {
		var empty V
		return empty, errors.New("refresh function not set")
	}

	ret, err = optns.funcx(ctx, key, trcache.RefreshFuncOptions[RD]{
		Data: optns.data,
	})
	if err != nil {
		if defaultOptions.metricsMetrics != nil {
			defaultOptions.metricsMetrics.Error(ctx, defaultOptions.metricsName, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, optns.setOptions...)
	if err != nil {
		if defaultOptions.metricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				defaultOptions.metricsMetrics.Error(ctx, defaultOptions.metricsName, trcache.MetricsErrorTypeEncode)
			} else {
				defaultOptions.metricsMetrics.Error(ctx, defaultOptions.metricsName, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
