package trfreecache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/codec"
	"github.com/RangelReale/trcache/mocks"
	"github.com/coocood/freecache"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	cache := freecache.NewCache(512)

	c, err := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
		WithValueCodec[string, string](codec.NewJSONCodec[string]()),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	_ = cache.Del([]byte("a"))

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

	cache := freecache.NewCache(512)

	c, err := New[string, string](cache,
		WithValueCodec[string, string](codec.NewJSONCodec[string]()),
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

	cache := freecache.NewCache(512)

	c, err := New[string, string](cache,
		WithValueCodec[string, string](codec.NewJSONCodec[string]()),
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

	_ = cache.Del([]byte("a"))

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheJSONCodec(t *testing.T) {
	ctx := context.Background()

	cache := freecache.NewCache(512)

	c, err := New[string, string](cache,
		WithValueCodec[string, string](codec.NewJSONCodec[string]()),
		WithDefaultDuration[string, string](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)
}

func TestCacheJSONCodecInt(t *testing.T) {
	ctx := context.Background()

	cache := freecache.NewCache(512)

	c, err := New[string, int](cache,
		WithValueCodec[string, int](codec.NewJSONCodec[int]()),
		WithDefaultDuration[string, int](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", 12)
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, 12, v)
}

func TestCacheCodecInvalidInt(t *testing.T) {
	ctx := context.Background()

	cache := freecache.NewCache(512)

	c, err := New[string, int](cache,
		WithValueCodec[string, int](codec.NewForwardCodec[int]()),
		WithDefaultDuration[string, int](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", 12)
	require.ErrorAs(t, err, new(*trcache.InvalidValueTypeError))
}

func TestCacheRefresh(t *testing.T) {
	ctx := context.Background()

	cache := freecache.NewCache(512)

	c, err := NewRefresh[string, string, int](cache,
		WithValueCodec[string, string](codec.NewJSONCodec[string]()),
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
