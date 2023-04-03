package trcache_test

import (
	"testing"

	"github.com/RangelReale/trcache"
)

// func TestOptions1(t *testing.T) {
// 	var opt Option
// 	opt = NewOptionCheckerImpl()
//
// 	if oc, ok := opt.(OptionChecker); ok {
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
	}
	_ = trcache.ParseRootOptions(&options, optionsParam)
	// require.NoError(t, err)
	// if optErr != nil && !options.ignoreOptionNotSupported {
	// 	return nil, optErr
	// }
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
	})
}

func With1Test12[K comparable, V any](value int) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions1[K, V]:
			opt.OptTest12(value)
			return true
		}
		return false
	})
}

type With1Test13 struct {
	trcache.IsRootOption
}

func (o With1Test13) isCacheRootOption() {
}

func (o With1Test13) ApplyCacheOpt(opt any) bool {
	return true
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

// Test Options 2

func With2Test21[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions2[K, V]:
			opt.OptTest21(name)
			return true
		}
		return false
	})
}

func With2Test22[K comparable, V any](value int) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case testOptions2[K, V]:
			opt.OptTest22(value)
			return true
		}
		return false
	})
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
