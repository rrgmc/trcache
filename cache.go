package trcache

import (
	"context"
)

type Cache[K comparable, V any] interface {
	Name() string
	Get(ctx context.Context, key K, options ...GetOption[K, V]) (V, error)
	Set(ctx context.Context, key K, value V, options ...SetOption[K, V]) error
	Delete(ctx context.Context, key K, options ...DeleteOption[K, V]) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, options ...RefreshOption[K, V]) (V, error)
}

type CacheRefreshFunc[K comparable, V any] func(ctx context.Context, key K, options RefreshFuncOptions) (V, error)
