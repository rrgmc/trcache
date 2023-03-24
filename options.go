package trcache

import "time"

// Cache set options

type CacheSetOption func(options *cacheSetOptions)

type cacheSetOptions struct {
	duration time.Duration
}

func WithCacheSetDuration(duration time.Duration) CacheSetOption {
	return func(options *cacheSetOptions) {
		options.duration = duration
	}
}

// Cache refresh options

type CacheRefreshOption[K comparable, V any] func(options *CacheRefreshOptions[K, V])

type CacheRefreshFuncOptions struct {
	Data any
}

type CacheRefreshOptions[K comparable, V any] struct {
	CacheRefreshFuncOptions
	CacheSetOpt []CacheSetOption
	RefreshFn   CacheRefreshFunc[K, V]
}

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
