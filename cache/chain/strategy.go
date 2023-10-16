package chain

import (
	"context"

	"github.com/rrgmc/trcache"
)

type StrategyLoopOrder int

const (
	StrategyLoopOrderFORWARD StrategyLoopOrder = iota
	StrategyLoopOrderBACKWARD
)

// Cache get strategy

type GetStrategyBeforeResult int
type GetStrategyAfterResult int
type GetStrategyBeforeSetResult int
type GetStrategyAfterSetResult int

const (
	GetStrategyBeforeResultGET GetStrategyBeforeResult = iota
	GetStrategyBeforeResultSKIP
)

const (
	GetStrategyAfterResultRETURN GetStrategyAfterResult = iota
	GetStrategyAfterResultSKIP
)

const (
	GetStrategyBeforeSetResultSET GetStrategyBeforeSetResult = iota
	GetStrategyBeforeSetResultSKIP
)

const (
	GetStrategyAfterSetResultCONTINUE GetStrategyAfterSetResult = iota
	GetStrategyAfterSetResultRETURN
)

// GetStrategy is the strategy to be used on [Chain.Get].
//
// [Chain.Get] uses this logic:
//   - **loop** on the list of caches from 0 to len().
//   - call [GetStrategy.BeforeGet] on the current cache.
//   - if SKIP, stop processing this cache and go to next loop.
//   - if GET, continue processing.
//   - call [Cache.Get] and store the result / error.
//   - call [GetStrategy.AfterGet] with the result of the previous call (including the error).
//   - if SKIP, continue to the next loop item.
//   - if RETURN, set this result as the current function result (even if it was an error) and exit the loop.
//   - **if no result was found**, or the last returned result is an error, return the error to the user.
//   - **loop** on the list of cache from len() to 0 (backwards)
//   - call [GetStrategy.BeforeSet] on the current cache.
//   - if SKIP, stop processing this cache and go to next loop.
//   - if SET, continue processing.
//   - call [Cache.Set] with the result found in the "Get" loop.
//   - call [GetStrategy.AfterSet] with the result of the previous call (including the error).
//   - if RETURN, return immediately with the result and error of the last "Set" call.
//   - if CONTINUE, continue to the next loop item.
//   - at the end, return the found result and error.
type GetStrategy[K comparable, V any] interface {
	GetLoopOrder(ctx context.Context) StrategyLoopOrder
	BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult
	AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult
	SetLoopOrder(ctx context.Context) StrategyLoopOrder
	BeforeSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult
	AfterSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult
}

type GetStrategyFunc[K comparable, V any] struct {
	GetLoopOrderFn func(ctx context.Context) StrategyLoopOrder
	BeforeGetFn    func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult
	AfterGetFn     func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult
	SetLoopOrderFn func(ctx context.Context) StrategyLoopOrder
	BeforeSetFn    func(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult
	AfterSetFn     func(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult
}

func (f GetStrategyFunc[K, V]) GetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return f.GetLoopOrderFn(ctx)
}

func (f GetStrategyFunc[K, V]) BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult {
	return f.BeforeGetFn(ctx, cacheIdx, cache, key)
}

func (f GetStrategyFunc[K, V]) AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult {
	return f.AfterGetFn(ctx, cacheIdx, cache, key, value, err)
}

func (f GetStrategyFunc[K, V]) SetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return f.SetLoopOrderFn(ctx)
}

func (f GetStrategyFunc[K, V]) BeforeSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult {
	return f.BeforeSetFn(ctx, gotCacheIdx, cacheIdx, cache, key, value)
}

func (f GetStrategyFunc[K, V]) AfterSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult {
	return f.AfterSetFn(ctx, gotCacheIdx, cacheIdx, cache, key, value, err)
}

// Cache set strategy

type SetStrategyBeforeResult int
type SetStrategyAfterResult int

const (
	SetStrategyBeforeResultSET SetStrategyBeforeResult = iota
	SetStrategyBeforeResultSKIP
)

