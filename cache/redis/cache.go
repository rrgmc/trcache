package redis

import (
	"context"
	"errors"

	"github.com/RangelReale/trcache"
	"github.com/redis/go-redis/v9"
)

type Cache[K comparable, V any] struct {
	redis      *redis.Client
	valueCodec trcache.Codec[V]
}

func NewCache[K comparable, V any](redis *redis.Client, option ...Option[K, V]) (*Cache[K, V], error) {
	ret := &Cache[K, V]{
		redis: redis,
	}
	for _, opt := range option {
		opt(ret)
	}
	if ret.valueCodec == nil {
		return nil, errors.New("value codec is required")
	}
	return ret, nil
}

func (c *Cache[K, V]) Get(ctx context.Context, key K) (V, error) {
	value, err := c.redis.Get(ctx, trcache.StringValue(key)).Result()
	if err != nil {
		var empty V
		return empty, err
	}
	ret, err := c.valueCodec.Unmarshal(ctx, value)
	if err != nil {
		var empty V
		return empty, trcache.CodecError{err}
	}
	return ret, nil
}

func (c *Cache[K, V]) Set(ctx context.Context, key K, value V, options ...trcache.CacheSetOption) error {
	value, err := c.valueCodec.Marshal(ctx, value)
	if err != nil {
		return trcache.CodecError{err}
	}
	return c.redis.Set(ctx, trcache.StringValue(key), value, 0).Err()
}

func (c *Cache[K, V]) Delete(ctx context.Context, key K) error {
	return c.redis.Del(ctx, trcache.StringValue(key)).Err()
}

func (c *Cache[K, V]) Clear(ctx context.Context, key K) error {
	return trcache.ErrNotSupported
}
