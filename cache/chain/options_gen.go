// Code generated by troptgen. DO NOT EDIT.

package chain

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

func WithGetGetStrategy[K comparable, V any](getStrategy GetStrategy[K, V]) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/getOptions.GetStrategy"
	const optionHash = uint64(0x6ab81482970279a6)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptGetStrategy(getStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithGetSetOptions[K comparable, V any](options ...trcache.SetOption) trcache.GetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/getOptions.SetOptions"
	const optionHash = uint64(0x20cdc9d4030ddb85)
	return trcache.GetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case getOptions[K, V]:
			opt.OptSetOptions(options...)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithSetSetStrategy[K comparable, V any](setStrategy SetStrategy[K, V]) trcache.SetOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/setOptions.SetStrategy"
	const optionHash = uint64(0xfc4183c47cd45f1e)
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptSetStrategy(setStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithDeleteDeleteStrategy[K comparable, V any](deleteStrategy DeleteStrategy[K, V]) trcache.DeleteOption {
	const optionName = "github.com/RangelReale/trcache/cache/chain/deleteOptions.DeleteStrategy"
	const optionHash = uint64(0x562ad4dca6e88296)
	return trcache.DeleteOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case deleteOptions[K, V]:
			opt.OptDeleteStrategy(deleteStrategy)
			return true
		}
		return false
	}, optionName, optionHash)
}

type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	name                     string
}

var _ options[string, string] = &rootOptionsImpl[string, string]{}

func (o *rootOptionsImpl[K, V]) OptCallDefaultDeleteOptions(options ...trcache.DeleteOption) {
	o.callDefaultDeleteOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultGetOptions(options ...trcache.GetOption) {
	o.callDefaultGetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptCallDefaultSetOptions(options ...trcache.SetOption) {
	o.callDefaultSetOptions = options
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}

type getOptionsImpl[K comparable, V any] struct {
	getStrategy GetStrategy[K, V]
	setOptions  []trcache.SetOption
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptGetStrategy(getStrategy GetStrategy[K, V]) {
	o.getStrategy = getStrategy
}
func (o *getOptionsImpl[K, V]) OptSetOptions(options ...trcache.SetOption) {
	o.setOptions = options
}

type setOptionsImpl[K comparable, V any] struct {
	duration    time.Duration
	setStrategy SetStrategy[K, V]
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}
func (o *setOptionsImpl[K, V]) OptSetStrategy(setStrategy SetStrategy[K, V]) {
	o.setStrategy = setStrategy
}

type deleteOptionsImpl[K comparable, V any] struct {
	deleteStrategy DeleteStrategy[K, V]
}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}

func (o *deleteOptionsImpl[K, V]) OptDeleteStrategy(deleteStrategy DeleteStrategy[K, V]) {
	o.deleteStrategy = deleteStrategy
}
