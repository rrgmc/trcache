package trredis

import (
	"context"
	"errors"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/redis/go-redis/v9"
)

type GetFunc[K comparable, V any] interface {
	Get(ctx context.Context, c *Cache[K, V], keyValue string, customParams any) (string, error)
}

type DefaultGetFunc[K comparable, V any] struct {
}

func (f DefaultGetFunc[K, V]) Get(ctx context.Context, c *Cache[K, V], keyValue string, _ any) (string, error) {
	value, err := c.Handle().Get(ctx, keyValue).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", trcache.ErrNotFound
		}
		return "", err
	}
	return value, nil
}

type SetFunc[K comparable, V any] interface {
	Set(ctx context.Context, c *Cache[K, V], keyValue string, value any, expiration time.Duration, customParams any) error
}

type DefaultSetFunc[K comparable, V any] struct {
}

func (f DefaultSetFunc[K, V]) Set(ctx context.Context, c *Cache[K, V], keyValue string, value any,
	expiration time.Duration, _ any) error {
	return c.Handle().Set(ctx, keyValue, value, expiration).Err()
}

type DeleteFunc[K comparable, V any] interface {
	Delete(ctx context.Context, c *Cache[K, V], keyValue string, customParams any) error
}

type DefaultDeleteFunc[K comparable, V any] struct {
}

func (f DefaultDeleteFunc[K, V]) Delete(ctx context.Context, c *Cache[K, V], keyValue string, _ any) error {
	return c.Handle().Del(ctx, keyValue).Err()
}
