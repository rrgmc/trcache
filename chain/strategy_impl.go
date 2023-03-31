package chain

import (
	"context"

	"github.com/RangelReale/trcache"
)

// Implementations: Get strategy

type GetStrategyGetFirstSetPrevious[K comparable, V any] struct {
}

func (f GetStrategyGetFirstSetPrevious[K, V]) BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult {
	return GetStrategyBeforeResultGet
}

func (f GetStrategyGetFirstSetPrevious[K, V]) AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult {
	if err == nil {
		return GetStrategyAfterResultReturn
	}
	return GetStrategyAfterResultSkip
}

func (f GetStrategyGetFirstSetPrevious[K, V]) BeforeSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult {
	if cacheIdx < gotCacheIdx {
		return GetStrategyBeforeSetResultSet
	}
	return GetStrategyBeforeSetResultSkip
}

func (f GetStrategyGetFirstSetPrevious[K, V]) AfterSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult {
	return GetStrategyAfterSetResultContinue
}

// Implementations: Set Strategy

type SetStrategySetAll[K comparable, V any] struct {
}

func (f SetStrategySetAll[K, V]) BeforeSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult {
	return SetStrategyBeforeResultSet
}

func (f SetStrategySetAll[K, V]) AfterSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult {
	return SetStrategyAfterResultContinue
}

// Implementations: Delete Strategy

type DeleteStrategyDeleteAll[K comparable, V any] struct {
}

func (f DeleteStrategyDeleteAll[K, V]) BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult {
	return DeleteStrategyBeforeResultDelete
}

func (f DeleteStrategyDeleteAll[K, V]) AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult {
	return DeleteStrategyAfterResultContinue
}
