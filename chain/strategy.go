package chain

import (
	"context"

	"github.com/RangelReale/trcache"
)

// Cache get strategy

type GetStrategyBeforeResult int
type GetStrategyAfterResult int
type GetStrategyBeforeSetResult int
type GetStrategyAfterSetResult int

const (
	GetStrategyBeforeResultGet GetStrategyBeforeResult = iota
	GetStrategyBeforeResultSkip
)

const (
	GetStrategyAfterResultReturn GetStrategyAfterResult = iota
	GetStrategyAfterResultSkip
)

const (
	GetStrategyBeforeSetResultSet GetStrategyBeforeSetResult = iota
	GetStrategyBeforeSetResultSkip
)

const (
	GetStrategyAfterSetResultContinue GetStrategyAfterSetResult = iota
	GetStrategyAfterSetResultReturn
)

type GetStrategy[K comparable, V any] interface {
	BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult
	AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult
	BeforeSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult
	AfterSet(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult
}

type GetStrategyFunc[K comparable, V any] struct {
	BeforeGetFn func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult
	AfterGetFn  func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult
	BeforeSetFn func(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V) GetStrategyBeforeSetResult
	AfterSetFn  func(ctx context.Context, gotCacheIdx, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterSetResult
}

func (f GetStrategyFunc[K, V]) BeforeGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) GetStrategyBeforeResult {
	return f.BeforeGetFn(ctx, cacheIdx, cache, key)
}

func (f GetStrategyFunc[K, V]) AfterGet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) GetStrategyAfterResult {
	return f.AfterGetFn(ctx, cacheIdx, cache, key, value, err)
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
	SetStrategyBeforeResultSet SetStrategyBeforeResult = iota
	SetStrategyBeforeResultSkip
)

const (
	SetStrategyAfterResultContinue SetStrategyAfterResult = iota
	SetStrategyAfterResultReturn
	SetStrategyAfterResultContinueWithError
)

type SetStrategy[K comparable, V any] interface {
	BeforeSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult
	AfterSet(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult
}

type SetStrategyFunc[K comparable, V any] struct {
	BeforeSetFn func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V) SetStrategyBeforeResult
	AfterSetFn  func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, value V, err error) SetStrategyAfterResult
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
	DeleteStrategyBeforeResultDelete DeleteStrategyBeforeResult = iota
	DeleteStrategyBeforeResultSkip
)

const (
	DeleteStrategyAfterResultContinue DeleteStrategyAfterResult = iota
	DeleteStrategyAfterResultReturn
	DeleteStrategyAfterResultContinueWithError
)

type DeleteStrategy[K comparable, V any] interface {
	BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult
	AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult
}

type DeleteStrategyFunc[K comparable, V any] struct {
	BeforeDeleteFn func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult
	AfterDeleteFn  func(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult
}

func (f DeleteStrategyFunc[K, V]) BeforeDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K) DeleteStrategyBeforeResult {
	return f.BeforeDeleteFn(ctx, cacheIdx, cache, key)
}

func (f DeleteStrategyFunc[K, V]) AfterDelete(ctx context.Context, cacheIdx int, cache trcache.Cache[K, V], key K, err error) DeleteStrategyAfterResult {
	return f.AfterDeleteFn(ctx, cacheIdx, cache, key, err)
}
