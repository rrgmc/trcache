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
	OptName(name string)
	OptKeyCodec(keyCodec trcache.KeyCodec[K])
	OptValueCodec(valueCodec trcache.Codec[V])
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
	OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V])
	OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V])
	OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V])
}

// Cache get options

// +troptgen get
type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptCustomParams(customParams any)
	OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V])
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
	OptCustomParams(customParams any)
	OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V])
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
	OptCustomParams(customParams any)
	OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V])
}

// helpers

func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFuncFunc(fn RedisDelFuncFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](fn))
	return ob
}

//go:generate troptgen
