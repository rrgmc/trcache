package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/wrap"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	caches  []trcache.Cache[K, V]
}

func New[K comparable, V any](cache []trcache.Cache[K, V],
	options ...RootOption) (*Chain[K, V], error) {
	ret := &Chain[K, V]{
		caches: cache,
	}
	optErr := trcache.ParseRootOptions(&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	return ret, nil
}

func NewRefresh[K comparable, V any, RD any](cache []trcache.Cache[K, V],
	options ...RootOption) (trcache.RefreshCache[K, V, RD], error) {
	c, err := New[K, V](cache,
		trcache.AppendRootOptions(options /* trcache.NewParseRootOptionChecker(), // TODO */)...,
	)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](c, options...)
}

func (c *Chain[K, V]) Name() string {
	return c.options.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K,
	options ...GetOption) (V, error) {
	var optns getOptionsImpl[K, V]
	optErr := trcache.ParseGetOptions(&optns,
		// trcache.NewParseGetOptionChecker(), // TODO
		c.options.callDefaultGetOptions, options)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	if optns.getStrategy == nil {
		optns.getStrategy = GetStrategyGetFirstSetPrevious[K, V]{}
	}

	var reterr error

	gotCacheIdx := -1
	var ret V

	callOpts := trcache.AppendGetOptions(c.options.callDefaultGetOptions, options)
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

	callSetOpts := trcache.AppendSetOptions(c.options.callDefaultSetOptions, optns.setOptions)
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
	options ...SetOption) error {
	var optns setOptionsImpl[K, V]
	optErr := trcache.ParseSetOptions(&optns,
		// trcache.NewParseSetOptionChecker(), // TODO
		c.options.callDefaultSetOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	if optns.setStrategy == nil {
		optns.setStrategy = &SetStrategySetAll[K, V]{}
	}

	var reterr error

	success := false
	callOpts := trcache.AppendSetOptions(c.options.callDefaultSetOptions, options)
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
	options ...DeleteOption) error {
	var optns deleteOptionsImpl[K, V]
	optErr := trcache.ParseDeleteOptions(&optns,
		// trcache.NewParseDeleteOptionChecker(), // TODO
		c.options.callDefaultDeleteOptions, options)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	if optns.deleteStrategy == nil {
		optns.deleteStrategy = &DeleteStrategyDeleteAll[K, V]{}
	}

	var reterr error

	// delete from all
	success := false
	callOpts := trcache.AppendDeleteOptions(c.options.callDefaultDeleteOptions, options)
	for cacheIdx, cache := range c.caches {
		switch optns.deleteStrategy.BeforeDelete(ctx, cacheIdx, cache, key) {
		case DeleteStrategyBeforeResultSkip:
			continue
		case DeleteStrategyBeforeResultDelete:
			break
		}

		err := cache.Delete(ctx, key, callOpts...)

		switch optns.deleteStrategy.AfterDelete(ctx, cacheIdx, cache, key, err) {
		case DeleteStrategyAfterResultReturn:
			return err
		case DeleteStrategyAfterResultContinueWithError:
			reterr = multierr.Append(reterr, err)
		case DeleteStrategyAfterResultContinue:
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
		return NewChainError(errType, "error deleting cache", reterr)
	}
	return nil
}
