package chain

import (
	"github.com/rrgmc/trcache"
)

// Option

//troptgen:root
type options[K comparable, V any] interface {
	trcache.Options[K, V]
	trcache.NameOptions[K, V]
	trcache.CallDefaultOptions[K, V]
	// OptDefaultStrategyCallback sets a callback function to receive strategy results.
	OptDefaultStrategyCallback(callback StrategyCallback)
	// OptGetStrategy sets the [GetStrategy] to use for the chain operation. The default is
	// [GetStrategyGetFirstSetPrevious].
	OptGetStrategy(getStrategy GetStrategy[K, V])
	// OptSetStrategy sets the [SetStrategy] to use for the chain operation. The default is
	// [SetStrategySetAll].
	OptSetStrategy(setStrategy SetStrategy[K, V])
	// OptDeleteStrategy sets the [DeleteStrategy] to use for the chain operation. The default is
	// [DeleteStrategyDeleteAll].
	OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V])
}

// Cache get options

//troptgen:get
type getOptions[K comparable, V any] interface {
	trcache.GetOptions[K, V]
	// OptStrategyCallback sets a callback function to receive strategy results.
	OptStrategyCallback(callback StrategyCallback)
	// OptSetOptions adds options to the [Cache.Set] call done after one of the [Cache.Get] function calls succeeds.
	OptSetOptions(options ...trcache.SetOption)
}

// Cache set options

//troptgen:set
type setOptions[K comparable, V any] interface {
	trcache.SetOptions[K, V]
	// OptStrategyCallback sets a callback function to receive strategy results.
	OptStrategyCallback(callback StrategyCallback)
}

// Cache delete options

//troptgen:delete
type deleteOptions[K comparable, V any] interface {
	trcache.DeleteOptions[K, V]
	// OptStrategyCallback sets a callback function to receive strategy results.
	OptStrategyCallback(callback StrategyCallback)
}

//go:generate troptgen
