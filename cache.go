package trcache

import (
	"context"
	"time"
)

type Cache[K comparable, V any] interface {
	Get(ctx context.Context, key K) (V, error)
	Set(ctx context.Context, key K, value V, options ...CacheSetOption) error
	Delete(ctx context.Context, key K) error
	Clear(ctx context.Context) error
}

type RefreshCache[K comparable, V any] interface {
	Cache[K, V]
	GetOrRefresh(ctx context.Context, key K, options ...CacheRefreshOption[K, V]) (V, error)
}

type Codec[V any] interface {
	Marshal(ctx context.Context, data V) (any, error)
	Unmarshal(ctx context.Context, data any) (V, error)
}

type Validator[V any] interface {
	ValidateGet(ctx context.Context, data V) error
}

type CacheRefreshFunc[K comparable, V any] func(ctx context.Context, key K, options CacheRefreshFuncOptions) (V, error)

type cacheSetOptions struct {
	duration time.Duration
}

type CacheSetOption func(options *cacheSetOptions)

func WithCacheSetDuration(duration time.Duration) CacheSetOption {
	return func(options *cacheSetOptions) {
		options.duration = duration
	}
}

type CacheRefreshFuncOptions struct {
	Data any
}

type CacheRefreshOptions[K comparable, V any] struct {
	CacheRefreshFuncOptions
	CacheSetOpt []CacheSetOption
	RefreshFn   CacheRefreshFunc[K, V]
}

type CacheRefreshOption[K comparable, V any] func(options *CacheRefreshOptions[K, V])

func WithCacheRefreshSetOptions[K comparable, V any](opt ...CacheSetOption) CacheRefreshOption[K, V] {
	return func(options *CacheRefreshOptions[K, V]) {
		options.CacheSetOpt = append(options.CacheSetOpt, opt...)
	}
}

func WithCacheRefreshData[K comparable, V any](data any) CacheRefreshOption[K, V] {
	return func(options *CacheRefreshOptions[K, V]) {
		options.Data = data
	}
}

func WithCacheRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) CacheRefreshOption[K, V] {
	return func(options *CacheRefreshOptions[K, V]) {
		options.RefreshFn = fn
	}
}
