package trttlcache

import (
	"time"

	"github.com/RangelReale/trcache"
)

// Option

type Options[K comparable, V any] interface {
	trcache.IsRootOptions
	trcache.CallDefaultOptions[K, V]
	OptName(string)
	OptValidator(trcache.Validator[V])
	OptDefaultDuration(time.Duration)
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

func (c *cacheOptions[K, V]) OptCallDefaultGetOpt(i []trcache.GetOption) {
	c.fnDefaultGet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultSetOpt(i []trcache.SetOption) {
	c.fnDefaultSet = i
}

func (c *cacheOptions[K, V]) OptCallDefaultDeleteOpt(i []trcache.DeleteOption) {
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

func WithName[K comparable, V any](name string) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptName(name)
			return true
		}
		return false
	})
}

func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	})
}

func WithDefaultDuration[K comparable, V any](defaultDuration time.Duration) trcache.RootOption {
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case Options[K, V]:
			opt.OptDefaultDuration(defaultDuration)
			return true
		}
		return false
	})
}

// Cache get options

type GetOptions[K comparable, V any] interface {
	trcache.IsGetOptions
	trcache.GetOptions[K, V]
	OptTouch(bool)
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

// Cache get options: declarations

func WithGetTouch[K comparable, V any](touch bool) trcache.GetOption {
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case GetOptions[K, V]:
			opt.OptTouch(touch)
			return true
		}
		return false
	})
}

// Cache set options

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

type DeleteOptions[K comparable, V any] interface {
	trcache.IsDeleteOptions
	trcache.DeleteOptions[K, V]
}

type deleteOptions[K comparable, V any] struct {
	trcache.IsDeleteOptionsImpl
}

var _ DeleteOptions[string, string] = &deleteOptions[string, string]{}
