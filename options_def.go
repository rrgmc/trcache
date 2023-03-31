package trcache

//
// Option
//

type Option interface {
	ApplyCacheOpt(any) bool
}

type OptionFunc func(any) bool

func (o OptionFunc) ApplyCacheOpt(c any) bool {
	return o(c)
}

//
// Root Options
//

type IsRootOption interface {
	isCacheRootOption()
}

type IsRootOptionImpl struct {
}

func (i IsRootOptionImpl) isCacheOption() {}

type IsRootOptions interface {
	isCacheOptions()
}

type IsRootOptionsImpl struct {
}

func (i IsRootOptionsImpl) isCacheOptions() {}

type RootOption interface {
	IsRootOption
	Option
}

type RootOptionFunc func(any) bool

func (f RootOptionFunc) isCacheRootOption() {}

func (f RootOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RootOption = RootOptionFunc(func(a any) bool {
	return true
})

// Root options: builder base

type RootOptionBuilderBase struct {
	IsRootOption
	optionBuilder[RootOption]
}

// Root options: functions

func ParseRootOptions(obj IsRootOptions, options ...[]RootOption) error {
	return parseOptions(obj, options...)
}

func AppendRootOptions[K comparable, V any](options ...[]RootOption) []RootOption {
	return appendOptions(options...)
}

//
// Get options
//

type IsGetOption interface {
	isCacheGetOption()
}

type IsGetOptionImpl struct {
}

func (i IsGetOptionImpl) isCacheGetOption() {}

type IsGetOptions interface {
	isCacheGetOptions()
}

type IsGetOptionsImpl struct {
}

func (i IsGetOptionsImpl) isCacheGetOptions() {}

type GetOption interface {
	IsGetOption
	Option
}

type GetOptionFunc func(any) bool

func (f GetOptionFunc) isCacheGetOption() {}

func (f GetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ GetOption = GetOptionFunc(func(a any) bool {
	return true
})

// Get options: builder base

type GetOptionBuilderBase struct {
	IsGetOption
	optionBuilder[GetOption]
}

// Get options: functions

func ParseGetOptions(obj IsGetOptions, options ...[]GetOption) error {
	return parseOptions(obj, options...)
}

func AppendGetOptions(options ...[]GetOption) []GetOption {
	return appendOptions(options...)
}

//
// Set options
//

type IsSetOption interface {
	isCacheSetOption()
}

type IsSetOptionImpl struct {
}

func (i IsSetOptionImpl) isCacheSetOption() {}

type IsSetOptions interface {
	isCacheSetOptions()
}

type IsSetOptionsImpl struct {
}

func (i IsSetOptionsImpl) isCacheSetOptions() {}

type SetOption interface {
	IsSetOption
	Option
}

type SetOptionFunc func(any) bool

func (f SetOptionFunc) isCacheSetOption() {}

func (f SetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ SetOption = SetOptionFunc(func(a any) bool {
	return true
})

// Set options: builder base

type SetOptionBuilderBase struct {
	IsSetOption
	optionBuilder[SetOption]
}

// Set options: functions

func ParseSetOptions(obj IsSetOptions, options ...[]SetOption) error {
	return parseOptions(obj, options...)
}

func AppendSetOptions(options ...[]SetOption) []SetOption {
	return appendOptions(options...)
}

//
// Delete options
//

type IsDeleteOption interface {
	isCacheDeleteOption()
}

type IsDeleteOptionImpl struct {
}

func (i IsDeleteOptionImpl) isCacheDeleteOption() {}

type IsDeleteOptions interface {
	isCacheDeleteOptions()
}

type IsDeleteOptionsImpl struct {
}

func (i IsDeleteOptionsImpl) isCacheDeleteOptions() {}

type DeleteOption interface {
	IsDeleteOption
	Option
}

type DeleteOptionFunc func(any) bool

func (f DeleteOptionFunc) isCacheDeleteOption() {}

func (f DeleteOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ DeleteOption = DeleteOptionFunc(func(a any) bool {
	return true
})

// Delete options: builder base

type DeleteOptionBuilderBase struct {
	IsDeleteOption
	optionBuilder[DeleteOption]
}

// Cache delete options: functions

func ParseDeleteOptions(obj IsDeleteOptions, options ...[]DeleteOption) error {
	return parseOptions(obj, options...)
}

func AppendDeleteOptions(options ...[]DeleteOption) []DeleteOption {
	return appendOptions(options...)
}

//
// Refresh options
//

type IsRefreshOption interface {
	isCacheRefreshOption()
}

type IsRefreshOptionImpl struct {
}

func (i IsRefreshOptionImpl) isCacheRefreshOption() {}

type IsRefreshOptions interface {
	isCacheRefreshOptions()
}

type IsRefreshOptionsImpl struct {
}

func (i IsRefreshOptionsImpl) isCacheRefreshOptions() {}

type RefreshOption interface {
	IsRefreshOption
	Option
}

type RefreshOptionFunc func(any) bool

func (f RefreshOptionFunc) isCacheRefreshOption() {}

func (f RefreshOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RefreshOption = RefreshOptionFunc(func(a any) bool {
	return true
})

// Refresh options: builder base

type RefreshOptionBuilderBase struct {
	IsRefreshOption
	optionBuilder[RefreshOption]
}

// Refresh options: functions

func ParseRefreshOptions(obj IsRefreshOptions, options ...[]RefreshOption) error {
	return parseOptions(obj, options...)
}

func AppendRefreshOptions(options ...[]RefreshOption) []RefreshOption {
	return appendOptions(options...)
}
