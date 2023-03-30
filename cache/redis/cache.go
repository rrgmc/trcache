package trredis

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/redis/go-redis/v9"
)

type Cache[K comparable, V any] struct {
	options cacheOptions[K, V]
	redis   *redis.Client
}

func New[K comparable, V any](redis *redis.Client, options ...trcache.RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		redis: redis,
		options: cacheOptions[K, V]{
			defaultDuration: 0, // 0 means default for go-redis
			redisGetFunc:    DefaultRedisGetFunc[K, V]{},
			redisSetFunc:    DefaultRedisSetFunc[K, V]{},
			redisDelFunc:    DefaultRedisDelFunc[K, V]{},
		},
	}
	_ = trcache.ParseRootOptions(&ret.options, options)
	if ret.options.valueCodec == nil {
		return nil, errors.New("value codec is required")
	}
	if ret.options.keyCodec == nil {
		ret.options.keyCodec = codec.NewStringKeyCodec[K]()
	}
	return ret, nil
}

func (c *Cache[K, V]) Handle() *redis.Client {
	return c.redis
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...trcache.GetOption) (V, error) {
	var optns getOptions[K, V]
	_ = trcache.ParseGetOptions(&optns, c.options.fnDefaultGet, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		var empty V
		return empty, err
	}

	value, err := FirstRedisGetFunc(optns.redisGetFunc, c.options.redisGetFunc).
		Get(ctx, c, keyValue, optns.customParams)
	if err != nil {
		var empty V
		return empty, err
	}

	dec, err := c.options.valueCodec.Unmarshal(ctx, value)
	if err != nil {
		var empty V
		return empty, trcache.CodecError{err}
	}

	if c.options.validator != nil {
		if err = c.options.validator.ValidateGet(ctx, dec); err != nil {
			var empty V
			return empty, err
		}
	}

	return dec, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.SetOption) error {
	var optns setOptions[K, V]
	_ = trcache.ParseSetOptions(&optns, c.options.fnDefaultSet, options)

	enc, err := c.options.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return FirstRedisSetFunc(optns.redisSetFunc, c.options.redisSetFunc).
		Set(ctx, c, keyValue, enc, c.options.defaultDuration, optns.customParams)
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K, options ...trcache.DeleteOption) error {
	var optns deleteOptions[K, V]
	_ = trcache.ParseDeleteOptions(&optns, c.options.fnDefaultDelete, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return FirstRedisDelFunc(optns.redisDelFunc, c.options.redisDelFunc).
		Delete(ctx, c, keyValue, optns.customParams)
}

func (c *Cache[K, V]) parseKey(ctx context.Context, key K) (string, error) {
	keyValue, err := c.options.keyCodec.Convert(ctx, key)
	if err != nil {
		return "", trcache.CodecError{err}
	}

	switch kv := keyValue.(type) {
	case string:
		return kv, nil
	case []byte:
		return string(kv), nil
	default:
		return "", trcache.CodecError{
			&trcache.ErrInvalidValueType{fmt.Sprintf("invalid type '%s' for redis key", getType(keyValue))},
		}
	}
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
