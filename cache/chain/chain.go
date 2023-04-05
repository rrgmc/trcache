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
	optErr := trcache.ParseRootOptions(&ret.options, options)
	if optErr.Err() != nil {
		return nil, optErr.Err()
	}
	return ret, nil
}

func (c *Chain[K, V]) Name() string {
	return c.options.name
}

func (c *Chain[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptionsImpl[K, V]

	callGetOpts := trcache.ConcatOptions(c.options.callDefaultGetOptions, options)
	getChecker := trcache.NewOptionChecker(callGetOpts)

	optErr := trcache.ParseGetOptionsChecker(getChecker, &optns, callGetOpts)
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

	for cacheIdx, cache := range c.caches {
		switch optns.getStrategy.BeforeGet(ctx, cacheIdx, cache, key) {
		case GetStrategyBeforeResultSkip:
			continue
		case GetStrategyBeforeResultGet:
			break
		}

		value, err := cache.Get(ctx, key, trcache.ConcatGetOptionsChecker(getChecker, callGetOpts)...)

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

	callSetOpts := trcache.ConcatOptions(c.options.callDefaultSetOptions, optns.setOptions)
	setChecker := trcache.NewOptionChecker(callSetOpts)

	for cacheIdx := len(c.caches) - 1; cacheIdx >= 0; cacheIdx-- {
		switch optns.getStrategy.BeforeSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret) {
		case GetStrategyBeforeSetResultSkip:
			continue
		case GetStrategyBeforeSetResultSet:
			break
		}

		err := c.caches[cacheIdx].Set(ctx, key, ret, trcache.ConcatSetOptionsChecker(setChecker, callSetOpts)...)

		switch optns.getStrategy.AfterSet(ctx, gotCacheIdx, cacheIdx, c.caches[cacheIdx], key, ret, err) {
		case GetStrategyAfterSetResultReturn:
			return ret, err
		case GetStrategyAfterSetResultContinue:
			break
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

	callOpts := trcache.ConcatOptions(c.options.callDefaultSetOptions, options)
	checker := trcache.NewOptionChecker(callOpts)

	optErr := trcache.ParseSetOptionsChecker(checker, &optns, callOpts)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	if optns.setStrategy == nil {
		optns.setStrategy = &SetStrategySetAll[K, V]{}
	}

	var reterr error

	success := false
	for cacheIdx, cache := range c.caches {
		switch optns.setStrategy.BeforeSet(ctx, cacheIdx, cache, key, value) {
		case SetStrategyBeforeResultSkip:
			continue
		case SetStrategyBeforeResultSet:
			break
		}

		err := cache.Set(ctx, key, value, trcache.ConcatSetOptionsChecker(checker, callOpts)...)

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
	callOpts := trcache.ConcatOptions(c.options.callDefaultDeleteOptions, options)
	checker := trcache.NewOptionChecker(callOpts)

	optErr := trcache.ParseDeleteOptionsChecker(checker, &optns, callOpts)
	if optErr.Err() != nil {
		return optErr.Err()
	}

	if optns.deleteStrategy == nil {
		optns.deleteStrategy = &DeleteStrategyDeleteAll[K, V]{}
	}

	var reterr error

	// delete from all
	success := false
	for cacheIdx, cache := range c.caches {
		switch optns.deleteStrategy.BeforeDelete(ctx, cacheIdx, cache, key) {
		case DeleteStrategyBeforeResultSkip:
			continue
		case DeleteStrategyBeforeResultDelete:
			break
		}

		err := cache.Delete(ctx, key, trcache.ConcatDeleteOptionsChecker(checker, callOpts)...)

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
