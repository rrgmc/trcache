package trredis

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.Options[K, V]
	trcache.CallDefaultOptions[K, V]
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
type getOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptCustomParams(customParams any)
	OptRedisGetFunc(redisGetFunc RedisGetFunc[K, V])
}

// helpers

func (ob *getOptionBuilder[K, V]) WithGetRedisGetFuncFunc(fn RedisGetFuncFunc[K, V]) *getOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](fn))
	return ob
}

// Cache set options

// +troptgen set
type setOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
	OptCustomParams(customParams any)
	OptRedisSetFunc(redisSetFunc RedisSetFunc[K, V])
}

// helpers

func (ob *setOptionBuilder[K, V]) WithSetRedisSetFuncFunc(fn RedisSetFuncFunc[K, V]) *setOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](fn))
	return ob
}

// Cache delete options

// +troptgen delete
type deleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
	OptCustomParams(customParams any)
	OptRedisDelFunc(redisDelFunc RedisDelFunc[K, V])
}

// helpers

func (ob *deleteOptionBuilder[K, V]) WithDeleteRedisDelFuncFunc(fn RedisDelFuncFunc[K, V]) *deleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](fn))
	return ob
}

//go:generate troptgen
