package trcache_test

import (
	"testing"

	"github.com/RangelReale/trcache"
	"github.com/stretchr/testify/require"
)

// func TestOptions1(t *testing.T) {
// 	var opt Option
// 	opt = NewOptionChecker()
//
// 	if oc, ok := opt.(optionCheckerImpl); ok {
//
// 	}
// }

func TestOptionsRecursive(t *testing.T) {
	options := testOptions1Impl[string, string]{}
	optionsParam := []trcache.RootOption{
		With1Test13{},
		With1Test11[string, string]("test11"),
		With1Test12[string, string](12),
		With2Test21[string, string]("test21"),
		With2Test22[string, string](22),
		// WithGet1Test115[string, string]("aaa"),
	}

	err := trcache.ParseRootOptions(&options,
		trcache.NewParseRootOptionChecker(optionsParam),
		trcache.NewParseRootOptionChecker(optionsParam),
		optionsParam)
	require.NoError(t, err.Err())

	err = trcache.ParseRootOptions(&options,
		trcache.NewParseRootOptionChecker(optionsParam),
		optionsParam)
	require.Error(t, err.Err())
}

// Test Options 1

func With1Test11[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions1[K, V]:
			opt.OptTest11(name)
			return true
		}
		return false
	}, "111", 111)
}

func With1Test12[K comparable, V any](value int) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions1[K, V]:
			opt.OptTest12(value)
			return true
		}
		return false
	}, "112", 112)
}

type With1Test13 struct {
	trcache.IsRootOption
}

func (o With1Test13) ApplyCacheOpt(opt any) bool {
	return true
}

func (o With1Test13) CacheOptName() string {
	return "113"
}

func (o With1Test13) CacheOptHash() uint64 {
	return 113
}

func WithGet1Test115[K comparable, V any](name string) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions1[K, V]:
			opt.OptTest11(name)
			return true
		}
		return false
	}, "115", 115)
}

type TestOptions1[K comparable, V any] interface {
	OptTest11(name string)
	OptTest12(value int)
}

type testOptions1[K comparable, V any] interface {
	TestOptions1[K, V]
}

type testOptions1Impl[K comparable, V any] struct {
	test11 string
	test12 int
}

var _ testOptions1[string, string] = &testOptions1Impl[string, string]{}

func (o *testOptions1Impl[K, V]) OptTest11(name string) {
	o.test11 = name
}

func (o *testOptions1Impl[K, V]) OptTest12(value int) {
	o.test12 = value
}

type TestGetOptions1[K comparable, V any] interface {
	OptGetTest15(name string)
	OptGetTest16(value int)
}

type testGetOptions1[K comparable, V any] interface {
	TestGetOptions1[K, V]
}

type testGetOptions1Impl[K comparable, V any] struct {
	test15 string
	test16 int
}

var _ testGetOptions1[string, string] = &testGetOptions1Impl[string, string]{}

func (o *testGetOptions1Impl[K, V]) OptGetTest15(name string) {
	o.test15 = name
}

func (o *testGetOptions1Impl[K, V]) OptGetTest16(value int) {
	o.test16 = value
}

// Test Options 2

func With2Test21[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions2[K, V]:
			opt.OptTest21(name)
			return true
		}
		return false
	}, "221", 221)
}

func With2Test22[K comparable, V any](value int) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions2[K, V]:
			opt.OptTest22(value)
			return true
		}
		return false
	}, "222", 222)
}

type TestOptions2[K comparable, V any] interface {
	OptTest21(name string)
	OptTest22(value int)
}

type testOptions2[K comparable, V any] interface {
	TestOptions2[K, V]
}

type testOptions2Impl[K comparable, V any] struct {
	test21 string
	test22 int
}

var _ testOptions2[string, string] = &testOptions2Impl[string, string]{}

func (o *testOptions2Impl[K, V]) OptTest21(name string) {
	o.test21 = name
}

func (o *testOptions2Impl[K, V]) OptTest22(value int) {
	o.test22 = value
}
