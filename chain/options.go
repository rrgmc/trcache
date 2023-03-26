package chain

import (
	"github.com/RangelReale/trcache"
)

// Option

type Options[K comparable, V any] interface {
	trcache.IsOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptRefreshFunc(trcache.CacheRefreshFunc[K, V])
	OptSetPreviousOnGet(bool)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsOptionImpl
	fnDefaultGet     []trcache.GetOption[K, V]
	fnDefaultSet     []trcache.SetOption[K, V]
	fnDefaultDelete  []trcache.DeleteOption[K, V]
	name             string
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	setPreviousOnGet bool
}

var _ Options[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptFnDefaultGetOpt(i []trcache.GetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultSetOpt(i []trcache.SetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultDeleteOpt(i []trcache.DeleteOption[K, V]) {
	c.fnDefaultDelete = i
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

func WithName[K comparable, V any](name string) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithRefreshFunc[K comparable, V any](refreshFunc trcache.CacheRefreshFunc[K, V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRefreshFunc(refreshFunc)
			return true
		}
		return false
	})
}

func WithSetPreviousOnGet[K comparable, V any](setPreviousOnGet bool) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptSetPreviousOnGet(setPreviousOnGet)
			return true
		}
		return false
	})
}

// Cache get options

type GetOptions[K comparable, V any] interface {
	trcache.IsGetOption
	trcache.GetOptions[K, V]
	OptSetPreviousOnGetOptions([]trcache.SetOption[K, V])
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionImpl
	customOptions           []any
	setPreviousOnGetOptions []trcache.SetOption[K, V]
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (c *getOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *getOptions[K, V]) OptSetPreviousOnGetOptions(i []trcache.SetOption[K, V]) {
	c.setPreviousOnGetOptions = i
}

// Cache get options: declarations

func WithGetSetPreviousOnGetOptions[K comparable, V any](optns ...trcache.SetOption[K, V]) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(options any) bool {
		switch opt := options.(type) {
		case GetOptions[K, V]:
			opt.OptSetPreviousOnGetOptions(optns)
			return true
		}
		return false
	})
}

// // Cache set options
//
// type SetOptions[K comparable, V any] interface {
// 	trcache.IsSetOption
// 	trcache.SetOptions[K, V]
// }
//
// type cacheSetOptions[K comparable, V any] struct {
// 	trcache.IsSetOptionImpl
// 	duration time.Duration
// }
//
// var _ SetOptions[string, string] = &cacheSetOptions[string, string]{}
//
// func (c *cacheSetOptions[K, V]) OptDuration(duration time.Duration) {
// 	c.duration = duration
// }
