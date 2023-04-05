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

func ParseRootOptionsChecker(checker OptionChecker[RootOption], obj any) ParseOptionsResult {
	return parseOptions(obj, ConcatRootOptionsChecker(checker, checker.CheckCacheOptList()))
}

func ConcatRootOptionsChecker(checker OptionChecker[RootOption], options ...[]RootOption) []RootOption {
	return append([]RootOption{checker}, ConcatOptions(options...)...)
}

func ForwardRootOptionsChecker(checker OptionChecker[RootOption]) []RootOption {
	return append([]RootOption{checker}, checker.CheckCacheOptList()...)
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

func ParseGetOptionsChecker(checker OptionChecker[GetOption], obj any) ParseOptionsResult {
	return parseOptions(obj, ConcatGetOptionsChecker(checker, checker.CheckCacheOptList()))
}

func ConcatGetOptionsChecker(checker OptionChecker[GetOption], options ...[]GetOption) []GetOption {
	return append([]GetOption{checker}, ConcatOptions(options...)...)
}

func ForwardGetOptionsChecker(checker OptionChecker[GetOption]) []GetOption {
	return append([]GetOption{checker}, checker.CheckCacheOptList()...)
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

func ParseSetOptionsChecker(checker OptionChecker[SetOption], obj any) ParseOptionsResult {
	return parseOptions(obj, ConcatSetOptionsChecker(checker, checker.CheckCacheOptList()))
}

func ConcatSetOptionsChecker(checker OptionChecker[SetOption], options ...[]SetOption) []SetOption {
	return append([]SetOption{checker}, ConcatOptions(options...)...)
}

func ForwardSetOptionsChecker(checker OptionChecker[SetOption]) []SetOption {
	return append([]SetOption{checker}, checker.CheckCacheOptList()...)
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

func ParseDeleteOptionsChecker(checker OptionChecker[DeleteOption], obj any) ParseOptionsResult {
	return parseOptions(obj, ConcatDeleteOptionsChecker(checker, checker.CheckCacheOptList()))
}

func ConcatDeleteOptionsChecker(checker OptionChecker[DeleteOption], options ...[]DeleteOption) []DeleteOption {
	return append([]DeleteOption{checker}, ConcatOptions(options...)...)
}

func ForwardDeleteOptionsChecker(checker OptionChecker[DeleteOption]) []DeleteOption {
	return append([]DeleteOption{checker}, checker.CheckCacheOptList()...)
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

func ParseRefreshOptionsChecker(checker OptionChecker[RefreshOption], obj any) ParseOptionsResult {
	return parseOptions(obj, ConcatRefreshOptionsChecker(checker, checker.CheckCacheOptList()))
}

func ConcatRefreshOptionsChecker(checker OptionChecker[RefreshOption], options ...[]RefreshOption) []RefreshOption {
	return append([]RefreshOption{checker}, ConcatOptions(options...)...)
}

func ForwardRefreshOptionsChecker(checker OptionChecker[RefreshOption]) []RefreshOption {
	return append([]RefreshOption{checker}, checker.CheckCacheOptList()...)
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
