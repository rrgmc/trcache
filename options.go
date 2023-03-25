package trcache

import "time"

// Cache get options

type CacheGetOption interface {
	ApplyCacheGetOpt(any) bool
}

type CacheGetOptions struct {
	CustomOptions []any
}

// Cache get options: declarations

func WithCacheGetCustomOption(cacheGetCustomOption ...any) CacheGetOption {
	return withCacheGetCustomOption{cacheGetCustomOption}
}

// Cache get options: implementations

type withCacheGetCustomOption struct {
	customOptions []any
}

func (o withCacheGetCustomOption) ApplyCacheGetOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheGetOptions:
		opt.CustomOptions = append(opt.CustomOptions, o.customOptions...)
		return true
	}
	return false
}

// Cache get options: functions

func ParseCacheGetOptions(objs []any, options ...CacheGetOption) {
	for _, opt := range options {
		for _, obj := range objs {
			if opt.ApplyCacheGetOpt(obj) {
				break
			}
		}
	}
}

// Cache set options

type CacheSetOption interface {
	ApplyCacheSetOpt(any) bool
}

type CacheSetOptions struct {
	Duration time.Duration
}

// Cache set options: declarations

func WithCacheSetDuration(duration time.Duration) CacheSetOption {
	return withCacheSetDuration{duration}
}

// Cache set options: implementations

type withCacheSetDuration struct {
	duration time.Duration
}

func (o withCacheSetDuration) ApplyCacheSetOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheSetOptions:
		opt.Duration = o.duration
		return true
	}
	return false
}

// Cache set options: functions

func ParseCacheSetOptions(objs []any, options ...CacheSetOption) {
	for _, opt := range options {
		for _, obj := range objs {
			if opt.ApplyCacheSetOpt(obj) {
				break
			}
		}
	}
}

// Cache refresh options

type CacheRefreshOption[K comparable, V any] interface {
	ApplyCacheRefreshOpt(any) bool
}

// type CacheRefreshOption[K comparable, V any] func(options *CacheRefreshOptions[K, V])

type CacheRefreshFuncOptions struct {
	Data any
}

type CacheRefreshOptions[K comparable, V any] struct {
	CacheRefreshFuncOptions
	CacheSetOpt []CacheSetOption
	RefreshFn   CacheRefreshFunc[K, V]
}

// Cache refresh options: declarations

func WithCacheRefreshSetOptions[K comparable, V any](opt ...CacheSetOption) CacheRefreshOption[K, V] {
	return withCacheRefreshSetOptions[K, V]{opt}
}

func WithCacheRefreshData[K comparable, V any](data any) CacheRefreshOption[K, V] {
	return withCacheRefreshData[K, V]{data}
}

func WithCacheRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) CacheRefreshOption[K, V] {
	return withCacheRefreshFunc[K, V]{fn}
}

// Cache refresh options: implementations

type withCacheRefreshSetOptions[K comparable, V any] struct {
	opt []CacheSetOption
}

func (o withCacheRefreshSetOptions[K, V]) ApplyCacheRefreshOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheRefreshOptions[K, V]:
		opt.CacheSetOpt = append(opt.CacheSetOpt, o.opt...)
		return true
	}
	return false
}

type withCacheRefreshData[K comparable, V any] struct {
	data any
}

func (o withCacheRefreshData[K, V]) ApplyCacheRefreshOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheRefreshOptions[K, V]:
		opt.Data = o.data
		return true
	}
	return false
}

type withCacheRefreshFunc[K comparable, V any] struct {
	fn CacheRefreshFunc[K, V]
}

func (o withCacheRefreshFunc[K, V]) ApplyCacheRefreshOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheRefreshOptions[K, V]:
		opt.RefreshFn = o.fn
		return true
	}
	return false
}

// Cache refresh options: functions

func ParseCacheRefreshOptions[K comparable, V any](objs []any, options ...CacheRefreshOption[K, V]) {
	for _, opt := range options {
		for _, obj := range objs {
			if opt.ApplyCacheRefreshOpt(obj) {
				break
			}
		}
	}
}
