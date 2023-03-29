package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	options cacheOptions[K, V]
	caches  []trcache.Cache[K, V]
}

func New[K comparable, V any](cache []trcache.Cache[K, V],
	options ...trcache.Option[K, V]) *Chain[K, V] {
	ret := &Chain[K, V]{
		caches: cache,
	}
	_ = trcache.ParseOptions[K, V](&ret.options, options)
	return ret
}

func NewRefresh[K comparable, V any](cache []trcache.Cache[K, V],
	options ...trcache.Option[K, V]) trcache.RefreshCache[K, V] {
	// var wopt []wrap.WrapRefreshOption[K, V]
	// if ret.refreshFunc != nil {
	// 	wopt = append(wopt, wrap.WithWrapRefreshFunc[K, V](ret.refreshFunc))
	// }
	return wrap.NewWrapRefreshCache[K, V](New(cache, options...), options...)
}

func (c *Chain[K, V]) Name() string {
	return c.options.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption[K, V]) (V, error) {
	var optns getOptions[K, V]
	_ = trcache.ParseGetOptions(&optns, c.options.fnDefaultGet, options)

	if optns.getStrategy == nil {
		optns.getStrategy = GetStrategyGetFirstSetPrevious[K, V]{}
	}

	var reterr error

	gotCacheIdx := -1
	var ret V

	callOpts := trcache.AppendGetOptions(c.options.fnDefaultGet, options)
	for cacheIdx, cache := range c.caches {
		switch optns.getStrategy.BeforeGet(ctx, cacheIdx, cache, key) {
		case GetStrategyBeforeResultSkip:
			continue
		case GetStrategyBeforeResultGet:
			break
		}

		value, err := cache.Get(ctx, key, callOpts...)

		switch optns.getStrategy.AfterGet(ctx, cacheIdx, cache, key, ret, err) {
		case GetStrategyAfterResultSkip:
			continue
		case GetStrategyAfterResultReturn:
			break
		}

		gotCacheIdx = cacheIdx
		ret = value
		reterr = err
		break
	}

	if reterr != nil {
		var empty V
		return empty, reterr
	}
	if gotCacheIdx == -1 {
		var empty V
		return empty, trcache.ErrNotFound
	}

	callSetOpts := trcache.AppendSetOptions(c.options.fnDefaultSet, optns.setOptions)
	for cacheIdx := len(c.caches) - 1; cacheIdx >= 0; cacheIdx-- {
		switch optns.getStrategy.BeforeSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret) {
		case GetStrategyBeforeSetResultSkip:
			continue
		case GetStrategyBeforeSetResultSet:
			break
		}

		err := c.caches[cacheIdx].Set(ctx, key, ret, callSetOpts...)

		switch optns.getStrategy.AfterSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret, err) {
		case GetStrategyAfterSetResultReturn:
			return ret, err
		case GetStrategyAfterSetResultContinue:
			break
		}
	}

	return ret, reterr
}

func (c *Chain[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption[K, V]) error {
	var optns setOptions[K, V]
	_ = trcache.ParseSetOptions(&optns, c.options.fnDefaultSet, options)

	if optns.setStrategy == nil {
		optns.setStrategy = &SetStrategySetAll[K, V]{}
	}

	var reterr error

	success := false
	callOpts := trcache.AppendSetOptions(c.options.fnDefaultSet, options)
	for cacheIdx, cache := range c.caches {
		switch optns.setStrategy.BeforeSet(ctx, cacheIdx, cache, key, value) {
		case SetStrategyBeforeResultSkip:
			continue
		case SetStrategyBeforeResultSet:
			break
		}

		err := cache.Set(ctx, key, value, callOpts...)

		switch optns.setStrategy.AfterSet(ctx, cacheIdx, cache, key, value, err) {
		case SetStrategyAfterResultReturn:
			return err
		case SetStrategyAfterResultContinueWithError:
			reterr = multierr.Append(reterr, err)
		case SetStrategyAfterResultContinue:
			break
		}

		if err != nil {
			success = true
		}
	}

	if reterr != nil {
		errType := ChainErrorTypeError
		if success {
			// at least one was set
			errType = ChainErrorTypeIncomplete
		}
		return NewChainError(errType, "error setting cache", reterr)
	}
	return nil
}

func (c *Chain[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption[K, V]) error {
	var reterr error

	// delete from all
	success := false
	callOpts := trcache.AppendDeleteOptions(c.options.fnDefaultDelete, options)
	for _, cache := range c.caches {
		if err := cache.Delete(ctx, key, callOpts...); err != nil {
			reterr = multierr.Append(reterr, err)
		} else {
			success = true
		}
	}

	if success || reterr == nil {
		return nil
	}
	return NewChainError(ChainErrorTypeError, "no cache to delete", reterr)
}
