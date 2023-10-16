package chain

import (
	"context"
	"fmt"
	"testing"

	"github.com/rrgmc/trcache"
	"github.com/rrgmc/trcache/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)
	mockCache3 := mocks.NewCache[string, string](t)

	// first cache will not find
	mockCache1.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// second cache will find
	mockCache2.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("12", nil)

	// first cache will receive the found value
	mockCache1.EXPECT().
		Set(mock.Anything, "a", "12", mock.Anything).
		Return(nil)

	c, err := New[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2, mockCache3,
	},
		// set all default strategies
		WithGetStrategy[string, string](&GetStrategyGetFirstSetPrevious[string, string]{}),
		WithSetStrategy[string, string](&SetStrategySetAll[string, string]{}),
		WithDeleteStrategy[string, string](&DeleteStrategyDeleteAll[string, string]{}),
	)
	require.NoError(t, err)

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)
}

func TestChainRefresh(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)

	// no cache will find
	mockCache1.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	mockCache2.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)

	// refresh will be called

	// all cache will be set
	mockCache1.EXPECT().
		Set(mock.Anything, "a", "abc123", mock.Anything).
		Return(nil)
	mockCache2.EXPECT().
		Set(mock.Anything, "a", "abc123", mock.Anything).
		Return(nil)

	c, err := NewRefresh[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2,
	},
		trcache.WithDefaultRefreshFunc[string, string](func(ctx context.Context, key string,
			options trcache.RefreshFuncOptions) (string, error) {
			return fmt.Sprintf("abc%d", options.Data), nil
		}),
	)
	require.NoError(t, err)

	value, err := c.GetOrRefresh(ctx, "a", trcache.WithRefreshData[string, string](123))
	require.NoError(t, err)
	require.Equal(t, "abc123", value)
}

func TestChainStrategyCallback(t *testing.T) {
	ctx := context.Background()

	mockCache1 := mocks.NewCache[string, string](t)
	mockCache2 := mocks.NewCache[string, string](t)
	mockCache3 := mocks.NewCache[string, string](t)

	mockCache1.EXPECT().Name().Return("mockCache1")
	mockCache2.EXPECT().Name().Return("mockCache2")
	mockCache3.EXPECT().Name().Return("mockCache3")

	// first cache will not find
	mockCache1.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("", trcache.ErrNotFound)
	// second cache will find
	mockCache2.EXPECT().
		Get(mock.Anything, "a", mock.Anything, mock.Anything).
		Return("12", nil)

	// first cache will receive the found value
	mockCache1.EXPECT().
		Set(mock.Anything, "a", "12", mock.Anything).
		Return(nil)

	mockCache1.EXPECT().
		Set(mock.Anything, "b", "23", mock.Anything).
		Return(nil)
	mockCache2.EXPECT().
		Set(mock.Anything, "b", "23", mock.Anything).
		Return(nil)
	mockCache3.EXPECT().
		Set(mock.Anything, "b", "23", mock.Anything).
		Return(nil)

	mockCache1.EXPECT().
		Delete(mock.Anything, "b", mock.Anything).
		Return(nil)
	mockCache2.EXPECT().
		Delete(mock.Anything, "b", mock.Anything).
		Return(nil)
	mockCache3.EXPECT().
		Delete(mock.Anything, "b", mock.Anything).
		Return(nil)

	getCalled1, getCalled2, getCalled3 := false, false, false
	getSetCalled1, getSetCalled2, getSetCalled3 := false, false, false
	setCalled1, setCalled2, setCalled3 := false, false, false
	deleteCalled1, deleteCalled2, deleteCalled3 := false, false, false

	c, err := New[string, string]([]trcache.Cache[string, string]{
		mockCache1, mockCache2, mockCache3,
	},
		// set all default strategies
		WithDefaultStrategyCallback[string, string](&StrategyCallbackFunc{
			GetFn: func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterResult) {
				if cacheIdx == 0 && cacheName == "mockCache1" && err != nil {
					getCalled1 = true
					return
				}
				if cacheIdx == 1 && cacheName == "mockCache2" && err == nil {
					getCalled2 = true
					return
				}
				if cacheIdx == 2 && cacheName == "mockCache3" && err == nil {
					getCalled3 = true
					return
				}
			},
			GetSetFn: func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result GetStrategyAfterSetResult) {
				if cacheIdx == 0 && cacheName == "mockCache1" && err == nil {
					getSetCalled1 = true
					return
				}
				if cacheIdx == 1 && cacheName == "mockCache2" && err == nil {
					getSetCalled2 = true
					return
				}
				if cacheIdx == 2 && cacheName == "mockCache3" && err == nil {
					getSetCalled3 = true
					return
				}
			},
			SetFn: func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result SetStrategyAfterResult) {
				if cacheIdx == 0 && cacheName == "mockCache1" && err == nil {
					setCalled1 = true
					return
				}
				if cacheIdx == 1 && cacheName == "mockCache2" && err == nil {
					setCalled2 = true
					return
				}
				if cacheIdx == 2 && cacheName == "mockCache3" && err == nil {
					setCalled3 = true
					return
				}
			},
			DeleteFn: func(ctx context.Context, cacheIdx int, cacheName string, key any, err error, result DeleteStrategyAfterResult) {
				if cacheIdx == 0 && cacheName == "mockCache1" && err == nil {
					deleteCalled1 = true
					return
				}
				if cacheIdx == 1 && cacheName == "mockCache2" && err == nil {
					deleteCalled2 = true
					return
				}
				if cacheIdx == 2 && cacheName == "mockCache3" && err == nil {
					deleteCalled3 = true
					return
				}
			},
		}),
	)
	require.NoError(t, err)

	value, err := c.Get(ctx, "a")
	require.NoError(t, err)
	require.Equal(t, "12", value)

	require.True(t, getCalled1)
	require.True(t, getCalled2)
	require.False(t, getCalled3)

	require.True(t, getSetCalled1)
	require.False(t, getSetCalled2)
	require.False(t, getSetCalled3)

	err = c.Set(ctx, "b", "23")
	require.NoError(t, err)

	require.True(t, setCalled1)
	require.True(t, setCalled2)
	require.True(t, setCalled3)

	err = c.Delete(ctx, "b")
	require.NoError(t, err)

	require.True(t, deleteCalled1)
	require.True(t, deleteCalled2)
	require.True(t, deleteCalled3)
}
