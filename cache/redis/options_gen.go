// Code generated by generator, DO NOT EDIT.
package trredis

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithCallDefaultDeleteOptions[K comparable, V any](options ...trcache.DeleteOption) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptCallDefaultDeleteOptions(options...)
			return true
		}
		return false
	})
}
func WithCallDefaultGetOptions[K comparable, V any](options ...trcache.GetOption) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptCallDefaultGetOptions(options...)
			return true
		}
		return false
	})
}
func WithCallDefaultSetOptions[K comparable, V any](options ...trcache.SetOption) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptCallDefaultSetOptions(options...)
			return true
		}
		return false
	})
}
func WithDefaultDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	})
}
func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	})
}
func WithName[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}
func WithRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	})
}
func WithRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	})
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	})
}
func WithGetCustomOptions[K comparable, V any](customOptions []interface{}) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomOptions(customOptions)
			return true
		}
		return false
	})
}
func WithGetCustomParams[K comparable, V any](customParams interface{}) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithGetRedisGetFunc[K comparable, V any](redisGetFunc RedisGetFunc[K, V]) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptRedisGetFunc(redisGetFunc)
			return true
		}
		return false
	})
}
func WithSetCustomParams[K comparable, V any](customParams interface{}) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithSetDuration[K comparable, V any](duration time.Duration) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptDuration(duration)
			return true
		}
		return false
	})
}
func WithSetRedisSetFunc[K comparable, V any](redisSetFunc RedisSetFunc[K, V]) trcache.SetOption {
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case SetOptions[K, V]:
			opt.OptRedisSetFunc(redisSetFunc)
			return true
		}
		return false
	})
}
func WithDeleteCustomParams[K comparable, V any](customParams interface{}) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptCustomParams(customParams)
			return true
		}
		return false
	})
}
func WithDeleteRedisDelFunc[K comparable, V any](redisDelFunc RedisDelFunc[K, V]) trcache.DeleteOption {
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case DeleteOptions[K, V]:
			opt.OptRedisDelFunc(redisDelFunc)
			return true
		}
		return false
	})
}

type RootOptionBuilder[K comparable, V any] struct {
	trcache.RootOptionBuilderBase
}

func RootOpt[K comparable, V any]() *RootOptionBuilder[K, V] {
	return &RootOptionBuilder[K, V]{}
}
func (ob *RootOptionBuilder[K, V]) WithCallDefaultDeleteOptions(options ...trcache.DeleteOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultDeleteOptions[K, V](options...))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithCallDefaultGetOptions(options ...trcache.GetOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultGetOptions[K, V](options...))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithCallDefaultSetOptions(options ...trcache.SetOption) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithCallDefaultSetOptions[K, V](options...))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithDefaultDuration(duration time.Duration) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithDefaultDuration[K, V](duration))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithKeyCodec(keyCodec trcache.KeyCodec[K]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithKeyCodec[K, V](keyCodec))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithName(name string) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithName[K, V](name))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisDelFunc[K, V](redisDelFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisGetFunc[K, V](redisGetFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithRedisSetFunc[K, V](redisSetFunc))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithValidator(validator trcache.Validator[V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithValidator[K, V](validator))
	return ob
}
func (ob *RootOptionBuilder[K, V]) WithValueCodec(valueCodec trcache.Codec[V]) *RootOptionBuilder[K, V] {
	ob.AppendOptions(WithValueCodec[K, V](valueCodec))
	return ob
}

type GetOptionBuilder[K comparable, V any] struct {
	trcache.GetOptionBuilderBase
}

func GetOpt[K comparable, V any]() *GetOptionBuilder[K, V] {
	return &GetOptionBuilder[K, V]{}
}
func (ob *GetOptionBuilder[K, V]) WithGetCustomOptions(customOptions []interface{}) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetCustomOptions[K, V](customOptions))
	return ob
}
func (ob *GetOptionBuilder[K, V]) WithGetCustomParams(customParams interface{}) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetCustomParams[K, V](customParams))
	return ob
}
func (ob *GetOptionBuilder[K, V]) WithGetRedisGetFunc(redisGetFunc RedisGetFunc[K, V]) *GetOptionBuilder[K, V] {
	ob.AppendOptions(WithGetRedisGetFunc[K, V](redisGetFunc))
	return ob
}

type SetOptionBuilder[K comparable, V any] struct {
	trcache.SetOptionBuilderBase
}

func SetOpt[K comparable, V any]() *SetOptionBuilder[K, V] {
	return &SetOptionBuilder[K, V]{}
}
func (ob *SetOptionBuilder[K, V]) WithSetCustomParams(customParams interface{}) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetCustomParams[K, V](customParams))
	return ob
}
func (ob *SetOptionBuilder[K, V]) WithSetDuration(duration time.Duration) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetDuration[K, V](duration))
	return ob
}
func (ob *SetOptionBuilder[K, V]) WithSetRedisSetFunc(redisSetFunc RedisSetFunc[K, V]) *SetOptionBuilder[K, V] {
	ob.AppendOptions(WithSetRedisSetFunc[K, V](redisSetFunc))
	return ob
}

type DeleteOptionBuilder[K comparable, V any] struct {
	trcache.DeleteOptionBuilderBase
}

func DeleteOpt[K comparable, V any]() *DeleteOptionBuilder[K, V] {
	return &DeleteOptionBuilder[K, V]{}
}
func (ob *DeleteOptionBuilder[K, V]) WithDeleteCustomParams(customParams interface{}) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteCustomParams[K, V](customParams))
	return ob
}
func (ob *DeleteOptionBuilder[K, V]) WithDeleteRedisDelFunc(redisDelFunc RedisDelFunc[K, V]) *DeleteOptionBuilder[K, V] {
	ob.AppendOptions(WithDeleteRedisDelFunc[K, V](redisDelFunc))
	return ob
}