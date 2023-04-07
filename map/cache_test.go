package trmap

import (
	"context"
	"fmt"
	"testing"

	"github.com/RangelReale/trcache"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	c, err := NewDefault[string, string]()
	require.NoError(t, err)

	err = c.Set(ctx, "a", "12")
	require.NoError(t, err)

	v, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", v)

	err = c.Delete(ctx, "a")
	require.NoError(t, err)

	v, err = c.Get(ctx, "a")
	require.ErrorIs(t, err, trcache.ErrNotFound)

	v, err = c.Get(ctx, "z")
	require.ErrorIs(t, err, trcache.ErrNotFound)
}

func TestCacheRefresh(t *testing.T) {
	ctx := context.Background()

	c, err := NewRefreshDefault[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string, options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprintf("abc%d", options.Data), nil
		}),
	)
	require.NoError(t, err)

	value, err := c.GetOrRefresh(ctx, "a", trcache.WithRefreshData[string, string](123))
	require.NoError(t, err)
	require.Equal(t, "abc123", value)
}
