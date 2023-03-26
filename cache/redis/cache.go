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

func New[K comparable, V any](redis *redis.Client, options ...trcache.CacheOption[K, V]) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		redis: redis,
		options: cacheOptions[K, V]{
			defaultDuration: 0, // 0 means default for go-redis
			getFunc:         DefaultGetFunc[K, V]{},
			setFunc:         DefaultSetFunc[K, V]{},
			delFunc:         DefaultDelFunc[K, V]{},
		},
	}
	_ = trcache.ParseCacheOptions[K, V](&ret.options, options)
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

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns cacheGetOptions[K, V]
	_ = trcache.ParseCacheGetOptions(&optns, c.options.fnDefaultGet, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		var empty V
		return empty, err
	}

	value, err := c.options.getFunc.Get(ctx, c, keyValue, optns.customParams)
	if err != nil {
		var empty V
		return empty, err
	}

	// value, err := c.redis.Get(ctx, keyValue).Result()
	// if err != nil {
	// 	var empty V
	// 	if errors.Is(err, redis.Nil) {
	// 		return empty, trcache.ErrNotFound
	// 	}
	// 	return empty, err
	// }

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

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	var optns cacheSetOptions[K, V]
	_ = trcache.ParseCacheSetOptions(&optns, c.options.fnDefaultSet, options)

	enc, err := c.options.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return c.options.setFunc.Set(ctx, c, keyValue, enc, c.options.defaultDuration, optns.customParams)

	// return c.redis.Set(ctx, keyValue, enc, c.options.defaultDuration).Err()
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return c.redis.Del(ctx, keyValue).Err()
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
