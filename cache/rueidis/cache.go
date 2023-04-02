package trrueidis

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/RangelReale/trcache/wrap"
	"github.com/rueian/rueidis"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	redis   rueidis.Client
}

func New[K comparable, V any](redis rueidis.Client, options ...RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		redis: redis,
		options: rootOptionsImpl[K, V]{
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

func NewRefresh[K comparable, V any, RD any](redis rueidis.Client,
	options ...RootOption) (trcache.RefreshCache[K, V, RD], error) {
	cache, err := New[K, V](redis, options...)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](cache, options...), nil
}

func (c *Cache[K, V]) Handle() rueidis.Client {
	return c.redis
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K, options ...GetOption) (V, error) {
	optns := getOptionsImpl[K, V]{
		clientSideDuration: c.options.defaultClientSideDuration,
		redisGetFunc:       c.options.redisGetFunc,
	}
	_ = trcache.ParseGetOptions(&optns, c.options.callDefaultGetOptions, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		var empty V
		return empty, err
	}

	value, err := optns.redisGetFunc.Get(ctx, c, keyValue, optns.customParams, optns.clientSideDuration)
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

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...SetOption) error {
	optns := setOptionsImpl[K, V]{
		redisSetFunc: c.options.redisSetFunc,
		duration:     c.options.defaultDuration,
	}
	_ = trcache.ParseSetOptions(&optns, c.options.callDefaultSetOptions, options)

	enc, err := c.options.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	var strvalue string
	switch tv := enc.(type) {
	case string:
		strvalue = tv
	case []byte:
		strvalue = string(tv)
	default:
		return &trcache.InvalidValueTypeError{fmt.Sprintf("invalid type '%s' for redis value", getType(enc))}
	}

	return optns.redisSetFunc.Set(ctx, c, keyValue, strvalue, c.options.defaultDuration, optns.customParams)
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K, options ...DeleteOption) error {
	optns := deleteOptionsImpl[K, V]{
		redisDelFunc: c.options.redisDelFunc,
	}
	_ = trcache.ParseDeleteOptions(&optns, c.options.callDefaultDeleteOptions, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	return optns.redisDelFunc.Delete(ctx, c, keyValue, optns.customParams)
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
			&trcache.InvalidValueTypeError{fmt.Sprintf("invalid type '%s' for redis key", getType(keyValue))},
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
