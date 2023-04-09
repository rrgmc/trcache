package chain

import (
	"context"

	"github.com/RangelReale/trcache"
)

// Implementations: Get strategy

type GetStrategyGetFirstSetPrevious[K comparable, V any] struct {
}

func (f GetStrategyGetFirstSetPrevious[K, V]) BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult {
	return GetStrategyBeforeResultGET
}

func (f GetStrategyGetFirstSetPrevious[K, V]) AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult {
	if err == nil {
		return GetStrategyAfterResultRETURN
	}
	return GetStrategyAfterResultSKIP
}

func (f GetStrategyGetFirstSetPrevious[K, V]) BeforeSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult {
	if cacheIdx < gotCacheIdx {
		return GetStrategyBeforeSetResultSET
	}
	return GetStrategyBeforeSetResultSKIP
}

func (f GetStrategyGetFirstSetPrevious[K, V]) AfterSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult {
	return GetStrategyAfterSetResultCONTINUE
}

// Implementations: Set Strategy

type SetStrategySetAll[K comparable, V any] struct {
}

func (f SetStrategySetAll[K, V]) BeforeSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult {
	return SetStrategyBeforeResultSET
}

func (f SetStrategySetAll[K, V]) AfterSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult {
	return SetStrategyAfterResultCONTINUE
}

// Implementations: Delete Strategy

type DeleteStrategyDeleteAll[K comparable, V any] struct {
}

func (f DeleteStrategyDeleteAll[K, V]) BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult {
	return DeleteStrategyBeforeResultDELETE
}

func (f DeleteStrategyDeleteAll[K, V]) AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult {
	return DeleteStrategyAfterResultCONTINUE
}
