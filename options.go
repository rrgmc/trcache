package trcache

import "time"

//
// Root options
//

// Root Call Default options

// +troptgen root
type CallDefaultOptions[K comparable, V any] interface {
	OptCallDefaultGetOptions(...GetOption)
	OptCallDefaultSetOptions(...SetOption)
	OptCallDefaultDeleteOptions(...DeleteOption)
}

// +troptgen root
type CallDefaultRefreshOptions[K comparable, V any] interface {
	OptCallDefaultRefreshOptions(...RefreshOption)
}

// // Root Call Default options
//
// func WithCallDefaultGetOptions[K comparable, V any](options ...GetOption) RootOption {
// 	return RootOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case CallDefaultOptions[K, V]:
// 			opt.OptCallDefaultGetOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// func WithCallDefaultSetOptions[K comparable, V any](options ...SetOption) RootOption {
// 	return RootOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case CallDefaultOptions[K, V]:
// 			opt.OptCallDefaultSetOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// func WithCallDefaultDeleteOptions[K comparable, V any](options ...DeleteOption) RootOption {
// 	return RootOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case CallDefaultOptions[K, V]:
// 			opt.OptCallDefaultDeleteOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// func WithCallDefaultRefreshOptions[K comparable, V any](options ...RefreshOption) RootOption {
// 	return RootOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case CallDefaultRefreshOptions[K, V]:
// 			opt.OptCallDefaultRefreshOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// // Root options builder
//
// type RootOptionBuilder[K comparable, V any] struct {
// 	RootOptionBuilderBase
// }
//
// func RootOpt[K comparable, V any]() *RootOptionBuilder[K, V] {
// 	return &RootOptionBuilder[K, V]{}
// }
//
// func (ob *RootOptionBuilder[K, V]) WithCallDefaultGetOptions(options ...GetOption) *RootOptionBuilder[K, V] {
// 	ob.AppendOptions(WithCallDefaultGetOptions[K, V](options...))
// 	return ob
// }
//
// func (ob *RootOptionBuilder[K, V]) WithCallDefaultSetOptions(options ...SetOption) *RootOptionBuilder[K, V] {
// 	ob.AppendOptions(WithCallDefaultSetOptions[K, V](options...))
// 	return ob
// }
//
// func (ob *RootOptionBuilder[K, V]) WithCallDefaultDeleteOptions(options ...DeleteOption) *RootOptionBuilder[K, V] {
// 	ob.AppendOptions(WithCallDefaultDeleteOptions[K, V](options...))
// 	return ob
// }

//
// Get options
//

// +troptgen get
type GetOptions[K comparable, V any] interface {
	OptCustomOptions([]any)
}

// func WithGetCustomOptions[K comparable, V any](options ...any) GetOption {
// 	return GetOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case GetOptions[K, V]:
// 			opt.OptCustomOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }

//
// Set options
//

// +troptgen set
type SetOptions[K comparable, V any] interface {
	OptDuration(time.Duration)
}

// func WithSetDuration[K comparable, V any](duration time.Duration) SetOption {
// 	return SetOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case SetOptions[K, V]:
// 			opt.OptDuration(duration)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// //
// // Delete options
// //

// +troptgen delete
type DeleteOptions[K comparable, V any] interface {
}

//
// Refresh options
//

type RefreshFuncOptions struct {
	Data any
}

// +troptgen refresh
type RefreshOptions[K comparable, V any] interface {
	OptData(any)
	OptSetOptions([]SetOption)
	OptFunc(CacheRefreshFunc[K, V])
}

// func WithRefreshSetOptions[K comparable, V any](options ...SetOption) RefreshOption {
// 	return RefreshOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case RefreshOptions[K, V]:
// 			opt.OptSetOptions(options)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// func WithRefreshData[K comparable, V any](data any) RefreshOption {
// 	return RefreshOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case RefreshOptions[K, V]:
// 			opt.OptData(data)
// 			return true
// 		}
// 		return false
// 	})
// }
//
// func WithRefreshFunc[K comparable, V any](fn CacheRefreshFunc[K, V]) RefreshOption {
// 	return RefreshOptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case RefreshOptions[K, V]:
// 			opt.OptFunc(fn)
// 			return true
// 		}
// 		return false
// 	})
// }

//go:generate go run github.com/RangelReale/trcache/optgen
