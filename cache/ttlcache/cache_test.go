package ttlcache

import (
	"context"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/mocks"
	"github.com/jellydator/ttlcache/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	cache := ttlcache.New[string, string]()

	c := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
	)

	err := c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a",
		WithCacheGetTouch[string, string](true),
		trcache.WithCacheGetCustomOption[string, string](1))
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

	c := New[string, string](ttlcache.New[string, string](),
		WithDefaultDuration[string, string](time.Minute),
		WithValidator[string, string](mockValidator),
	)

	err := c.Set(ctx, "a", "12")
	require.NoError(t, err)

	_, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheOptions(t *testing.T) {
	ctx := context.Background()

	cache := ttlcache.New[string, string]()

	c := New[string, string](cache,
		WithDefaultDuration[string, string](time.Minute),
		WithDefaultGetOptions[string, string](
			WithCacheGetTouch[string, string](true),
		),
	)

	err := c.Set(ctx, "a", "12")
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
