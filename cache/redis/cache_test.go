package redis

import (
	"context"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/mocks"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// //go:generate mockery --srcpkg github.com/spf13/afero --name Fs

func TestCache(t *testing.T) {
	ctx := context.Background()

	redis, mockRedis := redismock.NewClientMock()

	mockRedis.ExpectSet("a", "12", time.Minute).SetVal("12")
	mockRedis.ExpectGet("a").SetVal("12")
	mockRedis.ExpectGet("a").RedisNil() // simulate expiration
	mockRedis.ExpectGet("z").RedisNil()

	c, err := NewCache[string, string](redis,
		WithValueCodec[string, string](trcache.NewForwardCodec[string]()),
		WithDefaultDuration[string, string](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	if err := mockRedis.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCacheValidator(t *testing.T) {
	ctx := context.Background()

	redis, mockRedis := redismock.NewClientMock()
	mockValidator := mocks.NewValidator[string](t)

	mockRedis.ExpectSet("a", "12", time.Minute).SetVal("12")
	mockRedis.ExpectGet("a").SetVal("12")

	mockValidator.EXPECT().
		ValidateGet(mock.Anything, "12").
		Return(trcache.ErrNotFound).
		Once()

	c, err := NewCache[string, string](redis,
		WithValueCodec[string, string](trcache.NewForwardCodec[string]()),
		WithValidator[string, string](mockValidator),
		WithDefaultDuration[string, string](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	_, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	if err := mockRedis.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCacheJSONCodec(t *testing.T) {
	ctx := context.Background()

	redis, mockRedis := redismock.NewClientMock()

	mockRedis.ExpectSet("a", `"12"`, time.Minute).SetVal(`"12"`)
	mockRedis.ExpectGet("a").SetVal(`"12"`)
	mockRedis.ExpectGet("z").RedisNil()

	c, err := NewCache[string, string](redis,
		WithValueCodec[string, string](trcache.NewJSONCodec[string](trcache.WithJSONCodecReturnString(true))),
		WithDefaultDuration[string, string](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	if err := mockRedis.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestCacheJSONCodecInt(t *testing.T) {
	ctx := context.Background()

	redis, mockRedis := redismock.NewClientMock()

	mockRedis.ExpectSet("a", "12", time.Minute).SetVal("12")
	mockRedis.ExpectGet("a").SetVal("12")
	mockRedis.ExpectGet("z").RedisNil()

	c, err := NewCache[string, int](redis,
		WithValueCodec[string, int](trcache.NewJSONCodec[int](trcache.WithJSONCodecReturnString(true))),
		WithDefaultDuration[string, int](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", 12)
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, 12, v)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheCodecInvalidInt(t *testing.T) {
	ctx := context.Background()

	redis, mockRedis := redismock.NewClientMock()

	mockRedis.ExpectSet("a", 12, time.Minute).SetVal("12")
	mockRedis.ExpectGet("a").SetVal("12")

	c, err := NewCache[string, int](redis,
		WithValueCodec[string, int](trcache.NewForwardCodec[int]()),
		WithDefaultDuration[string, int](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", 12)
	require.NoError(t, err)

	_, err = c.Get(ctx, "a")
	require.ErrorAs(t, err, new(*trcache.ErrInvalidValueType))

	if err := mockRedis.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
