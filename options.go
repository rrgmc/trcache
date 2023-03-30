package trcache

import (
	"time"
)

//
// Option
//

type Option interface {
	ApplyCacheOpt(any) bool
}

type OptionFunc func(any) bool

func (o OptionFunc) ApplyCacheOpt(c any) bool {
	return o(c)
}

//
// Root Options
//

type IsRootOption interface {
	isCacheRootOption()
}

type IsRootOptionImpl struct {
}

func (i IsRootOptionImpl) isCacheOption() {}

type IsRootOptions interface {
	isCacheOptions()
}

type IsRootOptionsImpl struct {
}

func (i IsRootOptionsImpl) isCacheOptions() {}

type RootOption interface {
	IsRootOption
	Option
}

type RootOptionFunc func(any) bool

func (f RootOptionFunc) isCacheRootOption() {}

func (f RootOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RootOption = RootOptionFunc(func(a any) bool {
	return true
})

// Cache options: builder base

// type OptionBuilderBase struct {
// 	opt []Option
// }
//
// func (ob *OptionBuilderBase) AppendOptions(opt ...Option) {
// 	ob.opt = append(ob.opt, opt...)
// }
//
// func (ob *OptionBuilderBase) ApplyCacheOpt(o any) bool {
// 	found := false
// 	for _, opt := range ob.opt {
// 		if ok := opt.ApplyCacheOpt(o); ok {
// 			found = true
// 		}
// 	}
// 	return found
// }

// Cache options: functions

func ParseRootOptions(obj IsRootOptions, options ...[]RootOption) error {
	return parseOptions(obj, options...)
}

func AppendRootOptions[K comparable, V any](options ...[]RootOption) []RootOption {
	return appendOptions(options...)
}

// Cache Fn Default options

type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOpt([]GetOption)
	OptCallDefaultSetOpt([]SetOption)
	OptCallDefaultDeleteOpt([]DeleteOption)
}

type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOpt([]RefreshOption)
}

// Cache Fn Default options

func WithCallDefaultGetOptions[K comparable, V any](options ...GetOption) RootOption {
	return RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultGetOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultSetOptions[K comparable, V any](options ...SetOption) RootOption {
	return RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultSetOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultDeleteOptions[K comparable, V any](options ...DeleteOption) RootOption {
	return RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultDeleteOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultRefreshOptions[K comparable, V any](options ...RefreshOption) RootOption {
	return RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultRefreshOptions[K, V]:
			opt.OptCallDefaultRefreshOpt(options)
			return true
		}
		return false
	})
}

// Root options builder

type RootOptionBuilder[K comparable, V any] struct {
	optionBuilder[RootOption]
}

func RootOpt[K comparable, V any]() *RootOptionBuilder[K, V] {
	return &RootOptionBuilder[K, V]{}
}

func (ob *RootOptionBuilder[K, V]) WithCallDefaultGetOptions(options ...GetOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultGetOptions[K, V](options...))
	return ob
}

func (ob *RootOptionBuilder[K, V]) WithCallDefaultSetOptions(options ...SetOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultSetOptions[K, V](options...))
	return ob
}

func (ob *RootOptionBuilder[K, V]) WithCallDefaultDeleteOptions(options ...DeleteOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultDeleteOptions[K, V](options...))
	return ob
}

//
// Cache get options
//

type IsGetOption interface {
	isCacheGetOption()
}

type IsGetOptionImpl struct {
}

func (i IsGetOptionImpl) isCacheGetOption() {}

type IsGetOptions interface {
	isCacheGetOptions()
}

type IsGetOptionsImpl struct {
}

func (i IsGetOptionsImpl) isCacheGetOptions() {}

// type GetOption[K comparable, V any] interface {
// 	ApplyCacheGetOpt(any) bool
// }
//
// type GetOptionFunc func(any) bool
//
// func (o GetOptionFunc) ApplyCacheGetOpt(c any) bool {
// 	return o(c)
// }

type GetOption interface {
	IsGetOption
	Option
}

type GetOptionFunc func(any) bool

func (f GetOptionFunc) isCacheGetOption() {}

