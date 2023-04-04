package trttlcache

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/refresh"
	"github.com/jellydator/ttlcache/v3"
)

type RefreshCache[K comparable, V any, RD any] struct {
	*Cache[K, V]
	refreshOptions rootRefreshOptionsImpl[K, V, RD]
	helper         *refresh.Helper[K, V, RD]
}

func NewRefresh[K comparable, V any, RD any](cache *ttlcache.Cache[K, V],
	options ...RootOption) (*RefreshCache[K, V, RD], error) {
	checker := trcache.NewOptionChecker(options)

	c, err := New(cache, trcache.AppendRootOptionsChecker(checker, options)...)
	if err != nil {
		return nil, err
	}

	helper, err := refresh.NewHelper[K, V, RD](trcache.AppendRootOptionsChecker(checker, options)...)
	if err != nil {
		return nil, err
	}

	ret := &RefreshCache[K, V, RD]{
		Cache:          c,
		refreshOptions: rootRefreshOptionsImpl[K, V, RD]{},
		helper:         helper,
	}
	optErr := trcache.ParseRootOptions(&ret.refreshOptions, trcache.AppendRootOptionsChecker(checker, options))
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}

	if err = checker.CheckCacheError(); err != nil {
		return nil, err
	}
	return ret, nil
}

func NewDefaultRefresh[K comparable, V any, RD any](options ...RootOption) (*RefreshCache[K, V, RD], error) {
	return NewRefresh[K, V, RD](ttlcache.New[K, V](), options...)
}

func (c *RefreshCache[K, V, RD]) GetOrRefresh(ctx context.Context, key K, options ...trcache.RefreshOption) (V, error) {
	optns := refreshOptionsImpl[K, V, RD]{
		funcx: c.refreshOptions.defaultRefreshFunc,
	}
	optErr := trcache.ParseRefreshOptions(&optns, c.refreshOptions.callDefaultRefreshOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}
	return c.helper.GetOrRefresh(ctx, c, key, options...)
}
