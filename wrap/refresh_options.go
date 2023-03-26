package wrap

import (
	"github.com/RangelReale/trcache"
)

// Option

type WrapRefreshCacheOptions[K comparable, V any] interface {
	trcache.IsOption
	trcache.CacheFnDefaultRefreshOptions[K, V]
	OptRefreshFunc(trcache.CacheRefreshFunc[K, V])
}

type wrapRefreshCacheOptions[K comparable, V any] struct {
	trcache.IsOptionImpl
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	fnDefaultRefresh []trcache.RefreshOption[K, V]
}

var _ WrapRefreshCacheOptions[string, string] = &wrapRefreshCacheOptions[string, string]{}

func (w *wrapRefreshCacheOptions[K, V]) OptFnDefaultRefreshOpt(i []trcache.RefreshOption[K, V]) {
	w.fnDefaultRefresh = i
}

func (w *wrapRefreshCacheOptions[K, V]) OptRefreshFunc(t trcache.CacheRefreshFunc[K, V]) {
	w.refreshFunc = t
}

func WithWrapRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case WrapRefreshCacheOptions[K, V]:
			opt.OptRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}

// func WithWrapRefreshDefaultRefreshOptions[K comparable, V any](options ...trcache.RefreshOption[K, V]) WrapRefreshOption[K, V] {
// 	return func(o *wrapRefreshCache[K, V]) {
// 		trcache.WithDefaultRefreshOptions[K, V](options...)(&o.defaultRefreshOptions)
// 	}
// }

// Cache set options

type WrapRefreshCacheRefreshOptions[K comparable, V any] interface {
	trcache.IsRefreshOption
	trcache.RefreshOptions[K, V]
}

type wrapRefreshCacheRefreshOptions[K comparable, V any] struct {
	trcache.IsRefreshOptionImpl
	data        any
	cacheSetOpt []trcache.SetOption[K, V]
	refreshFn   trcache.CacheRefreshFunc[K, V]
}

var _ WrapRefreshCacheRefreshOptions[string, string] = &wrapRefreshCacheRefreshOptions[string, string]{}

func (w *wrapRefreshCacheRefreshOptions[K, V]) OptData(a any) {
	w.data = a
}

func (w *wrapRefreshCacheRefreshOptions[K, V]) OptCacheSetOpt(i []trcache.SetOption[K, V]) {
	w.cacheSetOpt = w.cacheSetOpt
}

func (w *wrapRefreshCacheRefreshOptions[K, V]) OptRefreshFn(c trcache.CacheRefreshFunc[K, V]) {
	w.refreshFn = c
}
