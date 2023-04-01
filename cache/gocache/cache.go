package trgocache

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/RangelReale/trcache/wrap"
	"github.com/patrickmn/go-cache"
)

type Cache[K comparable, V any] struct {
	options rootOptionsImpl[K, V]
	cache   *cache.Cache
}

var _ trcache.Cache[string, string] = &Cache[string, string]{}

func New[K comparable, V any](cache *cache.Cache,
	options ...trcache.RootOption) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		cache:   cache,
		options: rootOptionsImpl[K, V]{},
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

func NewRefresh[K comparable, V any, RD any](cache *cache.Cache,
	options ...trcache.RootOption) (trcache.RefreshCache[K, V, RD], error) {
	c, err := New[K, V](cache, options...)
	if err != nil {
		return nil, err
	}
	return wrap.NewWrapRefreshCache[K, V, RD](c, options...), nil
}

// func NewDefault[K comparable, V any](options ...trcache.RootOption) *Cache[K, V] {
// 	return New(cache.New(), options...)
// }

func (c *Cache[K, V]) Handle() *cache.Cache {
	return c.cache
}

func (c *Cache[K, V]) Name() string {
	return c.options.name
}

func (c *Cache[K, V]) Get(ctx context.Context, key K,
	options ...trcache.GetOption) (V, error) {
	var optns getOptionsImpl[K, V]
	_ = trcache.ParseGetOptions(&optns, c.options.callDefaultGetOptions, options)

	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		var empty V
		return empty, err
	}

	value, ok := c.cache.Get(keyValue)
	if !ok {
		var empty V
		return empty, trcache.ErrNotFound
	}

	dec, err := c.options.valueCodec.Unmarshal(ctx, value)
	if err != nil {
		var empty V
		return empty, trcache.CodecError{err}
	}

	if c.options.validator != nil {
		if err := c.options.validator.ValidateGet(ctx, dec); err != nil {
			var empty V
			return empty, err
		}
	}

	return dec, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V,
	options ...trcache.SetOption) error {
	optns := setOptionsImpl[K, V]{
		duration: c.options.defaultDuration,
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

	c.cache.Set(keyValue, enc, optns.duration)
	return nil
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K,
	options ...trcache.DeleteOption) error {
	keyValue, err := c.parseKey(ctx, key)
	if err != nil {
		return err
	}

	c.cache.Delete(keyValue)
	return nil
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
