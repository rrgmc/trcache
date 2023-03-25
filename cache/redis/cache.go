package redis

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/redis/go-redis/v9"
)

type Cache[K comparable, V any] struct {
	redis           *redis.Client
	name            string
	keycodec        trcache.KeyCodec[K]
	valueCodec      trcache.Codec[V]
	validator       trcache.Validator[V]
	defaultDuration time.Duration
}

func New[K comparable, V any](redis *redis.Client, option ...Option[K, V]) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		redis:           redis,
		defaultDuration: 0, // 0 means default for go-redis
	}
	for _, opt := range option {
		opt(ret)
	}
	if ret.valueCodec == nil {
		return nil, errors.New("value codec is required")
	}
	if ret.keycodec == nil {
		ret.keycodec = codec.NewStringKeyCodec[K]()
	}
	return ret, nil
}

func (c *Cache[K, V]) Name() string {
	return c.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...trcache.CacheGetOption[K, V]) (V, error) {
	var optns trcache.CacheGetOptions[K, V]
	trcache.ParseCacheGetOptions([]any{&optns}, options...)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		var empty V
		return empty, err
	}

	value, err := c.redis.Get(ctx, keyValue).Result()
	if err != nil {
		var empty V
		if errors.Is(err, redis.Nil) {
			return empty, trcache.ErrNotFound
		}
		return empty, err
	}

	dec, err := c.valueCodec.Unmarshal(ctx, value)
	if err != nil {
		var empty V
		return empty, trcache.CodecError{err}
	}

	if c.validator != nil {
		if err = c.validator.ValidateGet(ctx, dec); err != nil {
			var empty V
			return empty, err
		}
	}

	return dec, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption[K, V]) error {
	var optns trcache.CacheSetOptions[K, V]
	trcache.ParseCacheSetOptions([]any{&optns}, options...)

	enc, err := c.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return c.redis.Set(ctx, keyValue, enc, c.defaultDuration).Err()
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return c.redis.Del(ctx, keyValue).Err()
}

func (c *Cache[K, V]) parseKey(ctx context.Context, key K) (string, error) {
	keyValue, err := c.keycodec.Convert(ctx, key)
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