func (f GetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ GetOption = GetOptionFunc(func(a any) bool {
	return true
})

// Cache get options: builder base

// type GetOptionBuilderBase[K comparable, V any] struct {
// 	opt []GetOption[K, V]
// }
//
// func (ob *GetOptionBuilderBase[K, V]) AppendOptions(opt ...GetOption[K, V]) {
// 	ob.opt = append(ob.opt, opt...)
// }
//
// func (ob *GetOptionBuilderBase[K, V]) ApplyCacheGetOpt(o any) bool {
// 	found := false
// 	for _, opt := range ob.opt {
// 		if ok := opt.ApplyCacheGetOpt(o); ok {
// 			found = true
// 		}
// 	}
// 	return found
// }

// Cache get options: functions

func ParseGetOptions(obj IsGetOptions, options ...[]GetOption) error {
	return parseOptions(obj, options...)
}

func AppendGetOptions[K comparable, V any](options ...[]GetOption) []GetOption {
	return appendOptions(options...)
}

// Cache get options: default

type GetOptions[K comparable, V any] interface {
	OptCustomOptions([]any)
}

func WithGetCustomOption[K comparable, V any](options ...any) GetOption {
	return GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomOptions(options)
			return true
		}
		return false
	})
}

//
// Cache set options
//

type IsSetOption interface {
	isCacheSetOption()
}

type IsSetOptionImpl struct {
}

func (i IsSetOptionImpl) isCacheSetOption() {}

type IsSetOptions interface {
	isCacheSetOptions()
}

type IsSetOptionsImpl struct {
}

func (i IsSetOptionsImpl) isCacheSetOptions() {}

type SetOption interface {
	IsSetOption
	Option
}

type SetOptionFunc func(any) bool

func (f SetOptionFunc) isCacheSetOption() {}

func (f SetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ SetOption = SetOptionFunc(func(a any) bool {
	return true
})

// Cache set options: functions

func ParseSetOptions(obj IsSetOptions, options ...[]SetOption) error {
	return parseOptions(obj, options...)
}

func AppendSetOptions(options ...[]SetOption) []SetOption {
	return appendOptions(options...)
}

// Cache set options: declarations

type SetOptions[K comparable, V any] interface {
	OptDuration(time.Duration)
}

func WithSetDuration[K comparable, V any](duration time.Duration) SetOption {
	return SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptDuration(duration)
			return true
		}
		return false
	})
}

//
// Cache delete options
//

type IsDeleteOption interface {
	isCacheDeleteOption()
}

type IsDeleteOptionImpl struct {
}

func (i IsDeleteOptionImpl) isCacheDeleteOption() {}

type IsDeleteOptions interface {
	isCacheDeleteOptions()
}

type IsDeleteOptionsImpl struct {
}

func (i IsDeleteOptionsImpl) isCacheDeleteOptions() {}

type DeleteOption interface {
	IsDeleteOption
	Option
}

type DeleteOptionFunc func(any) bool

func (f DeleteOptionFunc) isCacheDeleteOption() {}

func (f DeleteOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ DeleteOption = DeleteOptionFunc(func(a any) bool {
	return true
})

// Cache delete options: functions

func ParseDeleteOptions(obj IsDeleteOptions, options ...[]DeleteOption) error {
	return parseOptions(obj, options...)
}

func AppendDeleteOptions(options ...[]DeleteOption) []DeleteOption {
	return appendOptions(options...)
}

// Cache delete options: declarations

type DeleteOptions[K comparable, V any] interface {
}

//
// Cache refresh options
//

type IsRefreshOption interface {
	isCacheRefreshOption()
}

type IsRefreshOptionImpl struct {
}

func (i IsRefreshOptionImpl) isCacheRefreshOption() {}

type IsRefreshOptions interface {
	isCacheRefreshOptions()
}

type IsRefreshOptionsImpl struct {
}

func (i IsRefreshOptionsImpl) isCacheRefreshOptions() {}

type RefreshOption interface {
	IsRefreshOption
	Option
}

type RefreshOptionFunc func(any) bool

func (f RefreshOptionFunc) isCacheRefreshOption() {}

func (f RefreshOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RefreshOption = RefreshOptionFunc(func(a any) bool {
	return true
})

// Cache refresh options: functions

func ParseRefreshOptions(obj IsRefreshOptions, options ...[]RefreshOption) error {
	return parseOptions(obj, options...)
}

func AppendRefreshOptions(options ...[]RefreshOption) []RefreshOption {
	return appendOptions(options...)
}

// Cache refresh options: declarations

type RefreshFuncOptions struct {
	Data any
}

type RefreshOptions[K comparable, V any] interface {
	// RefreshFuncOptions
	OptData(any)
	OptCacheSetOpt([]SetOption)
	OptRefreshFn(CacheRefreshFunc[K, V])
}

func WithRefreshSetOptions[K comparable, V any](options ...SetOption) RefreshOption {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptCacheSetOpt(options)
			return true
		}
		return false
	})
}

func WithRefreshData[K comparable, V any](data any) RefreshOption {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptData(data)
			return true
		}
		return false
	})
}

func WithRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) RefreshOption {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptRefreshFn(fn)
			return true
		}
		return false
	})
}
