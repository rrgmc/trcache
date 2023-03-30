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

	mockCache1.EXPECT().Name().Return("c1")
	// mockCache2.EXPECT().Get(mock.Anything, "a").Return("12", nil)

	c := New[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2,
	},
		trcache.WithCallDefaultRefreshOptions[string, string](
			trcache.WithRefreshData[string, string]("abc"),
		),
	)

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)
}

func TestChainGetStrategyDefault(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)
	mockCache3 := mocks.NewCache[string, string](t)

	// mockCache1.EXPECT().Name().Return("c1")
	// mockCache2.EXPECT().Name().Return("c2")
	// mockCache3.EXPECT().Name().Return("c3")

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
				WithGetStrategy[string, string](&GetStrategyGetFirstSetPrevious[string, string]{}),
			),
	)

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)
}
