package trgocache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/RangelReale/trcache/mocks"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	gocache := cache.New(5*time.Minute, 10*time.Minute)

	c, err := New[string, string](gocache,
		WithDefaultDuration[string, string](time.Minute),
		WithValueCodec[string, string](codec.NewForwardCodec[string]()),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	gocache.Delete("a")

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

	c, err := New[string, string](cache.New(5*time.Minute, 10*time.Minute),
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

	gocache := cache.New(5*time.Minute, 10*time.Minute)

	c, err := New[string, string](gocache,
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

	gocache.Delete("a")

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheRefresh(t *testing.T) {
	ctx := context.Background()

	gocache := cache.New(5*time.Minute, 10*time.Minute)

	c, err := NewRefresh[string, string, int](gocache,
		WithValueCodec[string, string](codec.NewForwardCodec[string]()),
		WithDefaultDuration[string, string](time.Minute),
		trcache.WithDefaultRefreshFunc[string, string, int](func(ctx context.Context, key string, options trcache.RefreshFuncOptions[int]) (string, error) {
			return fmt.Sprintf("abc%d", options.Data), nil
		}),
	)
	require.NoError(t, err)

	value, err := c.GetOrRefresh(ctx, "a", trcache.WithRefreshData[string, string, int](123))
	require.NoError(t, err)
	require.Equal(t, "abc123", value)
}
