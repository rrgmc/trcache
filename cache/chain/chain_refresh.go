package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/refresh"
)

type ChainRefresh[K comparable, V any] struct {
	*Chain[K, V]
	helper *refresh.Helper[K, V]
}

var _ trcache.RefreshCache[string, string] = &ChainRefresh[string, string]{}

func NewRefresh[K comparable, V any](cache []trcache.Cache[K, V],
	options ...trcache.RootOption) (*ChainRefresh[K, V], error) {
	checker := trcache.NewOptionChecker(options)

	c, err := New[K, V](cache, trcache.ForwardOptionsChecker(checker)...)
	if err != nil {
		return nil, err
	}

	helper, err := refresh.NewHelper[K, V](trcache.ForwardOptionsChecker(checker)...)
	if err != nil {
		return nil, err
	}

	if err = checker.CheckCacheError(); err != nil {
		return nil, err
	}

	return &ChainRefresh[K, V]{
		Chain:  c,
		helper: helper,
	}, nil
}

func (c *ChainRefresh[K, V]) GetOrRefresh(ctx context.Context, key K, options ...trcache.RefreshOption) (V, error) {
	return c.helper.GetOrRefresh(ctx, c, key, options...)
}
