// Code generated by troptgen. DO NOT EDIT.

package trgocache

import (
	trcache "github.com/RangelReale/trcache"
	"time"
)

type RootOption = trcache.RootOption

func WithCallDefaultDeleteOptions[K comparable, V any](options ...trcache.DeleteOption) RootOption {
	return trcache.WithCallDefaultDeleteOptions[K, V](options...)
}
func WithCallDefaultGetOptions[K comparable, V any](options ...trcache.GetOption) RootOption {
	return trcache.WithCallDefaultGetOptions[K, V](options...)
}
func WithCallDefaultSetOptions[K comparable, V any](options ...trcache.SetOption) RootOption {
	return trcache.WithCallDefaultSetOptions[K, V](options...)
}
func WithDefaultDuration[K comparable, V any](duration time.Duration) RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/gocache/options.DefaultDuration"
	const optionHash = uint64(0x360142a555d80c4b)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptDefaultDuration(duration)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithKeyCodec[K comparable, V any](keyCodec trcache.KeyCodec[K]) RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/gocache/options.KeyCodec"
	const optionHash = uint64(0x24c0d0030feea421)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptKeyCodec(keyCodec)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithName[K comparable, V any](name string) RootOption {
	return trcache.WithName[K, V](name)
}
func WithValidator[K comparable, V any](validator trcache.Validator[V]) RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/gocache/options.Validator"
	const optionHash = uint64(0xa6226955eb6acc5e)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValidator(validator)
			return true
		}
		return false
	}, optionName, optionHash)
}
func WithValueCodec[K comparable, V any](valueCodec trcache.Codec[V]) RootOption {
	const optionName = "github.com/RangelReale/trcache/cache/gocache/options.ValueCodec"
	const optionHash = uint64(0xb4eb97184a079517)
	return trcache.RootOptionFunc(func(o any) bool {
		switch opt := o.(type) {
		case options[K, V]:
			opt.OptValueCodec(valueCodec)
			return true
		}
		return false
	}, optionName, optionHash)
}

type GetOption = trcache.GetOption

func WithGetNoop[K comparable, V any](x bool) GetOption {
	return trcache.WithGetNoop[K, V](x)
}

type SetOption = trcache.SetOption

func WithSetDuration[K comparable, V any](duration time.Duration) SetOption {
	return trcache.WithSetDuration[K, V](duration)
}

type DeleteOption = trcache.DeleteOption
type rootOptionsImpl[K comparable, V any] struct {
	callDefaultDeleteOptions []trcache.DeleteOption
	callDefaultGetOptions    []trcache.GetOption
	callDefaultSetOptions    []trcache.SetOption
	defaultDuration          time.Duration
	keyCodec                 trcache.KeyCodec[K]
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
func (o *rootOptionsImpl[K, V]) OptKeyCodec(keyCodec trcache.KeyCodec[K]) {
	o.keyCodec = keyCodec
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

type getOptionsImpl[K comparable, V any] struct {
	noop bool
}

var _ getOptions[string, string] = &getOptionsImpl[string, string]{}

func (o *getOptionsImpl[K, V]) OptNoop(x bool) {
	o.noop = x
}

type setOptionsImpl[K comparable, V any] struct {
	duration time.Duration
}

var _ setOptions[string, string] = &setOptionsImpl[string, string]{}

func (o *setOptionsImpl[K, V]) OptDuration(duration time.Duration) {
	o.duration = duration
}

type deleteOptionsImpl[K comparable, V any] struct{}

var _ deleteOptions[string, string] = &deleteOptionsImpl[string, string]{}
