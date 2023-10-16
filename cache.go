package trcache

import (
	"context"
)

// Cache is the base interface for all cache implementations.
type Cache[K comparable, V any] interface {
	// Name returns the name for the cache. It can be useful in cases where multiple caches are used at the same time,
	// like on [github.com/rrgmc/trcache/cache/chain.Chain]. It often is blank.
	Name() string

	// Get gets the value of an item from the cache with this key. If not found, it returns [ErrNotFound] as the error.
	Get(ctx context.Context, key K, options ...GetOption) (V, error)

	// Set sets the value of the key on the cache.
	Set(ctx context.Context, key K, value V, options ...SetOption) error

	// Delete deletes the key from the cache.
	Delete(ctx context.Context, key K, options ...DeleteOption) error
}

// RefreshCache is the base type for cache implementations that allows refreshing on get, using the
// "GetOrRefresh" method.
type RefreshCache[K comparable, V any] interface {
	Cache[K, V]

	// GetOrRefresh tries to get the value for the key on the cache, calling a refresh function if not found,
	// and then storing the newly loaded value to cache before returning it.
	GetOrRefresh(ctx context.Context, key K, options ...RefreshOption) (V, error)
}

// CacheRefreshFunc is the function signature to use for refreshing values on [RefreshCache].
type CacheRefreshFunc[K comparable, V any] func(ctx context.Context, key K, options RefreshFuncOptions) (V, error)
