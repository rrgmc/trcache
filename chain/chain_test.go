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

	mockCache := mocks.NewCache[string, string](t)

	mockCache.EXPECT().Get(mock.Anything, "a").Return("12", nil)

	c := NewChain[string, string]([]trcache.Cache[string, string]{
		mockCache,
	})

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)
}
