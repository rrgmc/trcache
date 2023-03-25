package trcache

import "time"

// Cache options

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

func ParseCacheOptions[K comparable, V any](obj IsCacheOption,
	options ...[]CacheOption[K, V]) {
	for _, optinstance := range options {
		for _, opt := range optinstance {
			_ = opt.ApplyCacheOpt(obj)
		}
	}
}

// Fn Default options

// type DefaultOption[K comparable, V any] func(*CacheFnDefaultOptions[K, V])

// type CacheFnDefaultOptions[K comparable, V any] struct {
// 	FnDefaultGet []CacheGetOption[K, V]
// 	FnDefaultSet []CacheSetOption[K, V]
// }

type CacheFnDefaultOptions[K comparable, V any] interface {
	OptFnDefaultGet([]CacheGetOption[K, V])
	OptFnDefaultSet([]CacheSetOption[K, V])
}

// type DefaultRefreshOption[K comparable, V any] func(*CacheFnDefaultRefreshOptions[K, V])

// type CacheFnDefaultRefreshOptions[K comparable, V any] struct {
// 	FnDefaultRefresh []CacheRefreshOption[K, V]
// }

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

// Cache get options

type CacheGetOption[K comparable, V any] interface {
	ApplyCacheGetOpt(any) bool
}

type CacheGetOptionFunc func(any) bool

func (o CacheGetOptionFunc) ApplyCacheGetOpt(c any) bool {
	return o(c)
}

type CacheGetOptions[K comparable, V any] struct {
	CustomOptions []any
}

// Cache get options: declarations

func WithCacheGetCustomOption[K comparable, V any](options ...any) CacheGetOption[K, V] {
	return CacheGetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheGetOptions[K, V]:
			opt.CustomOptions = append(opt.CustomOptions, options...)
			return true
		}
		return false
	})
}

// // Cache get options: implementations
//
// type withCacheGetCustomOption[K comparable, V any] struct {
// 	customOptions []any
// }
//
// func (o withCacheGetCustomOption[K, V]) ApplyCacheGetOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheGetOptions[K, V]:
// 		opt.CustomOptions = append(opt.CustomOptions, o.customOptions...)
// 		return true
// 	}
// 	return false
// }

// Cache get options: functions

func ParseCacheGetOptions[K comparable, V any](objs []any,
	options ...CacheGetOption[K, V]) {
	for _, opt := range options {
		for _, obj := range objs {
			_ = opt.ApplyCacheGetOpt(obj)
		}
	}
}

// func ParseCacheGetOptions[K comparable, V any](objs []any, defaultOptions []CacheGetOption[K, V],
// 	options ...CacheGetOption[K, V]) {
// 	for _, optitem := range [][]CacheGetOption[K, V]{defaultOptions, options} {
// 		for _, opt := range optitem {
// 			for _, obj := range objs {
// 				if opt.ApplyCacheGetOpt(obj) {
// 					break
// 				}
// 			}
// 		}
// 	}
// }

func AppendCacheGetOptions[K comparable, V any](opt1, opt2 []CacheGetOption[K, V]) []CacheGetOption[K, V] {
	ret := append([]CacheGetOption[K, V]{}, opt1...)
	return append(ret, opt2...)
}

// Cache set options

type CacheSetOption[K comparable, V any] interface {
	ApplyCacheSetOpt(any) bool
}

type CacheSetOptionFunc func(any) bool

func (o CacheSetOptionFunc) ApplyCacheSetOpt(c any) bool {
	return o(c)
}

type CacheSetOptions[K comparable, V any] struct {
	Duration time.Duration
}

// Cache set options: declarations

func WithCacheSetDuration[K comparable, V any](duration time.Duration) CacheSetOption[K, V] {
	return CacheSetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheSetOptions[K, V]:
			opt.Duration = duration
			return true
		}
		return false
	})
}

// // Cache set options: implementations
//
// type withCacheSetDuration[K comparable, V any] struct {
// 	duration time.Duration
// }
//
// func (o withCacheSetDuration[K, V]) ApplyCacheSetOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheSetOptions[K, V]:
// 		opt.Duration = o.duration
// 		return true
// 	}
// 	return false
// }

// Cache set options: functions

