package chain

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type Options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(string)
	OptRefreshFunc(trcache.CacheRefreshFunc[K, V])
	OptSetPreviousOnGet(bool)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	fnDefaultGet     []trcache.GetOption
	fnDefaultSet     []trcache.SetOption
	fnDefaultDelete  []trcache.DeleteOption
	name             string
	refreshFunc      trcache.CacheRefreshFunc[K, V]
	setPreviousOnGet bool
}

var _ Options[string, string] = &cacheOptions[string, string]{}

func (c *cacheOptions[K, V]) OptCallDefaultGetOptions(i ...trcache.GetOption) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultSetOptions(i ...trcache.SetOption) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultDeleteOptions(i ...trcache.DeleteOption) {
	c.fnDefaultDelete = i
}

func (c *cacheOptions[K, V]) OptName(s string) {
	c.name = s
}

func (c *cacheOptions[K, V]) OptRefreshFunc(t trcache.CacheRefreshFunc[K, V]) {
	c.refreshFunc = t
}

func (c *cacheOptions[K, V]) OptSetPreviousOnGet(b bool) {
	c.setPreviousOnGet = b
}

// Cache get options

// +troptgen get
type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptSetOptions([]trcache.SetOption)
	OptGetStrategy(GetStrategy[K, V])
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []any
	setOptions    []trcache.SetOption
	getStrategy   GetStrategy[K, V]
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (c *getOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *getOptions[K, V]) OptSetOptions(i []trcache.SetOption) {
	c.setOptions = i
}

func (c *getOptions[K, V]) OptGetStrategy(s GetStrategy[K, V]) {
	c.getStrategy = s
}

// Cache set options

// +troptgen set
type SetOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
	OptSetStrategy(SetStrategy[K, V])
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	duration    time.Duration
	setStrategy SetStrategy[K, V]
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (c *setOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

func (c *setOptions[K, V]) OptSetStrategy(s SetStrategy[K, V]) {
	c.setStrategy = s
}

// Cache delete options

// +troptgen delete
type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
	OptDeleteStrategy(DeleteStrategy[K, V])
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
	deleteStrategy DeleteStrategy[K, V]
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}

func (c *deleteOptions[K, V]) OptDeleteStrategy(s DeleteStrategy[K, V]) {
	c.deleteStrategy = s
}

//go:generate troptgen
