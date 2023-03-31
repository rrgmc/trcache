package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

// +troptgen root
type Options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(name string)
	OptValidator(validator trcache.Validator[V])
	OptDefaultDuration(duration time.Duration)
}

type cacheOptions[K comparable, V any] struct {
	trcache.IsRootOptionsImpl
	fnDefaultGet    []trcache.GetOption
	fnDefaultSet    []trcache.SetOption
	fnDefaultDelete []trcache.DeleteOption
	name            string
	validator       trcache.Validator[V]
	defaultDuration time.Duration
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

func (c *cacheOptions[K, V]) OptValidator(t trcache.Validator[V]) {
	c.validator = t
}

func (c *cacheOptions[K, V]) OptDefaultDuration(duration time.Duration) {
	c.defaultDuration = duration
}

// Cache get options

// +troptgen get
type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptTouch(touch bool)
}

type getOptions[K comparable, V any] struct {
	trcache.IsGetOptionsImpl
	customOptions []any
	touch         bool
}

var _ GetOptions[string, string] = &getOptions[string, string]{}

func (c *getOptions[K, V]) OptCustomOptions(anies []any) {
	c.customOptions = anies
}

func (c *getOptions[K, V]) OptTouch(b bool) {
	c.touch = b
}

// Cache set options

// +troptgen set
type SetOptions[K comparable, V any] interface {
	trcache.IsSetOptions
	trcache.SetOptions[K, V]
}

type setOptions[K comparable, V any] struct {
	trcache.IsSetOptionsImpl
	duration time.Duration
}

var _ SetOptions[string, string] = &setOptions[string, string]{}

func (c *setOptions[K, V]) OptDuration(duration time.Duration) {
	c.duration = duration
}

// Cache delete options

// +troptgen delete
type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}

//go:generate troptgen
