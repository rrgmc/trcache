package chain

import (
	"context"
	"testing"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)
	mockCache3 := mocks.NewCache[string, string](t)

	// first cache will not find
	mockCache1.EXPECT().Get(mock.Anything, "a", mock.Anything).Return("", trcache.ErrNotFound)
	// second cache will find
	mockCache2.EXPECT().Get(mock.Anything, "a", mock.Anything).Return("12", nil)

	// first cache will receive the found value
	mockCache1.EXPECT().Set(mock.Anything, "a", "12", mock.Anything).Return(nil)

	c := New[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2, mockCache3,
	},
		trcache.RootOpt[string, string]().
			WithCallDefaultGetOptions(
				WithGetGetStrategy[string, string](&GetStrategyGetFirstSetPrevious[string, string]{}),
			),
	)

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)
}

func TestChainRefresh(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)

	// no cache will find
	mockCache1.EXPECT().Get(mock.Anything, "a", mock.Anything).Return("", trcache.ErrNotFound)
	mockCache2.EXPECT().Get(mock.Anything, "a", mock.Anything).Return("", trcache.ErrNotFound)

	// refresh will be called

	// all cache will be set
	mockCache1.EXPECT().Set(mock.Anything, "a", "abc123", mock.Anything).Return(nil)
	mockCache2.EXPECT().Set(mock.Anything, "a", "abc123", mock.Anything).Return(nil)

	c := NewRefresh[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2,
	},
		trcache.WithCallDefaultRefreshOptions[string, string](trcache.RefreshOpt[string, string]().
			WithRefreshRefreshFunc(func(ctx context.Context, key string, options trcache.RefreshFuncOptions) (string, error) {
				return "abc" + options.Data.(string), nil
			}).
			WithRefreshData("123"),
		),
	)

	value, err := c.GetOrRefresh(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "abc123", value)
}

