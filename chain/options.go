package chain

import (
	"context"
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type Options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(string)
	OptRefreshFunc(trcache.CacheRefreshFunc[K, V])
	OptSetPreviousOnGet(bool)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	fnDefaultGet     []trcache.GetOption
	fnDefaultSet     []trcache.SetOption
	fnDefaultDelete  []trcache.DeleteOption
	name             string
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	setPreviousOnGet bool
}

var _ Options[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptCallDefaultGetOptions(i ...trcache.GetOption) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultSetOptions(i ...trcache.SetOption) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultDeleteOptions(i ...trcache.DeleteOption) {
	c.fnDefaultDelete = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptRefreshFunc(t trcache.CacheRefreshFunc[K, V]) {
	c.refreshFunc = t
}

func (c *cacheOptions[K, V]) OptSetPreviousOnGet(b bool) {
	c.setPreviousOnGet = b
}

func WithName[K comparable, V any](name string) trcache.Option {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.Option {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}

func WithSetPreviousOnGet[K comparable, V any](setPreviousOnGet bool) trcache.Option {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptSetPreviousOnGet(setPreviousOnGet)
			return true
		}
		return false
	})
}

// Cache get options definitions

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

// Cache get options

type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptSetOptions([]trcache.SetOption)
	OptGetStrategy(GetStrategy[K, V])
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []any
	setOptions    []trcache.SetOption
	getStrategy   GetStrategy[K, V]
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (c *getOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *getOptions[K, V]) OptSetOptions(i []trcache.SetOption) {
	c.setOptions = i
}

func (c *getOptions[K, V]) OptGetStrategy(s GetStrategy[K, V]) {
	c.getStrategy = s
}

// Cache get options: declarations

func WithGetSetOptions[K comparable, V any](optns ...trcache.SetOption) trcache.GetOption {
	return trcache.GetOptionFunc(func(options any) bool {
		switch opt := options.(type) {
		case GetOptions[K, V]:
			opt.OptSetOptions(optns)
			return true
		}
		return false
	})
}

func WithGetStrategy[K comparable, V any](s GetStrategy[K, V]) trcache.GetOption {
	return trcache.GetOptionFunc(func(options any) bool {
		switch opt := options.(type) {
		case GetOptions[K, V]:
			opt.OptGetStrategy(s)
			return true
		}
		return false
	})
}

// Cache set options definitions

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

// Cache set options

type SetOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
	OptSetStrategy(SetStrategy[K, V])
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	duration    time.Duration
	setStrategy SetStrategy[K, V]
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (c *setOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

func (c *setOptions[K, V]) OptSetStrategy(s SetStrategy[K, V]) {
	c.setStrategy = s
}

// Cache delete options definitions

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

// Cache delete options

type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
	OptDeleteStrategy(DeleteStrategy[K, V])
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
	deleteStrategy DeleteStrategy[K, V]
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}

func (c *deleteOptions[K, V]) OptDeleteStrategy(s DeleteStrategy[K, V]) {
	c.deleteStrategy = s
}

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
