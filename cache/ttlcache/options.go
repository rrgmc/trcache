package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// type CacheOptions[K comparable, V any] struct {
// 	trcache.CacheFnDefaultOptions[K, V]
// 	name            string
// 	validator       trcache.Validator[V]
// 	defaultDuration time.Duration
// }

type CacheOptions[K comparable, V any] interface {
	trcache.IsOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsOptionImpl
	fnDefaultGet    []trcache.GetOption[K, V]
	fnDefaultSet    []trcache.SetOption[K, V]
	fnDefaultDelete []trcache.DeleteOption[K, V]
	name            string
	validator       trcache.Validator[V]
	defaultDuration time.Duration
}

var _ CacheOptions[string, string] = &cacheOptions[string, string]{}

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

func (c *cacheOptions[K, V]) OptValidator(t trcache.Validator[V]) {
	c.validator = t
}

func (c *cacheOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	c.defaultDuration = duration
}

// type Option[K comparable, V any] func(*Cache[K, V])

func WithName[K comparable, V any](name string) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptDefaultDuration(defaultDuration)
			return true
		}
		return false
	})
}

// Cache get options

type CacheGetOptions[K comparable, V any] interface {
	trcache.IsGetOption
	trcache.GetOptions[K, V]
	OptTouch(bool)
}

type cacheGetOptions[K comparable, V any] struct {
	trcache.IsGetOptionImpl
	customOptions []any
	touch         bool
}

var _ CacheGetOptions[string, string] = &cacheGetOptions[string, string]{}

func (c *cacheGetOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *cacheGetOptions[K, V]) OptTouch(b bool) {
	c.touch = b
}

// Cache get options: declarations

func WithCacheGetTouch[K comparable, V any](touch bool) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptTouch(touch)
			return true
		}
		return false
	})
}

// Cache set options

type CacheSetOptions[K comparable, V any] interface {
	trcache.IsSetOption
	trcache.SetOptions[K, V]
}

type cacheSetOptions[K comparable, V any] struct {
	trcache.IsSetOptionImpl
	duration time.Duration
}

var _ CacheSetOptions[string, string] = &cacheSetOptions[string, string]{}

func (c *cacheSetOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

// Cache delete options

type CacheDeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOption
	trcache.DeleteOptions[K, V]
}

type cacheDeleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionImpl
}

var _ CacheDeleteOptions[string, string] = &cacheDeleteOptions[string, string]{}
