package trredis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type CacheOptions[K comparable, V any] interface {
	trcache.IsCacheOption
	trcache.CacheFnDefaultOptions[K, V]
	OptName(string)
	OptKeyCodec(trcache.KeyCodec[K])
	OptValueCodec(trcache.Codec[V])
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
	OptGetFunc(GetFunc[K, V])
	OptSetFunc(SetFunc[K, V])
	OptDelFunc(DelFunc[K, V])
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsCacheOptionImpl
	fnDefaultGet    []trcache.CacheGetOption[K, V]
	fnDefaultSet    []trcache.CacheSetOption[K, V]
	name            string
	keyCodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
	getFunc         GetFunc[K, V]
	setFunc         SetFunc[K, V]
	delFunc         DelFunc[K, V]
}

var _ CacheOptions[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptFnDefaultGetOpt(i []trcache.CacheGetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptFnDefaultSetOpt(i []trcache.CacheSetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptKeyCodec(t trcache.KeyCodec[K]) {
	c.keyCodec = t
}

func (c *cacheOptions[K, V]) OptValueCodec(t trcache.Codec[V]) {
	c.valueCodec = t
}

func (c *cacheOptions[K, V]) OptValidator(t trcache.Validator[V]) {
	c.validator = t
}

func (c *cacheOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	c.defaultDuration = duration
}

func (c *cacheOptions[K, V]) OptGetFunc(fn GetFunc[K, V]) {
	c.getFunc = fn
}

func (c *cacheOptions[K, V]) OptSetFunc(fn SetFunc[K, V]) {
	c.setFunc = fn
}

func (c *cacheOptions[K, V]) OptDelFunc(fn DelFunc[K, V]) {
	c.delFunc = fn
}

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

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheOptions[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.CacheOption[K, V] {
	return trcache.CacheOptionFunc(func(o any) bool {
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
	trcache.IsCacheGetOption
	trcache.CacheGetOptions[K, V]
	OptCustomParams(any)
}

type cacheGetOptions[K comparable, V any] struct {
	trcache.IsCacheGetOptionImpl
	customOptions []any
	customParams  any
}

var _ CacheGetOptions[string, string] = &cacheGetOptions[string, string]{}

func (c *cacheGetOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *cacheGetOptions[K, V]) OptCustomParams(params any) {
	c.customParams = params
}

func WithCacheGetCustomParam[K comparable, V any](param any) trcache.CacheGetOption[K, V] {
	return trcache.CacheGetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case CacheGetOptions[K, V]:
			opt.OptCustomParams(param)
			return true
		}
		return false
	})
}

// Cache set options

type CacheSetOptions[K comparable, V any] interface {
	trcache.IsCacheSetOption
	trcache.CacheSetOptions[K, V]
	OptCustomParams(any)
}

type cacheSetOptions[K comparable, V any] struct {
	trcache.IsCacheSetOptionImpl
	duration     time.Duration
	customParams any
}

var _ CacheSetOptions[string, string] = &cacheSetOptions[string, string]{}

func (c *cacheSetOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

func (c *cacheSetOptions[K, V]) OptCustomParams(customParams any) {
	c.customParams = customParams
}
