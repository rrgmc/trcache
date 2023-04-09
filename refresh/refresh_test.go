package refresh

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRefresh(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// will set after refresh
	mockCache.EXPECT().
		Set(mock.Anything, "a", "123").
		Return(nil)

	helper, err := NewHelper[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return "123", nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	value, err := helper.GetOrRefresh(ctx, mockCache, "a")
	require.NoError(t, err)

	require.Equal(t, "123", value)
}

func TestNoLoader(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)

	helper, err := NewHelper[string, string]()
	require.NoError(t, err)

	ctx := context.Background()

	_, err = helper.GetOrRefresh(ctx, mockCache, "a")
	require.Error(t, err)
}

func TestGetError(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	getErr := errors.New("get error")

	// get will error
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", getErr)

	helper, err := NewHelper[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return "123", nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	_, err = helper.GetOrRefresh(ctx, mockCache, "a")
	require.ErrorIs(t, err, getErr)
}

func TestSetError(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	setErr := errors.New("set error")

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// set will error
	mockCache.EXPECT().
		Set(mock.Anything, "a", "123").
		Return(setErr)

	helper, err := NewHelper[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return "123", nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	_, err = helper.GetOrRefresh(ctx, mockCache, "a")
	require.Error(t, setErr)
}

func TestCustomLoader(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// will set after refresh
	mockCache.EXPECT().
		Set(mock.Anything, "a", "456").
		Return(nil)

	helper, err := NewHelper[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return "123", nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	value, err := helper.GetOrRefresh(ctx, mockCache, "a",
		trcache.WithRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return "456", nil
		}))
	require.NoError(t, err)

	require.Equal(t, "456", value)
}

func TestRefreshData(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// will set after refresh
	mockCache.EXPECT().
		Set(mock.Anything, "a", "123456").
		Return(nil)

	helper, err := NewHelper[string, string](
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprint("123", options.Data), nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	value, err := helper.GetOrRefresh(ctx, mockCache, "a",
		trcache.WithRefreshData[string, string]("456"))
	require.NoError(t, err)

	require.Equal(t, "123456", value)
}

func TestRefreshCallDefault(t *testing.T) {
	mockCache := mocks.NewCache[string, string](t)

	// get will not find
	mockCache.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// will set after refresh
	mockCache.EXPECT().
		Set(mock.Anything, "a", "123789").
		Return(nil)

	helper, err := NewHelper[string, string](
		trcache.WithCallDefaultRefreshOptions[string, string](
			trcache.WithRefreshData[string, string]("789"),
		),
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprint("123", options.Data), nil
		}),
	)
	require.NoError(t, err)

	ctx := context.Background()

	value, err := helper.GetOrRefresh(ctx, mockCache, "a")
	require.NoError(t, err)

	require.Equal(t, "123789", value)
}

func TestInvalidOption(t *testing.T) {
	_, err := NewHelper[string, string](
		trcache.WithCallDefaultSetOptions[string, string](
			trcache.WithSetDuration[string, string](time.Minute),
		),
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprint("123", options.Data), nil
		}),
	)
	var optErr trcache.OptionNotSupportedError
	require.ErrorAs(t, err, &optErr)
}
