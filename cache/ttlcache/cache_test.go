package trttlcache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/rrgmc/trcache"
	"github.com/rrgmc/trcache/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	cache := ttlcache.New[string, string]()

	c, err := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	cache.Delete("a")

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

	c, err := New[string, string](ttlcache.New[string, string](),
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

	cache := ttlcache.New[string, string]()

	c, err := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
		trcache.WithCallDefaultGetOptions[string, string](
			WithGetTouch[string, string](true),
		),
	)
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	cache.Delete("a")

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheRefresh(t *testing.T) {
	ctx := context.Background()

	cache := ttlcache.New[string, string]()

	c, err := NewRefresh[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string, options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprintf("abc%d", options.Data), nil
		}),
	)
	require.NoError(t, err)

	value, err := c.GetOrRefresh(ctx, "a", trcache.WithRefreshData[string, string](123))
	require.NoError(t, err)
	require.Equal(t, "abc123", value)
}
