package wrap

import (
	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type WrapRefreshOptions[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultRefreshOptions[K, V]
	OptDefaultRefreshFunc(trcache.CacheRefreshFunc[K, V])
}

type wrapRefreshOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	fnDefaultRefresh []trcache.RefreshOption
}

var _ WrapRefreshOptions[string, string] = &wrapRefreshOptions[string, string]{}

func (w *wrapRefreshOptions[K, V]) OptCallDefaultRefreshOptions(i ...trcache.RefreshOption) {
	w.fnDefaultRefresh = i
}

func (w *wrapRefreshOptions[K, V]) OptDefaultRefreshFunc(t trcache.CacheRefreshFunc[K, V]) {
	w.refreshFunc = t
}

// func WithWrapRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.Option {
// 	return trcache.OptionFunc(func(o any) bool {
// 		switch opt := o.(type) {
// 		case WrapRefreshOptions[K, V]:
// 			opt.OptDefaultRefreshFunc(refreshFunc)
// 			return true
// 		}
// 		return false
// 	})
// }

// func WithWrapRefreshDefaultRefreshOptions[K comparable, V any](options ...trcache.RefreshOption) WrapRefreshOption {
// 	return func(o *wrapRefreshCache[K, V]) {
// 		trcache.WithDefaultRefreshOptions[K, V](options...)(&o.defaultRefreshOptions)
// 	}
// }

// Cache refresh options

// +troptgen refresh
type WrapRefreshRefreshOptions[K comparable, V any] interface {
	trcache.IsRefreshOptions
	trcache.RefreshOptions[K, V]
}

type wrapRefreshRefreshOptions[K comparable, V any] struct {
	trcache.IsRefreshOptionsImpl
	data        any
	cacheSetOpt []trcache.SetOption
	refreshFn   trcache.CacheRefreshFunc[K, V]
}

var _ WrapRefreshRefreshOptions[string, string] = &wrapRefreshRefreshOptions[string, string]{}

func (w *wrapRefreshRefreshOptions[K, V]) OptData(a any) {
	w.data = a
}

func (w *wrapRefreshRefreshOptions[K, V]) OptSetOptions(i []trcache.SetOption) {
	w.cacheSetOpt = w.cacheSetOpt
}

func (w *wrapRefreshRefreshOptions[K, V]) OptRefreshFunc(c trcache.CacheRefreshFunc[K, V]) {
	w.refreshFn = c
}

//go:generate troptgen