func ParseCacheSetOptions[K comparable, V any](objs []any,
	options ...CacheSetOption[K, V]) {
	for _, opt := range options {
		for _, obj := range objs {
			_ = opt.ApplyCacheSetOpt(obj)
		}
	}
}

// func ParseCacheSetOptions[K comparable, V any](objs []any, defaultOptions []CacheSetOption[K, V],
// 	options ...CacheSetOption[K, V]) {
// 	for _, optitem := range [][]CacheSetOption[K, V]{defaultOptions, options} {
// 		for _, opt := range optitem {
// 			for _, obj := range objs {
// 				if opt.ApplyCacheSetOpt(obj) {
// 					break
// 				}
// 			}
// 		}
// 	}
// }

func AppendCacheSetOptions[K comparable, V any](opt1, opt2 []CacheSetOption[K, V]) []CacheSetOption[K, V] {
	ret := append([]CacheSetOption[K, V]{}, opt1...)
	return append(ret, opt2...)
}

// Cache refresh options

type CacheRefreshOption[K comparable, V any] interface {
	ApplyCacheRefreshOpt(any) bool
}

type CacheRefreshOptionFunc func(any) bool

func (o CacheRefreshOptionFunc) ApplyCacheRefreshOpt(c any) bool {
	return o(c)
}

type CacheRefreshFuncOptions struct {
	Data any
}

type CacheRefreshOptions[K comparable, V any] struct {
	CacheRefreshFuncOptions
	CacheSetOpt []CacheSetOption[K, V]
	RefreshFn   CacheRefreshFunc[K, V]
}

// Cache refresh options: declarations

func WithCacheRefreshSetOptions[K comparable, V any](options ...CacheSetOption[K, V]) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheRefreshOptions[K, V]:
			opt.CacheSetOpt = append(opt.CacheSetOpt, options...)
			return true
		}
		return false
	})
}

func WithCacheRefreshData[K comparable, V any](data any) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheRefreshOptions[K, V]:
			opt.Data = data
			return true
		}
		return false
	})
}

func WithCacheRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) CacheRefreshOption[K, V] {
	return CacheRefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case *CacheRefreshOptions[K, V]:
			opt.RefreshFn = fn
			return true
		}
		return false
	})
}

// // Cache refresh options: implementations
//
// type withCacheRefreshSetOptions[K comparable, V any] struct {
// 	opt []CacheSetOption[K, V]
// }
//
// func (o withCacheRefreshSetOptions[K, V]) ApplyCacheRefreshOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheRefreshOptions[K, V]:
// 		opt.CacheSetOpt = append(opt.CacheSetOpt, o.opt...)
// 		return true
// 	}
// 	return false
// }
//
// type withCacheRefreshData[K comparable, V any] struct {
// 	data any
// }
//
// func (o withCacheRefreshData[K, V]) ApplyCacheRefreshOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheRefreshOptions[K, V]:
// 		opt.Data = o.data
// 		return true
// 	}
// 	return false
// }
//
// type withCacheRefreshFunc[K comparable, V any] struct {
// 	fn CacheRefreshFunc[K, V]
// }
//
// func (o withCacheRefreshFunc[K, V]) ApplyCacheRefreshOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheRefreshOptions[K, V]:
// 		opt.RefreshFn = o.fn
// 		return true
// 	}
// 	return false
// }

// Cache refresh options: functions

func ParseCacheRefreshOptions[K comparable, V any](objs []any,
	options ...CacheRefreshOption[K, V]) {
	for _, opt := range options {
		for _, obj := range objs {
			_ = opt.ApplyCacheRefreshOpt(obj)
		}
	}
}

// func ParseCacheRefreshOptions[K comparable, V any](objs []any, defaultOptions []CacheRefreshOption[K, V],
// 	options ...CacheRefreshOption[K, V]) {
// 	for _, optitem := range [][]CacheRefreshOption[K, V]{defaultOptions, options} {
// 		for _, opt := range optitem {
// 			for _, obj := range objs {
// 				if opt.ApplyCacheRefreshOpt(obj) {
// 					break
// 				}
// 			}
// 		}
// 	}
// }

func AppendCacheRefreshOptions[K comparable, V any](opt1, opt2 []CacheRefreshOption[K, V]) []CacheRefreshOption[K, V] {
	ret := append([]CacheRefreshOption[K, V]{}, opt1...)
	return append(ret, opt2...)
}
