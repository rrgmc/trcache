package trcache

import (
	"context"
)

type Cache[K comparable, V any] interface {
	Name() string
	Get(ctx context.Context, key K, options ...GetOption) (V, error)
	Set(ctx context.Context, key K, value V, options ...SetOption) error
	Delete(ctx context.Context, key K, options ...DeleteOption) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, options ...RefreshOption) (V, error)
}

type CacheRefreshFunc[K comparable, V any] func(ctx context.Context, key K, options RefreshFuncOptions) (V, error)
