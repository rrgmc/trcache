package chain

import (
	"context"
	"testing"

	"github.com/RangelReale/trcache"
	"github.com/RangelReale/trcache/mocks"
	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	ctx := context.Background()

	mockCache := mocks.NewCache[string, string](t)

	c := NewChain[string, string]([]trcache.Cache[string, string]{
		mockCache,
	})

	_, err := c.Get(ctx, "a")
	require.NoError(t, err)
}
