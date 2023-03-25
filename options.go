package trcache

import "time"

// Default options

type DefaultOption[K comparable, V any] func(*DefaultOptions[K, V])

type DefaultOptions[K comparable, V any] struct {
	Get     []CacheGetOption[K, V]
	Set     []CacheSetOption[K, V]
	Refresh []CacheRefreshOption[K, V]
}

// Defaut options: declarations

func WithDefaultGetOptions[K comparable, V any](options ...CacheGetOption[K, V]) DefaultOption[K, V] {
	return func(o *DefaultOptions[K, V]) {
		o.Get = append(o.Get, options...)
	}
}

func WithDefaultSetOptions[K comparable, V any](options ...CacheSetOption[K, V]) DefaultOption[K, V] {
	return func(o *DefaultOptions[K, V]) {
		o.Set = append(o.Set, options...)
	}
}

func WithDefaultRefreshOptions[K comparable, V any](options ...CacheRefreshOption[K, V]) DefaultOption[K, V] {
	return func(o *DefaultOptions[K, V]) {
		o.Refresh = append(o.Refresh, options...)
	}
}

// Cache get options

type CacheGetOption[K comparable, V any] interface {
	ApplyCacheGetOpt(any) bool
}

type CacheGetOptions[K comparable, V any] struct {
	CustomOptions []any
}

// Cache get options: declarations

func WithCacheGetCustomOption[K comparable, V any](cacheGetCustomOption ...any) CacheGetOption[K, V] {
	return withCacheGetCustomOption[K, V]{cacheGetCustomOption}
}

// Cache get options: implementations

type withCacheGetCustomOption[K comparable, V any] struct {
	customOptions []any
}

func (o withCacheGetCustomOption[K, V]) ApplyCacheGetOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheGetOptions[K, V]:
		opt.CustomOptions = append(opt.CustomOptions, o.customOptions...)
		return true
	}
	return false
}

// Cache get options: functions

func ParseCacheGetOptions[K comparable, V any](objs []any, options ...CacheGetOption[K, V]) {
	for _, opt := range options {
		for _, obj := range objs {
			if opt.ApplyCacheGetOpt(obj) {
				break
			}
		}
	}
}

// Cache set options

type CacheSetOption[K comparable, V any] interface {
	ApplyCacheSetOpt(any) bool
}

type CacheSetOptions[K comparable, V any] struct {
	Duration time.Duration
}

// Cache set options: declarations

func WithCacheSetDuration[K comparable, V any](duration time.Duration) CacheSetOption[K, V] {
	return withCacheSetDuration[K, V]{duration}
}

// Cache set options: implementations

type withCacheSetDuration[K comparable, V any] struct {
	duration time.Duration
}

func (o withCacheSetDuration[K, V]) ApplyCacheSetOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheSetOptions[K, V]:
		opt.Duration = o.duration
		return true
	}
	return false
}

// Cache set options: functions

func ParseCacheSetOptions[K comparable, V any](objs []any, options ...CacheSetOption[K, V]) {
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
	CacheSetOpt []CacheSetOption[K, V]
	RefreshFn   CacheRefreshFunc[K, V]
}

// Cache refresh options: declarations

func WithCacheRefreshSetOptions[K comparable, V any](opt ...CacheSetOption[K, V]) CacheRefreshOption[K, V] {
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
	opt []CacheSetOption[K, V]
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