const (
	SetStrategyAfterResultCONTINUE SetStrategyAfterResult = iota
	SetStrategyAfterResultRETURN
	SetStrategyAfterResultCONTINUEWITHERROR
)

// SetStrategy is the strategy to be used on [Chain.Set].
//
// [Chain.Set] uses this logic:
//   - **loop** on the list of caches from 0 to len().
//   - call [SetStrategy.BeforeSet] on the current cache.
//   - if SKIP, stop processing this cache and go to next loop.
//   - if SET, continue processing.
//   - call [Cache.Set] and store the error.
//   - call [SetStrategy.AfterSet] with the result of the previous call (including the error).
//   - if RETURN, return immediately the previous call error result.
//   - if CONTINUEWITHERROR, continue to the next loop item, appending the error to a list of errors to return.
//   - if CONTINUE, continue to the next loop item.
//   - at the end, return the error list if any.
type SetStrategy[K comparable, V any] interface {
	SetLoopOrder(ctx context.Context) StrategyLoopOrder
	BeforeSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult
	AfterSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult
}

type SetStrategyFunc[K comparable, V any] struct {
	SetLoopOrderFn func(ctx context.Context) StrategyLoopOrder
	BeforeSetFn    func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult
	AfterSetFn     func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult
}

func (f SetStrategyFunc[K, V]) SetLoopOrder(ctx context.Context) StrategyLoopOrder {
	return f.SetLoopOrderFn(ctx)
}

func (f SetStrategyFunc[K, V]) BeforeSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult {
	return f.BeforeSetFn(ctx, cacheIdx, cache, key, value)
}

func (f SetStrategyFunc[K, V]) AfterSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult {
	return f.AfterSetFn(ctx, cacheIdx, cache, key, value, err)
}

// Cache delete strategy

type DeleteStrategyBeforeResult int
type DeleteStrategyAfterResult int

const (
	DeleteStrategyBeforeResultDELETE DeleteStrategyBeforeResult = iota
	DeleteStrategyBeforeResultSKIP
)

const (
	DeleteStrategyAfterResultCONTINUE DeleteStrategyAfterResult = iota
	DeleteStrategyAfterResultRETURN
	DeleteStrategyAfterResultCONTINUEWITHERROR
)

// DeleteStrategy is the strategy to be used on [Chain.Delete].
//
// [Chain.Delete] uses this logic:
//   - **loop** on the list of caches from 0 to len().
//   - call [DeleteStrategy.BeforeDelete] on the current cache.
//   - if SKIP, stop processing this cache and go to next loop.
//   - if DELETE, continue processing.
//   - call [Cache.Delete] and store the error
//   - call [DeleteStrategy.AfterDelete] with the result of the previous call (including the error).
//   - if RETURN, return immediately the previous call error result.
//   - if CONTINUEWITHERROR, continue to the next loop item, appending the error to a list of errors to return.
//   - if CONTINUE, continue to the next loop item.
//   - at the end, return the error list if any.
type DeleteStrategy[K comparable, V any] interface {
	DeleteLoopOrder(ctx context.Context) StrategyLoopOrder
	BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult
	AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult
}

type DeleteStrategyFunc[K comparable, V any] struct {
	DeleteLoopOrderFn func(ctx context.Context) StrategyLoopOrder
	BeforeDeleteFn    func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult
	AfterDeleteFn     func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult
}

func (f DeleteStrategyFunc[K, V]) DeleteLoopOrder(ctx context.Context) StrategyLoopOrder {
	return f.DeleteLoopOrderFn(ctx)
}

func (f DeleteStrategyFunc[K, V]) BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult {
	return f.BeforeDeleteFn(ctx, cacheIdx, cache, key)
}

func (f DeleteStrategyFunc[K, V]) AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult {
	return f.AfterDeleteFn(ctx, cacheIdx, cache, key, err)
}

// Callback

type StrategyCallback interface {
	Get(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterResult)
	GetSet(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterSetResult)
	Set(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result SetStrategyAfterResult)
	Delete(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result DeleteStrategyAfterResult)
}
