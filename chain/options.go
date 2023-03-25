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

type CacheGetOptions struct {
	trcache.CacheGetOptions
	SetPreviousOnGetOptions []trcache.CacheSetOption
}

// Cache get options: declarations

func WithCacheGetSetPreviousOnGetOptions(options ...trcache.CacheSetOption) trcache.CacheGetOption {
	return &withCacheGetSetPreviousOnGetOptions{options}
}

// Cache get options: implementations

type withCacheGetSetPreviousOnGetOptions struct {
	options []trcache.CacheSetOption
}

func (o withCacheGetSetPreviousOnGetOptions) ApplyCacheGetOpt(options any) bool {
	switch opt := options.(type) {
	case *CacheGetOptions:
		opt.SetPreviousOnGetOptions = append(opt.SetPreviousOnGetOptions, o.options...)
		return true
	}
	return false
}
