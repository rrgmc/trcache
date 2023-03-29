package trcache

import (
	"time"

	"go.uber.org/multierr"
)

//
// Cache options
//

type IsOption interface {
	isCacheOption()
}

type IsOptionImpl struct {
}

func (i IsOptionImpl) isCacheOption() {}

type Option[K comparable, V any] interface {
	ApplyCacheOpt(any) bool
}

type OptionFunc func(any) bool

func (o OptionFunc) ApplyCacheOpt(c any) bool {
	return o(c)
}

// Cache options: builder

type OptionBuilderBase[K comparable, V any] struct {
	opt []Option[K, V]
}

func (ob *OptionBuilderBase[K, V]) AppendOptions(opt ...Option[K, V]) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *OptionBuilderBase[K, V]) ApplyCacheOpt(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheOpt(o); ok {
			found = true
		}
	}
	return found
}

// Cache options: functions

func parseOptions[I any, O any](obj I, apply func(O, I) bool, options ...[]O) error {
	var err error
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if !apply(opt, obj) {
				err = multierr.Append(err, NewOptionNotSupportedError(opt))
			}
		}
	}
	return err
}

func ParseOptions[K comparable, V any](obj IsOption, options ...[]Option[K, V]) error {
	return parseOptions(obj, func(i Option[K, V], o IsOption) bool {
		return i.ApplyCacheOpt(o)
	}, options...)
}

// Cache Fn Default options

type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOpt([]GetOption[K, V])
	OptCallDefaultSetOpt([]SetOption[K, V])
	OptCallDefaultDeleteOpt([]DeleteOption[K, V])
}

type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOpt([]RefreshOption[K, V])
}

// Cache Fn Default options

func WithCallDefaultGetOptions[K comparable, V any](options ...GetOption[K, V]) Option[K, V] {
	return OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultGetOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultSetOptions[K comparable, V any](options ...SetOption[K, V]) Option[K, V] {
	return OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultSetOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultDeleteOptions[K comparable, V any](options ...DeleteOption[K, V]) Option[K, V] {
	return OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultOptions[K, V]:
			opt.OptCallDefaultDeleteOpt(options)
			return true
		}
		return false
	})
}

func WithCallDefaultRefreshOptions[K comparable, V any](options ...RefreshOption[K, V]) Option[K, V] {
	return OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CallDefaultRefreshOptions[K, V]:
			opt.OptCallDefaultRefreshOpt(options)
			return true
		}
		return false
	})
}

// Default options builder

type OptionBuilder[K comparable, V any] struct {
	OptionBuilderBase[K, V]
}

func NewOptionBuilder[K comparable, V any]() *OptionBuilder[K, V] {
	return &OptionBuilder[K, V]{}
}

func (ob *OptionBuilder[K, V]) WithCallDefaultGetOptions(options ...GetOption[K, V]) *OptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultGetOptions[K, V](options...))
	return ob
}

func (ob *OptionBuilder[K, V]) WithCallDefaultSetOptions(options ...SetOption[K, V]) *OptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultSetOptions[K, V](options...))
	return ob
}

func (ob *OptionBuilder[K, V]) WithCallDefaultDeleteOptions(options ...DeleteOption[K, V]) *OptionBuilder[K, V] {
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

type GetOption[K comparable, V any] interface {
	ApplyCacheGetOpt(any) bool
}

type GetOptionFunc func(any) bool

func (o GetOptionFunc) ApplyCacheGetOpt(c any) bool {
	return o(c)
}

// Cache get options: builder

type GetOptionBuilderBase[K comparable, V any] struct {
	opt []GetOption[K, V]
}

func (ob *GetOptionBuilderBase[K, V]) AppendOptions(opt ...GetOption[K, V]) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *GetOptionBuilderBase[K, V]) ApplyCacheGetOpt(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheGetOpt(o); ok {
			found = true
		}
	}
	return found
}

// Cache get options: functions

func ParseGetOptions[K comparable, V any](obj IsGetOption, options ...[]GetOption[K, V]) error {
	return parseOptions(obj, func(i GetOption[K, V], o IsGetOption) bool {
		return i.ApplyCacheGetOpt(o)
	}, options...)
}

// Cache get options: default

type GetOptions[K comparable, V any] interface {
	OptCustomOptions([]any)
}

func WithGetCustomOption[K comparable, V any](options ...any) GetOption[K, V] {
	return GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomOptions(options)
			return true
		}
		return false
	})
}

func AppendGetOptions[K comparable, V any](options ...[]GetOption[K, V]) []GetOption[K, V] {
	var ret []GetOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache set options

type IsSetOption interface {
	isCacheSetOption()
}

type IsSetOptionImpl struct {
}

func (i IsSetOptionImpl) isCacheSetOption() {}

type SetOption[K comparable, V any] interface {
	ApplyCacheSetOpt(any) bool
}

type SetOptionFunc func(any) bool

func (o SetOptionFunc) ApplyCacheSetOpt(c any) bool {
	return o(c)
}

// Cache set options: builder

type SetOptionBuilderBase[K comparable, V any] struct {
	opt []SetOption[K, V]
}

func (ob *SetOptionBuilderBase[K, V]) AppendOptions(opt ...SetOption[K, V]) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *SetOptionBuilderBase[K, V]) ApplyCacheSetOpt(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheSetOpt(o); ok {
			found = true
		}
	}
	return found
}

