package trcache

//
// Option
//

type Option interface {
	ApplyCacheOpt(any) bool
	CacheOptHash() uint64
}

type optionFunc struct {
	f func(any) bool
	h uint64
}

func (o optionFunc) ApplyCacheOpt(c any) bool {
	return o.f(c)
}

func (o optionFunc) CacheOptHash() uint64 {
	return o.h
}

var _ Option = &optionFunc{}

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

func RootOptionFunc(f func(any) bool, hash uint64) RootOption {
	return &rootOptionFunc{
		optionFunc: optionFunc{
			f: f,
			h: hash,
		},
	}
}

type rootOptionFunc struct {
	optionFunc
}

func (f rootOptionFunc) isCacheRootOption() {}

// var _ RootOption = &rootOptionFunc{optionFunc{func(a any) bool {
// 	return true
// }, 0}}

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

func GetOptionFunc(f func(any) bool, hash uint64) GetOption {
	return &getOptionFunc{
		optionFunc: optionFunc{
			f: f,
			h: hash,
		},
	}
}

type getOptionFunc struct {
	optionFunc
}

func (f getOptionFunc) isCacheGetOption() {}

// var _ GetOption = &getOptionFunc{optionFunc{func(a any) bool {
// 	return true
// }, 0}}

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

func SetOptionFunc(f func(any) bool, hash uint64) SetOption {
	return &setOptionFunc{
		optionFunc: optionFunc{
			f: f,
			h: hash,
		},
	}
}

type setOptionFunc struct {
	optionFunc
}

func (f setOptionFunc) isCacheSetOption() {}

// var _ SetOption = &setOptionFunc{optionFunc{func(a any) bool {
// 	return true
// }, 0}}

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

func DeleteOptionFunc(f func(any) bool, hash uint64) DeleteOption {
	return &deleteOptionFunc{
		optionFunc: optionFunc{
			f: f,
			h: hash,
		},
	}
}

type deleteOptionFunc struct {
	optionFunc
}

func (f deleteOptionFunc) isCacheDeleteOption() {}

// var _ DeleteOption = &deleteOptionFunc{optionFunc{func(a any) bool {
// 	return true
// }, 0}}

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

func RefreshOptionFunc(f func(any) bool, hash uint64) RefreshOption {
	return &refreshOptionFunc{
		optionFunc: optionFunc{
			f: f,
			h: hash,
		},
	}
}

type refreshOptionFunc struct {
	optionFunc
}

func (f refreshOptionFunc) isCacheRefreshOption() {}

// var _ RefreshOption = &refreshOptionFunc{optionFunc{func(a any) bool {
// 	return true
// }, 0}}

// Refresh options: functions

func ParseRefreshOptions(obj any, options ...[]RefreshOption) error {
	return parseOptions(obj, options...)
}

func AppendRefreshOptions(options ...[]RefreshOption) []RefreshOption {
	return appendOptions(options...)
}
