// Code generated by troptgen. DO NOT EDIT.

package trristretto

import (
	trcache "github.com/rrgmc/trcache"
	"time"
)

func WithDefaultDuration[K comparable, V any](duration time.Duration) trcache.RootOption {
	const optionName = "github.com/rrgmc/trcache/cache/ristretto/options.DefaultDuration"
	const optionHash = uint64(0x2e7224a23dcd9543)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithEventualConsistency[K comparable, V any](eventualConsistency bool) trcache.RootOption {
	const optionName = "github.com/rrgmc/trcache/cache/ristretto/options.EventualConsistency"
	const optionHash = uint64(0x237a93cc1fe22bf6)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptEventualConsistency(eventualConsistency)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) trcache.RootOption {
	const optionName = "github.com/rrgmc/trcache/cache/ristretto/options.Validator"
	const optionHash = uint64(0x49fcb7b8d9427a66)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) trcache.RootOption {
	const optionName = "github.com/rrgmc/trcache/cache/ristretto/options.ValueCodec"
	const optionHash = uint64(0xffc96d014122347f)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithSetCost[K comparable, V any](cost int64) trcache.SetOption {
	const optionName = "github.com/rrgmc/trcache/cache/ristretto/setOptions.Cost"
	const optionHash = uint64(0xf8e2d52e844e33eb)
	return trcache.SetOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case setOptions[K, V]:
			opt.OptCost(cost)
			return true
		}
		return false
	}, optionName, optionHash)
}

type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	eventualConsistency      bool
	name                     string
	validator                trcache.Validator[V]
	valueCodec               trcache.Codec[V]
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
func (o *rootOptionsImpl[K, V]) OptDefaultDuration(duration time.Duration) {
	o.defaultDuration = duration
}
func (o *rootOptionsImpl[K, V]) OptEventualConsistency(eventualConsistency bool) {
	o.eventualConsistency = eventualConsistency
}
func (o *rootOptionsImpl[K, V]) OptName(name string) {
	o.name = name
}
func (o *rootOptionsImpl[K, V]) OptValidator(validator trcache.Validator[V]) {
	o.validator = validator
}
func (o *rootOptionsImpl[K, V]) OptValueCodec(valueCodec trcache.Codec[V]) {
	o.valueCodec = valueCodec
}

type getOptionsImpl[K comparable, V any] struct{}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

type setOptionsImpl[K comparable, V any] struct {
	cost     int64
	duration time.Duration
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptCost(cost int64) {
	o.cost = cost
}
func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}

type deleteOptionsImpl[K comparable, V any] struct{}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}
