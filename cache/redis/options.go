package trredis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type Options[K comparable, V any] interface {
	trcache.IsOptions
	trcache.CallDefaultOptions[K, V]
	OptName(string)
	OptKeyCodec(trcache.KeyCodec[K])
	OptValueCodec(trcache.Codec[V])
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
	OptRedisGetFunc(RedisGetFunc[K, V])
	OptRedisSetFunc(RedisSetFunc[K, V])
	OptRedisDelFunc(RedisDelFunc[K, V])
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsOptionsImpl
	fnDefaultGet    []trcache.GetOption[K, V]
	fnDefaultSet    []trcache.SetOption[K, V]
	fnDefaultDelete []trcache.DeleteOption[K, V]
	name            string
	keyCodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
	redisGetFunc    RedisGetFunc[K, V]
	redisSetFunc    RedisSetFunc[K, V]
	redisDelFunc    RedisDelFunc[K, V]
}

var _ Options[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptCallDefaultGetOpt(i []trcache.GetOption[K, V]) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultSetOpt(i []trcache.SetOption[K, V]) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultDeleteOpt(i []trcache.DeleteOption[K, V]) {
	c.fnDefaultDelete = i
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

func (c *cacheOptions[K, V]) OptRedisGetFunc(fn RedisGetFunc[K, V]) {
	c.redisGetFunc = fn
}

func (c *cacheOptions[K, V]) OptRedisSetFunc(fn RedisSetFunc[K, V]) {
	c.redisSetFunc = fn
}

func (c *cacheOptions[K, V]) OptRedisDelFunc(fn RedisDelFunc[K, V]) {
	c.redisDelFunc = fn
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

func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}

func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.Option[K, V] {
	return trcache.OptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptDefaultDuration(defaultDuration)
			return true
		}
		return false
	})
}

// Options builder

type OptionBuilder[K comparable, V any] struct {
	trcache.OptionBuilderBase[K, V]
}

func Opt[K comparable, V any]() *OptionBuilder[K, V] {
	return &OptionBuilder[K, V]{}
}

func (ob *OptionBuilder[K, V]) WithName(name string) *OptionBuilder[K, V] {
	ob.AppendOptions(WithName[K, V](name))
	return ob
}

func (ob *OptionBuilder[K, V]) WithKeyCodec(keyCodec trcache.KeyCodec[K]) *OptionBuilder[K, V] {
	ob.AppendOptions(WithKeyCodec[K, V](keyCodec))
	return ob
}

func (ob *OptionBuilder[K, V]) WithValueCodec(valueCodec trcache.Codec[V]) *OptionBuilder[K, V] {
	ob.AppendOptions(WithValueCodec[K, V](valueCodec))
	return ob
}

func (ob *OptionBuilder[K, V]) WithValidator(validator trcache.Validator[V]) *OptionBuilder[K, V] {
	ob.AppendOptions(WithValidator[K, V](validator))
	return ob
}

func (ob *OptionBuilder[K, V]) WithDefaultDuration(defaultDuration time.Duration) *OptionBuilder[K, V] {
	ob.AppendOptions(WithDefaultDuration[K, V](defaultDuration))
	return ob
}

// Cache get options

type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptCustomParams(any)
	OptRedisGetFunc(RedisGetFunc[K, V])
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []any
	customParams  any
	redisGetFunc  RedisGetFunc[K, V]
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (c *getOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *getOptions[K, V]) OptCustomParams(params any) {
	c.customParams = params
}

func (c *getOptions[K, V]) OptRedisGetFunc(fn RedisGetFunc[K, V]) {
	c.redisGetFunc = fn
}

func WithGetCustomParam[K comparable, V any](param any) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomParams(param)
			return true
		}
		return false
	})
}

func WithGetRedisGetFunc[K comparable, V any](fn RedisGetFunc[K, V]) trcache.GetOption[K, V] {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptRedisGetFunc(fn)
			return true
		}
		return false
	})
}

// Options Get builder

type GetOptionBuilder[K comparable, V any] struct {
	trcache.GetOptionBuilderBase[K, V]
}

func GetOpt[K comparable, V any]() *GetOptionBuilder[K, V] {
	return &GetOptionBuilder[K, V]{}
}

func (ob *GetOptionBuilder[K, V]) WithGetRedisGetFunc(fn RedisGetFunc[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](fn))
	return ob
}

func (ob *GetOptionBuilder[K, V]) WithGetRedisGetFuncFunc(fn RedisGetFuncFunc[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](fn))
	return ob
}

// Cache set options

type SetOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
	OptCustomParams(any)
	OptRedisSetFunc(RedisSetFunc[K, V])
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	duration     time.Duration
	customParams any
	redisSetFunc RedisSetFunc[K, V]
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (c *setOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

func (c *setOptions[K, V]) OptCustomParams(customParams any) {
	c.customParams = customParams
}

func (c *setOptions[K, V]) OptRedisSetFunc(fn RedisSetFunc[K, V]) {
	c.redisSetFunc = fn
}

func WithSetRedisSetFunc[K comparable, V any](fn RedisSetFunc[K, V]) trcache.SetOption[K, V] {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptRedisSetFunc(fn)
			return true
		}
		return false
	})
}

// Options Set builder

type SetOptionBuilder[K comparable, V any] struct {
	trcache.SetOptionBuilderBase[K, V]
}

func SetOpt[K comparable, V any]() *SetOptionBuilder[K, V] {
	return &SetOptionBuilder[K, V]{}
}

func (ob *SetOptionBuilder[K, V]) WithSetRedisSetFunc(fn RedisSetFunc[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](fn))
	return ob
}

func (ob *SetOptionBuilder[K, V]) WithSetRedisSetFuncFunc(fn RedisSetFuncFunc[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](fn))
	return ob
}

// Cache delete options

type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
	OptCustomParams(any)
	OptRedisDelFunc(RedisDelFunc[K, V])
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
	customParams any
	redisDelFunc RedisDelFunc[K, V]
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}

func (c *deleteOptions[K, V]) OptCustomParams(customParams any) {
	c.customParams = customParams
}

func (c *deleteOptions[K, V]) OptRedisDelFunc(fn RedisDelFunc[K, V]) {
	c.redisDelFunc = fn
}

func WithDeleteRedisDelFunc[K comparable, V any](fn RedisDelFunc[K, V]) trcache.DeleteOption[K, V] {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptRedisDelFunc(fn)
			return true
		}
		return false
	})
}

// Options Delete builder

type DeleteOptionBuilder[K comparable, V any] struct {
	trcache.DeleteOptionBuilderBase[K, V]
}

func DeleteOpt[K comparable, V any]() *DeleteOptionBuilder[K, V] {
	return &DeleteOptionBuilder[K, V]{}
}

func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFunc(fn RedisDelFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](fn))
	return ob
}

func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFuncFunc(fn RedisDelFuncFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](fn))
	return ob
}