// Cache set options: functions

func ParseSetOptions[K comparable, V any](obj IsSetOption,
	options ...[]SetOption[K, V]) error {
	return parseOptions(obj, func(i SetOption[K, V], o IsSetOption) bool {
		return i.ApplyCacheSetOpt(o)
	}, options...)
}

func AppendSetOptions[K comparable, V any](options ...[]SetOption[K, V]) []SetOption[K, V] {
	var ret []SetOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache set options: declarations

type SetOptions[K comparable, V any] interface {
	OptDuration(time.Duration)
}

func WithSetDuration[K comparable, V any](duration time.Duration) SetOption[K, V] {
	return SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptDuration(duration)
			return true
		}
		return false
	})
}

// Cache delete options

type IsDeleteOption interface {
	isCacheDeleteOption()
}

type IsDeleteOptionImpl struct {
}

func (i IsDeleteOptionImpl) isCacheDeleteOption() {}

type DeleteOption[K comparable, V any] interface {
	ApplyCacheDeleteOpt(any) bool
}

type DeleteOptionFunc func(any) bool

func (o DeleteOptionFunc) ApplyCacheDeleteOpt(c any) bool {
	return o(c)
}

// Cache delete options: builder

type DeleteOptionBuilderBase[K comparable, V any] struct {
	opt []DeleteOption[K, V]
}

func (ob *DeleteOptionBuilderBase[K, V]) AppendOptions(opt ...DeleteOption[K, V]) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *DeleteOptionBuilderBase[K, V]) ApplyCacheDeleteOpt(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheDeleteOpt(o); ok {
			found = true
		}
	}
	return found
}

// Cache delete options: functions

func ParseDeleteOptions[K comparable, V any](obj IsDeleteOption,
	options ...[]DeleteOption[K, V]) error {
	return parseOptions(obj, func(i DeleteOption[K, V], o IsDeleteOption) bool {
		return i.ApplyCacheDeleteOpt(o)
	}, options...)
}

func AppendDeleteOptions[K comparable, V any](options ...[]DeleteOption[K, V]) []DeleteOption[K, V] {
	var ret []DeleteOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache delete options: declarations

type DeleteOptions[K comparable, V any] interface {
}

// Cache refresh options

type IsRefreshOption interface {
	isCacheRefreshOption()
}

type IsRefreshOptionImpl struct {
}

func (i IsRefreshOptionImpl) isCacheRefreshOption() {}

type RefreshOption[K comparable, V any] interface {
	ApplyCacheRefreshOpt(any) bool
}

type RefreshOptionFunc func(any) bool

func (o RefreshOptionFunc) ApplyCacheRefreshOpt(c any) bool {
	return o(c)
}

// Cache refresh options: builder

type RefreshOptionBuilderBase[K comparable, V any] struct {
	opt []RefreshOption[K, V]
}

func (ob *RefreshOptionBuilderBase[K, V]) AppendOptions(opt ...RefreshOption[K, V]) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *RefreshOptionBuilderBase[K, V]) ApplyCacheRefreshOpt(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheRefreshOpt(o); ok {
			found = true
		}
	}
	return found
}

// Cache refresh options: functions

func ParseRefreshOptions[K comparable, V any](obj IsRefreshOption,
	options ...[]RefreshOption[K, V]) error {
	return parseOptions(obj, func(i RefreshOption[K, V], o IsRefreshOption) bool {
		return i.ApplyCacheRefreshOpt(o)
	}, options...)
}

func AppendRefreshOptions[K comparable, V any](options ...[]RefreshOption[K, V]) []RefreshOption[K, V] {
	var ret []RefreshOption[K, V]
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// Cache refresh options: declarations

type RefreshFuncOptions struct {
	Data any
}

type RefreshOptions[K comparable, V any] interface {
	// RefreshFuncOptions
	OptData(any)
	OptCacheSetOpt([]SetOption[K, V])
	OptRefreshFn(CacheRefreshFunc[K, V])
}

func WithRefreshSetOptions[K comparable, V any](options ...SetOption[K, V]) RefreshOption[K, V] {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptCacheSetOpt(options)
			return true
		}
		return false
	})
}

func WithRefreshData[K comparable, V any](data any) RefreshOption[K, V] {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptData(data)
			return true
		}
		return false
	})
}

func WithRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) RefreshOption[K, V] {
	return RefreshOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case RefreshOptions[K, V]:
			opt.OptRefreshFn(fn)
			return true
		}
		return false
	})
}
