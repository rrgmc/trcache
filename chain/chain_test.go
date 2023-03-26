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

	mockCache1.EXPECT().Get(mock.Anything, "a").Return("12", nil)
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
