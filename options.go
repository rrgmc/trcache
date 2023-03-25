package trcache

import "time"

//
// Cache options
//

type IsCacheOption interface {
	isCacheOption()
}

type IsCacheOptionImpl struct {
}

func (i IsCacheOptionImpl) isCacheOption() {}

type CacheOption[K comparable, V any] interface {
	ApplyCacheOpt(any) bool
}

type CacheOptionFunc func(any) bool

func (o CacheOptionFunc) ApplyCacheOpt(c any) bool {
	return o(c)
}

// Cache options: functions

func ParseCacheOptions[K comparable, V any](obj IsCacheOption, options ...[]CacheOption[K, V]) {
	for _, optinstance := range options {
		for _, opt := range optinstance {
			_ = opt.ApplyCacheOpt(obj)
		}
	}
}

// Cache Fn Default options

type CacheFnDefaultOptions[K comparable, V any] interface {
	OptFnDefaultGet([]CacheGetOption[K, V])
	OptFnDefaultSet([]CacheSetOption[K, V])
}

type CacheFnDefaultRefreshOptions[K comparable, V any] interface {
	OptFnDefaultRefresh([]CacheRefreshOption[K, V])
}

// Cache Fn Default options

func WithCacheFnDefaultGetOptions[K comparable, V any](options ...CacheGetOption[K, V]) CacheOption[K, V] {
	return CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheFnDefaultOptions[K, V]:
			opt.OptFnDefaultGet(options)
			return true
		}
		return false
	})
}

func WithCacheFnDefaultSetOptions[K comparable, V any](options ...CacheSetOption[K, V]) CacheOption[K, V] {
	return CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheFnDefaultOptions[K, V]:
			opt.OptFnDefaultSet(options)
			return true
		}
		return false
	})
}

func WithCacheFnDefaultRefreshOptions[K comparable, V any](options ...CacheRefreshOption[K, V]) CacheOption[K, V] {
	return CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheFnDefaultRefreshOptions[K, V]:
			opt.OptFnDefaultRefresh(options)
			return true
		}
		return false
	})
}

//
// Cache get options
//

type IsCacheGetOption interface {
	isCacheGetOption()
}

type IsCacheGetOptionImpl struct {
}

func (i IsCacheGetOptionImpl) isCacheGetOption() {}

type CacheGetOption[K comparable, V any] interface {
	ApplyCacheGetOpt(any) bool
}

type CacheGetOptionFunc func(any) bool

func (o CacheGetOptionFunc) ApplyCacheGetOpt(c any) bool {
	return o(c)
}

// Cache get options: functions

func ParseCacheGetOptions[K comparable, V any](obj IsCacheGetOption,
	options ...[]CacheGetOption[K, V]) {
	for _, optinstance := range options {
		for _, opt := range optinstance {
			_ = opt.ApplyCacheGetOpt(obj)
		}
	}
}

// Cache get options: default

type CacheGetOptions[K comparable, V any] interface {
	OptCustomOptions([]any)
}

func WithCacheGetCustomOption[K comparable, V any](options ...any) CacheGetOption[K, V] {
	return CacheGetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptCustomOptions(options)
			return true
		}
		return false
	})
}

func AppendCacheGetOptions[K comparable, V any](options ...[]CacheGetOption[K, V]) []CacheGetOption[K, V] {
	var ret []CacheGetOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache set options

type IsCacheSetOption interface {
	isCacheSetOption()
}

type IsCacheSetOptionImpl struct {
}

func (i IsCacheSetOptionImpl) isCacheSetOption() {}

type CacheSetOption[K comparable, V any] interface {
	ApplyCacheSetOpt(any) bool
}

type CacheSetOptionFunc func(any) bool

func (o CacheSetOptionFunc) ApplyCacheSetOpt(c any) bool {
	return o(c)
}

// Cache set options: functions

func ParseCacheSetOptions[K comparable, V any](obj IsCacheSetOption,
	options ...[]CacheSetOption[K, V]) {
	for _, optinstance := range options {
		for _, opt := range optinstance {
			_ = opt.ApplyCacheSetOpt(obj)
		}
	}
}

func AppendCacheSetOptions[K comparable, V any](options ...[]CacheSetOption[K, V]) []CacheSetOption[K, V] {
	var ret []CacheSetOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache set options: declarations

type CacheSetOptions[K comparable, V any] interface {
	OptDuration(time.Duration)
}

func WithCacheSetDuration[K comparable, V any](duration time.Duration) CacheSetOption[K, V] {
	return CacheSetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheSetOptions[K, V]:
			opt.OptDuration(duration)
			return true
		}
		return false
	})
}

// Cache refresh options

type IsCacheRefreshOption interface {
	isCacheRefreshOption()
}

type IsCacheRefreshOptionImpl struct {
}

func (i IsCacheRefreshOptionImpl) isCacheRefreshOption() {}

type CacheRefreshOption[K comparable, V any] interface {
	ApplyCacheRefreshOpt(any) bool
}

type CacheRefreshOptionFunc func(any) bool

func (o CacheRefreshOptionFunc) ApplyCacheRefreshOpt(c any) bool {
	return o(c)
}

// Cache refresh options: functions

func ParseCacheRefreshOptions[K comparable, V any](obj IsCacheRefreshOption,
	options ...[]CacheRefreshOption[K, V]) {
	for _, optinstance := range options {
		for _, opt := range optinstance {
			_ = opt.ApplyCacheRefreshOpt(obj)
		}
	}
}

func AppendCacheRefreshOptions[K comparable, V any](options ...[]CacheRefreshOption[K, V]) []CacheRefreshOption[K, V] {
	var ret []CacheRefreshOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache refresh options: declarations

type CacheRefreshFuncOptions struct {
	Data any
}

type CacheRefreshOptions[K comparable, V any] interface {
	// CacheRefreshFuncOptions
	OptData(any)
	OptCacheSetOpt([]CacheSetOption[K, V])
	OptRefreshFn(CacheRefreshFunc[K, V])
}

func WithCacheRefreshSetOptions[K comparable, V any](options ...CacheSetOption[K, V]) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheRefreshOptions[K, V]:
			opt.OptCacheSetOpt(options)
			return true
		}
		return false
	})
}

func WithCacheRefreshData[K comparable, V any](data any) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheRefreshOptions[K, V]:
			opt.OptData(data)
			return true
		}
		return false
	})
}

func WithCacheRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheRefreshOptions[K, V]:
			opt.OptRefreshFn(fn)
			return true
		}
		return false
	})
}
