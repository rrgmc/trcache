// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	trcache "github.com/RangelReale/trcache"
	mock "github.com/stretchr/testify/mock"
)

// CacheRefreshFunc is an autogenerated mock type for the CacheRefreshFunc type
type CacheRefreshFunc[K comparable, V interface{}] struct {
	mock.Mock
}

type CacheRefreshFunc_Expecter[K comparable, V interface{}] struct {
	mock *mock.Mock
}

func (_m *CacheRefreshFunc[K, V]) EXPECT() *CacheRefreshFunc_Expecter[K, V] {
	return &CacheRefreshFunc_Expecter[K, V]{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, key, options
func (_m *CacheRefreshFunc[K, V]) Execute(ctx context.Context, key K, options trcache.CacheRefreshFuncOptions) (V, error) {
	ret := _m.Called(ctx, key, options)

	var r0 V
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, K, trcache.CacheRefreshFuncOptions) (V, error)); ok {
		return rf(ctx, key, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, K, trcache.CacheRefreshFuncOptions) V); ok {
		r0 = rf(ctx, key, options)
	} else {
		r0 = ret.Get(0).(V)
	}

	if rf, ok := ret.Get(1).(func(context.Context, K, trcache.CacheRefreshFuncOptions) error); ok {
		r1 = rf(ctx, key, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheRefreshFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type CacheRefreshFunc_Execute_Call[K comparable, V interface{}] struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - key K
//   - options trcache.CacheRefreshFuncOptions
func (_e *CacheRefreshFunc_Expecter[K, V]) Execute(ctx interface{}, key interface{}, options interface{}) *CacheRefreshFunc_Execute_Call[K, V] {
	return &CacheRefreshFunc_Execute_Call[K, V]{Call: _e.mock.On("Execute", ctx, key, options)}
}

func (_c *CacheRefreshFunc_Execute_Call[K, V]) Run(run func(ctx context.Context, key K, options trcache.CacheRefreshFuncOptions)) *CacheRefreshFunc_Execute_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(K), args[2].(trcache.CacheRefreshFuncOptions))
	})
	return _c
}

func (_c *CacheRefreshFunc_Execute_Call[K, V]) Return(_a0 V, _a1 error) *CacheRefreshFunc_Execute_Call[K, V] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheRefreshFunc_Execute_Call[K, V]) RunAndReturn(run func(context.Context, K, trcache.CacheRefreshFuncOptions) (V, error)) *CacheRefreshFunc_Execute_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCacheRefreshFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheRefreshFunc creates a new instance of CacheRefreshFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheRefreshFunc[K comparable, V interface{}](t mockConstructorTestingTNewCacheRefreshFunc) *CacheRefreshFunc[K, V] {
	mock := &CacheRefreshFunc[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}