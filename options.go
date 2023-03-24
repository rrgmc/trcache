package trcache

import "time"

// Cache get options

type CacheGetOption func(options *CacheGetOptions)

type CacheGetOptions struct {
	Touch         bool
	CustomOptions []any
}

func WithCacheGetTouch(touch bool) CacheGetOption {
	return func(options *CacheGetOptions) {
		options.Touch = touch
	}
}

func WithCacheGetCustomOption(cacheGetCustomOption ...any) CacheGetOption {
	return func(options *CacheGetOptions) {
		options.CustomOptions = append(options.CustomOptions, cacheGetCustomOption...)
	}
}

// Cache set options

type CacheSetOption func(options *CacheSetOptions)

type CacheSetOptions struct {
	Duration time.Duration
}

func WithCacheSetDuration(duration time.Duration) CacheSetOption {
	return func(options *CacheSetOptions) {
		options.Duration = duration
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
