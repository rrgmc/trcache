package trcache

import (
	"context"
	"time"
)

type Cache[K comparable, V any] interface {
	Get(ctx context.Context, key K) (V, error)
	Set(ctx context.Context, key K, value V, options ...CacheSetOption) error
	Delete(ctx context.Context, key K) error
	Clear(ctx context.Context, key K) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, refreshFn CacheRefreshFunc[V], options ...CacheSetOption) (V, error)
}

type Codec[T any] interface {
	Marshal(ctx context.Context, data T) (any, error)
	Unmarshal(ctx context.Context, data any) (T, error)
}

type CacheRefreshFunc[V any] func(ctx context.Context) (V, error)

type cacheSetOptions struct {
	duration time.Duration
}

type CacheSetOption func(options *cacheSetOptions)

func WithCacheSetDuration(duration time.Duration) CacheSetOption {
	return func(options *cacheSetOptions) {
		options.duration = duration
	}
}
