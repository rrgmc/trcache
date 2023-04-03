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

type RootOption interface {
	Option
	isCacheRootOption()
}

type IsRootOption struct {
}

func (i IsRootOption) isCacheRootOption() {}

type RootOptionFunc func(any) bool

func (f RootOptionFunc) isCacheRootOption() {}

func (f RootOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RootOption = RootOptionFunc(func(a any) bool {
	return true
})

// Root options: functions

func ParseRootOptions(obj any, options ...[]RootOption) error {
	return parseOptions(obj, options...)
}

func AppendRootOptions(options ...[]RootOption) []RootOption {
	return appendOptions(options...)
}

//
// Get options
//

type GetOption interface {
	Option
	isCacheGetOption()
}

type IsGetOption struct {
}

func (i IsGetOption) isCacheGetOption() {}

type GetOptionFunc func(any) bool

func (f GetOptionFunc) isCacheGetOption() {}

func (f GetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ GetOption = GetOptionFunc(func(a any) bool {
	return true
})

// Get options: functions

func ParseGetOptions(obj any, options ...[]GetOption) error {
	return parseOptions(obj, options...)
}

func AppendGetOptions(options ...[]GetOption) []GetOption {
	return appendOptions(options...)
}

//
// Set options
//

type SetOption interface {
	Option
	isCacheSetOption()
}

type IsSetOption struct {
}

func (i IsSetOption) isCacheSetOption() {}

type SetOptionFunc func(any) bool

func (f SetOptionFunc) isCacheSetOption() {}

func (f SetOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ SetOption = SetOptionFunc(func(a any) bool {
	return true
})

// Set options: functions

func ParseSetOptions(obj any, options ...[]SetOption) error {
	return parseOptions(obj, options...)
}

func AppendSetOptions(options ...[]SetOption) []SetOption {
	return appendOptions(options...)
}

//
// Delete options
//

type DeleteOption interface {
	Option
	isCacheDeleteOption()
}

type IsDeleteOption struct {
}

func (i IsDeleteOption) isCacheDeleteOption() {}

type DeleteOptionFunc func(any) bool

func (f DeleteOptionFunc) isCacheDeleteOption() {}

func (f DeleteOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ DeleteOption = DeleteOptionFunc(func(a any) bool {
	return true
})

// Cache delete options: functions

func ParseDeleteOptions(obj any, options ...[]DeleteOption) error {
	return parseOptions(obj, options...)
}

func AppendDeleteOptions(options ...[]DeleteOption) []DeleteOption {
	return appendOptions(options...)
}

//
// Refresh options
//

type RefreshOption interface {
	Option
	isCacheRefreshOption()
}

type IsRefreshOption struct {
}

func (i IsRefreshOption) isCacheRefreshOption() {}

type RefreshOptionFunc func(any) bool

func (f RefreshOptionFunc) isCacheRefreshOption() {}

func (f RefreshOptionFunc) ApplyCacheOpt(c any) bool {
	return f(c)
}

var _ RefreshOption = RefreshOptionFunc(func(a any) bool {
	return true
})

// Refresh options: functions

func ParseRefreshOptions(obj any, options ...[]RefreshOption) error {
	return parseOptions(obj, options...)
}

func AppendRefreshOptions(options ...[]RefreshOption) []RefreshOption {
	return appendOptions(options...)
}
