package trcache

import (
	"context"
)

type Cache[K comparable, V any] interface {
	Name() string
	Get(ctx context.Context, key K, options ...CacheGetOption[K, V]) (V, error)
	Set(ctx context.Context, key K, value V, options ...CacheSetOption[K, V]) error
	Delete(ctx context.Context, key K) error
	Clear(ctx context.Context) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, options ...CacheRefreshOption[K, V]) (V, error)
}

type CacheRefreshFunc[K comparable, V any] func(ctx context.Context, key K, options CacheRefreshFuncOptions) (V, error)
