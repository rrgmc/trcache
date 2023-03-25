package chain

import "github.com/RangelReale/trcache"

type Option[K comparable, V any] func(*chainOptions[K, V])

type chainOptions[K comparable, V any] struct {
	name             string
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	setPreviousOnGet bool
}

func WithName[K comparable, V any](name string) Option[K, V] {
	return func(c *chainOptions[K, V]) {
		c.name = name
	}
}

func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) Option[K, V] {
	return func(c *chainOptions[K, V]) {
		c.refreshFunc = refreshFunc
	}
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
