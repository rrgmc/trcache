package trristretto

import (
	"context"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/RangelReale/trcache/mocks"
	"github.com/dgraph-io/ristretto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.NoError(t, err)

	c, err := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
		WithValueCodec[string, string](codec.NewForwardCodec[string]()),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	cache.Del("a")

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheValidator(t *testing.T) {
	ctx := context.Background()

	mockValidator := mocks.NewValidator[string](t)

	mockValidator.EXPECT().
		ValidateGet(mock.Anything, "12").
		Return(trcache.ErrNotFound).
		Once()

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.NoError(t, err)

	c, err := New[string, string](cache,
		WithValueCodec[string, string](codec.NewForwardCodec[string]()),
		WithDefaultDuration[string, string](time.Minute),
		WithValidator[string, string](mockValidator),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	_, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheOptions(t *testing.T) {
	ctx := context.Background()

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	require.NoError(t, err)

	c, err := New[string, string](cache,
		WithValueCodec[string, string](codec.NewForwardCodec[string]()),
		WithDefaultDuration[string, string](time.Minute),
		// redis.WithDefaultDuration[string, string](time.Minute),
		trcache.WithCallDefaultGetOptions[string, string](),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	cache.Del("a")

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}
