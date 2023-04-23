package chain

import (
	"context"

	"github.com/RangelReale/trcache"
)

func NewDefaultGetStrategy[K comparable, V any]() GetStrategy[K, V] {
	return &GetStrategyGetFirstSetPrevious[K, V]{}
}

func NewDefaultSetStrategy[K comparable, V any]() SetStrategy[K, V] {
	return &SetStrategySetAll[K, V]{}
}

func NewDefaultDeleteStrategy[K comparable, V any]() DeleteStrategy[K, V] {
	return &DeleteStrategyDeleteAll[K, V]{}
}

// Implementations: Get strategy

type GetStrategyGetFirstSetPrevious[K comparable, V any] struct {
}

func (f GetStrategyGetFirstSetPrevious[K, V]) GetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return StrategyLoopOrderFORWARD
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

func (f GetStrategyGetFirstSetPrevious[K, V]) SetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return StrategyLoopOrderBACKWARD
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

func (f SetStrategySetAll[K, V]) SetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return StrategyLoopOrderFORWARD
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

func (f DeleteStrategyDeleteAll[K, V]) DeleteLoopOrder(ctx context.Context) StrategyLoopOrder {
	return StrategyLoopOrderFORWARD
}

func (f DeleteStrategyDeleteAll[K, V]) BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult {
	return DeleteStrategyBeforeResultDELETE
}

func (f DeleteStrategyDeleteAll[K, V]) AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult {
	return DeleteStrategyAfterResultCONTINUE
}

// Implementations: StrategyCallback

type StrategyCallbackFunc struct {
	GetFn    func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterResult)
	GetSetFn func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterSetResult)
	SetFn    func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result SetStrategyAfterResult)
	DeleteFn func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result DeleteStrategyAfterResult)
}

func (s *StrategyCallbackFunc) Get(ctx context.Context, cacheIdx int, cacheName string, key any, err error,
	result GetStrategyAfterResult) {
	if s.GetFn != nil {
		s.GetFn(ctx, cacheIdx, cacheName, key, err, result)
	}
}

func (s *StrategyCallbackFunc) GetSet(ctx context.Context, cacheIdx int, cacheName string, key any, err error,
	result GetStrategyAfterSetResult) {
	if s.GetSetFn != nil {
		s.GetSetFn(ctx, cacheIdx, cacheName, key, err, result)
	}
}

func (s *StrategyCallbackFunc) Set(ctx context.Context, cacheIdx int, cacheName string, key any, err error,
	result SetStrategyAfterResult) {
	if s.SetFn != nil {
		s.SetFn(ctx, cacheIdx, cacheName, key, err, result)
	}
}

func (s *StrategyCallbackFunc) Delete(ctx context.Context, cacheIdx int, cacheName string, key any, err error,
	result DeleteStrategyAfterResult) {
	if s.DeleteFn != nil {
		s.DeleteFn(ctx, cacheIdx, cacheName, key, err, result)
	}
}
