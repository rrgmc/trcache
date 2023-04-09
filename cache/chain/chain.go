package chain

import (
	"context"

	"github.com/RangelReale/trcache"
	"go.uber.org/multierr"
)

type Chain[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	caches  []trcache.Cache[K, V]
}

var _ trcache.Cache[string, string] = &Chain[string, string]{}

func New[K comparable, V any](cache []trcache.Cache[K, V],
	options ...trcache.RootOption) (*Chain[K, V], error) {
	ret := &Chain[K, V]{
		caches: cache,
	}
	optErr := trcache.ParseOptions(&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}

	if ret.options.getStrategy == nil {
		ret.options.getStrategy = GetStrategyGetFirstSetPrevious[K, V]{}
	}
	if ret.options.setStrategy == nil {
		ret.options.setStrategy = &SetStrategySetAll[K, V]{}
	}
	if ret.options.deleteStrategy == nil {
		ret.options.deleteStrategy = &DeleteStrategyDeleteAll[K, V]{}
	}

	return ret, nil
}

func (c *Chain[K, V]) Name() string {
	return c.options.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptionsImpl[K, V]

	getChecker := trcache.NewOptionChecker(c.options.callDefaultGetOptions, options)

	optErr := trcache.ParseOptionsChecker(getChecker, &optns)
	if optErr.Err() != nil {
		var empty V
		return empty, optErr.Err()
	}

	var reterr error

	gotCacheIdx := -1
	var ret V

	for cacheIdx, cache := range c.caches {
		switch c.options.getStrategy.BeforeGet(ctx, cacheIdx, cache, key) {
		case GetStrategyBeforeResultSKIP:
			continue
		case GetStrategyBeforeResultGET:
			break
		default:
			var empty V
			return empty, ErrInvalidStrategyResult
		}

		value, err := cache.Get(ctx, key, trcache.ForwardOptionsChecker(getChecker)...)

		switch c.options.getStrategy.AfterGet(ctx, cacheIdx, cache, key, ret, err) {
		case GetStrategyAfterResultSKIP:
			continue
		case GetStrategyAfterResultRETURN:
			break
		default:
			var empty V
			return empty, ErrInvalidStrategyResult
		}

		gotCacheIdx = cacheIdx
		ret = value
		reterr = err
		break
	}

	if err := getChecker.CheckCacheError(); err != nil {
		var empty V
		return empty, err
	}

	if reterr != nil {
		var empty V
		return empty, reterr
	}
	if gotCacheIdx == -1 {
		var empty V
		return empty, trcache.ErrNotFound
	}

	setChecker := trcache.NewOptionChecker(c.options.callDefaultSetOptions, optns.setOptions)

	for cacheIdx := len(c.caches) - 1; cacheIdx >= 0; cacheIdx-- {
		switch c.options.getStrategy.BeforeSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret) {
		case GetStrategyBeforeSetResultSKIP:
			continue
		case GetStrategyBeforeSetResultSET:
			break
		default:
			var empty V
			return empty, ErrInvalidStrategyResult
		}

		err := c.caches[cacheIdx].Set(ctx, key, ret, trcache.ForwardOptionsChecker(setChecker)...)

		switch c.options.getStrategy.AfterSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret, err) {
		case GetStrategyAfterSetResultRETURN:
			return ret, err
		case GetStrategyAfterSetResultCONTINUE:
			break
		default:
			var empty V
			return empty, ErrInvalidStrategyResult
		}
	}

	if err := setChecker.CheckCacheError(); err != nil {
		var empty V
		return empty, err
	}

	return ret, reterr
}

func (c *Chain[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	var optns setOptionsImpl[K, V]

	checker := trcache.NewOptionChecker(c.options.callDefaultSetOptions, options)

	optErr := trcache.ParseOptionsChecker(checker, &optns)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	var reterr error

	success := false
	for cacheIdx, cache := range c.caches {
		switch c.options.setStrategy.BeforeSet(ctx, cacheIdx, cache, key, value) {
		case SetStrategyBeforeResultSKIP:
			continue
		case SetStrategyBeforeResultSET:
			break
		default:
			return ErrInvalidStrategyResult
		}

		err := cache.Set(ctx, key, value, trcache.ForwardOptionsChecker(checker)...)

		switch c.options.setStrategy.AfterSet(ctx, cacheIdx, cache, key, value, err) {
		case SetStrategyAfterResultRETURN:
			return err
		case SetStrategyAfterResultCONTINUEWITHERROR:
			reterr = multierr.Append(reterr, err)
		case SetStrategyAfterResultCONTINUE:
			break
		default:
			return ErrInvalidStrategyResult
		}

		if err != nil {
			success = true
		}
	}

	if err := checker.CheckCacheError(); err != nil {
		return err
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
	options ...trcache.DeleteOption) error {
	var optns deleteOptionsImpl[K, V]
	checker := trcache.NewOptionChecker(c.options.callDefaultDeleteOptions, options)

	optErr := trcache.ParseOptionsChecker(checker, &optns)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	var reterr error

	// delete from all
	success := false
	for cacheIdx, cache := range c.caches {
		switch c.options.deleteStrategy.BeforeDelete(ctx, cacheIdx, cache, key) {
		case DeleteStrategyBeforeResultSKIP:
			continue
		case DeleteStrategyBeforeResultDELETE:
			break
		default:
			return ErrInvalidStrategyResult
		}

		err := cache.Delete(ctx, key, trcache.ForwardOptionsChecker(checker)...)

		switch c.options.deleteStrategy.AfterDelete(ctx, cacheIdx, cache, key, err) {
		case DeleteStrategyAfterResultRETURN:
			return err
		case DeleteStrategyAfterResultCONTINUEWITHERROR:
			reterr = multierr.Append(reterr, err)
		case DeleteStrategyAfterResultCONTINUE:
			break
		default:
			return ErrInvalidStrategyResult
		}

		if err != nil {
			success = true
		}
	}

	if err := checker.CheckCacheError(); err != nil {
		return err
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
