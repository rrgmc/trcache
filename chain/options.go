package chain

import (
	"github.com/RangelReale/trcache"
)

// Option

type CacheOptions[K comparable, V any] interface {
	trcache.IsCacheOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptRefreshFunc(trcache.CacheRefreshFunc[K, V])
	OptSetPreviousOnGet(bool)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsCacheOptionImpl
	name             string
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	setPreviousOnGet bool
	fnDefaultGet     []trcache.CacheGetOption[K, V]
	fnDefaultSet     []trcache.CacheSetOption[K, V]
}

var _ CacheOptions[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptFnDefaultGet(i []trcache.CacheGetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultSet(i []trcache.CacheSetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptRefreshFunc(t trcache.CacheRefreshFunc[K, V]) {
	c.refreshFunc = t
}

func (c *cacheOptions[K, V]) OptSetPreviousOnGet(b bool) {
	c.setPreviousOnGet = b
}

// type Option[K comparable, V any] func(*chainOptions[K, V])

// type chainOptions[K comparable, V any] struct {
// 	name             string
// 	refreshFunc      trcache.CacheRefreshFunc[K, V]
// 	setPreviousOnGet bool
// }

func WithName[K comparable, V any](name string) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}

func WithSetPreviousOnGet[K comparable, V any](setPreviousOnGet bool) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptSetPreviousOnGet(setPreviousOnGet)
			return true
		}
		return false
	})
}

// Cache get options

type CacheGetOptions[K comparable, V any] struct {
	trcache.CacheGetOptions[K, V]
	SetPreviousOnGetOptions []trcache.CacheSetOption[K, V]
}

// Cache get options: declarations

func WithCacheGetSetPreviousOnGetOptions[K comparable, V any](optns ...trcache.CacheSetOption[K, V]) trcache.CacheGetOption[K, V] {
	return trcache.CacheGetOptionFunc(func(options any) bool {
		switch opt := options.(type) {
		case *CacheGetOptions[K, V]:
			opt.SetPreviousOnGetOptions = append(opt.SetPreviousOnGetOptions, optns...)
			return true
		}
		return false
	})
}

// // Cache get options: implementations
//
// type withCacheGetSetPreviousOnGetOptions[K comparable, V any] struct {
// 	options []trcache.CacheSetOption[K, V]
// }
//
// func (o withCacheGetSetPreviousOnGetOptions[K, V]) ApplyCacheGetOpt(options any) bool {
// 	switch opt := options.(type) {
// 	case *CacheGetOptions[K, V]:
// 		opt.SetPreviousOnGetOptions = append(opt.SetPreviousOnGetOptions, o.options...)
// 		return true
// 	}
// 	return false
// }
