package trredis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type Options[K comparable, V any] interface {
	trcache.IsRootOptions
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
	trcache.IsRootOptionsImpl
	fnDefaultGet    []trcache.GetOption
	fnDefaultSet    []trcache.SetOption
	fnDefaultDelete []trcache.DeleteOption
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

func (c *cacheOptions[K, V]) OptCallDefaultGetOptions(i ...trcache.GetOption) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultSetOptions(i ...trcache.SetOption) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultDeleteOptions(i ...trcache.DeleteOption) {
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

// Cache get options

// +troptgen get
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

// helpers

func (ob *GetOptionBuilder[K, V]) WithGetRedisGetFuncFunc(fn RedisGetFuncFunc[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](fn))
	return ob
}

// Cache set options

// +troptgen set
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

// helpers

func (ob *SetOptionBuilder[K, V]) WithSetRedisSetFuncFunc(fn RedisSetFuncFunc[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](fn))
	return ob
}

// Cache delete options

// +troptgen delete
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

// helpers

func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFuncFunc(fn RedisDelFuncFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](fn))
	return ob
}

//go:generate troptgen
