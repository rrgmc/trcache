package refresh

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
)

func GetOrRefresh[K comparable, V any, RD any](ctx context.Context, c trcache.Cache[K, V], key K,
	defaultOptions DefaultRefreshOptions[K, V, RD], options ...trcache.RefreshOption) (V, error) {
	optns := refreshOptionsImpl[K, V, RD]{
		funcx: defaultOptions.DefaultRefreshFunc,
	}
	optErr := trcache.ParseRefreshOptions(&optns, defaultOptions.CallDefaultRefreshOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	ret, err := c.Get(ctx, key, trcache.AppendGetOptions(defaultOptions.CallDefaultGetOptions, optns.getOptions)...)
	if err == nil {
		if defaultOptions.MetricsMetrics != nil {
			defaultOptions.MetricsMetrics.Hit(ctx, defaultOptions.MetricsName)
		}
		return ret, nil
	} else if err != nil && !errors.Is(err, trcache.ErrNotFound) {
		if defaultOptions.MetricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				defaultOptions.MetricsMetrics.Error(ctx, defaultOptions.MetricsName, trcache.MetricsErrorTypeDecode)
			} else {
				defaultOptions.MetricsMetrics.Error(ctx, defaultOptions.MetricsName, trcache.MetricsErrorTypeGet)
			}
		}
		var empty V
		return empty, err
	}

	if defaultOptions.MetricsMetrics != nil {
		defaultOptions.MetricsMetrics.Miss(ctx, defaultOptions.MetricsName)
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
		if defaultOptions.MetricsMetrics != nil {
			defaultOptions.MetricsMetrics.Error(ctx, defaultOptions.MetricsName, trcache.MetricsErrorTypeRefresh)
		}
		var empty V
		return empty, err
	}

	err = c.Set(ctx, key, ret, trcache.AppendSetOptions(defaultOptions.CallDefaultSetOptions, optns.setOptions)...)
	if err != nil {
		if defaultOptions.MetricsMetrics != nil {
			var cerr *trcache.CodecError
			if errors.As(err, &cerr) {
				defaultOptions.MetricsMetrics.Error(ctx, defaultOptions.MetricsName, trcache.MetricsErrorTypeEncode)
			} else {
				defaultOptions.MetricsMetrics.Error(ctx, defaultOptions.MetricsName, trcache.MetricsErrorTypeSet)
			}
		}
		var empty V
		return empty, err
	}

	return ret, nil
}
