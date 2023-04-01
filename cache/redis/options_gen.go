// Code generated by generator, DO NOT EDIT.
package trredis

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithDefaultDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	})
}
func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}
func WithRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	})
}
func WithRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	})
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}
func WithGetCustomParams[K comparable, V any](customParams interface{}) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithGetRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithSetCustomParams[K comparable, V any](customParams interface{}) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithSetRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	})
}
func WithDeleteCustomParams[K comparable, V any](customParams interface{}) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithDeleteRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	})
}

type rootOptionsImpl[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	keyCodec                 trcache.KeyCodec[K]
	name                     string
	redisDelFunc             RedisDelFunc[K, V]
	redisGetFunc             RedisGetFunc[K, V]
	redisSetFunc             RedisSetFunc[K, V]
	validator                trcache.Validator[V]
	valueCodec               trcache.Codec[V]
}

var _ options[string, string] = &rootOptionsImpl[string, string]{}

func (o *rootOptionsImpl[K, V]) OptCallDefaultDeleteOptions(options ...trcache.DeleteOption) {
	o.callDefaultDeleteOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultGetOptions(options ...trcache.GetOption) {
	o.callDefaultGetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultSetOptions(options ...trcache.SetOption) {
	o.callDefaultSetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptDefaultDuration(duration time.Duration) {
	o.defaultDuration = duration
}
func (o *rootOptionsImpl[K, V]) OptKeyCodec(keyCodec trcache.KeyCodec[K]) {
	o.keyCodec = keyCodec
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
func (o *rootOptionsImpl[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}
func (o *rootOptionsImpl[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}
func (o *rootOptionsImpl[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}
func (o *rootOptionsImpl[K, V]) OptValueCodec(valueCodec trcache.Codec[V]) {
	o.valueCodec = valueCodec
}

type getOptionsImpl[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []interface{}
	customParams  interface{}
	redisGetFunc  RedisGetFunc[K, V]
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptCustomOptions(customOptions []interface{}) {
	o.customOptions = customOptions
}
func (o *getOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *getOptionsImpl[K, V]) OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) {
	o.redisGetFunc = redisGetFunc
}

type setOptionsImpl[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	customParams interface{}
	duration     time.Duration
	redisSetFunc RedisSetFunc[K, V]
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}
func (o *setOptionsImpl[K, V]) OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) {
	o.redisSetFunc = redisSetFunc
}

type deleteOptionsImpl[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
	customParams interface{}
	redisDelFunc RedisDelFunc[K, V]
}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}

func (o *deleteOptionsImpl[K, V]) OptCustomParams(customParams interface{}) {
	o.customParams = customParams
}
func (o *deleteOptionsImpl[K, V]) OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) {
	o.redisDelFunc = redisDelFunc
}
