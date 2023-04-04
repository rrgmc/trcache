package trcache

//
// Option
//

type Option interface {
	ApplyCacheOpt(any) bool
	CacheOptName() string
	CacheOptHash() uint64
}

type optionFunc struct {
	f func(any) bool
	n string
	h uint64
}

func (o optionFunc) ApplyCacheOpt(c any) bool {
	return o.f(c)
}

func (o optionFunc) CacheOptName() string {
	return o.n
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

func RootOptionFunc(f func(any) bool, name string, hash uint64) RootOption {
	return &rootOptionFunc{
		optionFunc: optionFunc{f, name, hash},
	}
}

type rootOptionFunc struct {
	optionFunc
}

func (f rootOptionFunc) isCacheRootOption() {}

// Root options: functions

func ParseRootOptions(obj any, options ...[]RootOption) ParseOptionsResult {
	return parseOptions(obj, options...)
}

func ParseRootOptionsChecker(checker OptionChecker, obj any, options ...[]RootOption) ParseOptionsResult {
	return parseOptions(obj, AppendRootOptionsChecker(checker, options...))
}

func AppendRootOptions(options ...[]RootOption) []RootOption {
	return appendOptions(options...)
}

func AppendRootOptionsChecker(checker OptionChecker, options ...[]RootOption) []RootOption {
	return append([]RootOption{checker}, AppendRootOptions(options...)...)
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

func GetOptionFunc(f func(any) bool, name string, hash uint64) GetOption {
	return &getOptionFunc{
		optionFunc: optionFunc{f, name, hash},
	}
}

type getOptionFunc struct {
	optionFunc
}

func (f getOptionFunc) isCacheGetOption() {}

// Get options: functions

func ParseGetOptions(obj any, options ...[]GetOption) ParseOptionsResult {
	return parseOptions(obj, options...)
}

func ParseGetOptionsChecker(checker OptionChecker, obj any, options ...[]GetOption) ParseOptionsResult {
	return parseOptions(obj, AppendGetOptionsChecker(checker, options...))
}

func AppendGetOptions(options ...[]GetOption) []GetOption {
	return appendOptions(options...)
}

func AppendGetOptionsChecker(checker OptionChecker, options ...[]GetOption) []GetOption {
	return append([]GetOption{checker}, AppendGetOptions(options...)...)
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

func SetOptionFunc(f func(any) bool, name string, hash uint64) SetOption {
	return &setOptionFunc{
		optionFunc: optionFunc{f, name, hash},
	}
}

type setOptionFunc struct {
	optionFunc
}

func (f setOptionFunc) isCacheSetOption() {}

// Set options: functions

func ParseSetOptions(obj any, options ...[]SetOption) ParseOptionsResult {
	return parseOptions(obj, options...)
}

func ParseSetOptionsChecker(checker OptionChecker, obj any, options ...[]SetOption) ParseOptionsResult {
	return parseOptions(obj, AppendSetOptionsChecker(checker, options...))
}

func AppendSetOptions(options ...[]SetOption) []SetOption {
	return appendOptions(options...)
}

func AppendSetOptionsChecker(checker OptionChecker, options ...[]SetOption) []SetOption {
	return append([]SetOption{checker}, AppendSetOptions(options...)...)
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

func DeleteOptionFunc(f func(any) bool, name string, hash uint64) DeleteOption {
	return &deleteOptionFunc{
		optionFunc: optionFunc{f, name, hash},
	}
}

type deleteOptionFunc struct {
	optionFunc
}

func (f deleteOptionFunc) isCacheDeleteOption() {}

// Cache delete options: functions

func ParseDeleteOptions(obj any, options ...[]DeleteOption) ParseOptionsResult {
	return parseOptions(obj, options...)
}

func ParseDeleteOptionsChecker(checker OptionChecker, obj any, options ...[]DeleteOption) ParseOptionsResult {
	return parseOptions(obj, AppendDeleteOptionsChecker(checker, options...))
}

func AppendDeleteOptions(options ...[]DeleteOption) []DeleteOption {
	return appendOptions(options...)
}

func AppendDeleteOptionsChecker(checker OptionChecker, options ...[]DeleteOption) []DeleteOption {
	return append([]DeleteOption{checker}, AppendDeleteOptions(options...)...)
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

func RefreshOptionFunc(f func(any) bool, name string, hash uint64) RefreshOption {
	return &refreshOptionFunc{
		optionFunc: optionFunc{f, name, hash},
	}
}

type refreshOptionFunc struct {
	optionFunc
}

func (f refreshOptionFunc) isCacheRefreshOption() {}

// Refresh options: functions

func ParseRefreshOptions(obj any, options ...[]RefreshOption) ParseOptionsResult {
	return parseOptions(obj, options...)
}

func ParseRefreshOptionsChecker(checker OptionChecker, obj any, options ...[]RefreshOption) ParseOptionsResult {
	return parseOptions(obj, AppendRefreshOptionsChecker(checker, options...))
}

func AppendRefreshOptions(options ...[]RefreshOption) []RefreshOption {
	return appendOptions(options...)
}

func AppendRefreshOptionsChecker(checker OptionChecker, options ...[]RefreshOption) []RefreshOption {
	return append([]RefreshOption{checker}, AppendRefreshOptions(options...)...)
}

// Any option

type AnyOption interface {
	RootOption
	GetOption
	SetOption
	DeleteOption
	RefreshOption
}

type IsAnyOption struct {
	IsRootOption
	IsGetOption
	IsSetOption
	IsDeleteOption
	IsRefreshOption
}
