package trcache

//
// Option
//

// Option is the base interface for all options.
type Option interface {
	ApplyCacheOpt(any) bool
	CacheOptName() string
	CacheOptHash() uint64
}

// optionFunc is a functional implementation of [Option].
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

// Option type markers.
// As Go type declarations creates new unrelated types even if they have the same underlying type,
// we leverage this to make methods accept only one type and not the others.

type IRootOpt int
type IGetOpt int
type ISetOpt int
type IDeleteOpt int
type IRefreshOpt int

// IOption is used to segregate option types, not allowing different types to be mixed.
// The combination of a "isCacheOption" function and a single parameter of a type marker makes interfaces
// compatible only between the same marker type.
type IOption[T any] interface {
	Option
	isCacheOption(T)
}

// IIsOption is an implementation of IOption, meant to be embedded in structs.
type IIsOption[T any] struct {
}

func (i IIsOption[T]) isCacheOption(T) {}

// iOptionFunc is a functional implementation of IOption.
type iOptionFunc[T any] struct {
	IIsOption[T]
	optionFunc
}

//
// Root Options
//

type RootOption = IOption[IRootOpt]

type IsRootOption = IIsOption[IRootOpt]

func RootOptionFunc(f func(any) bool, name string, hash uint64) RootOption {
	return &iOptionFunc[IRootOpt]{
		optionFunc: optionFunc{f, name, hash},
	}
}

//
// Get options
//

type GetOption = IOption[IGetOpt]

type IsGetOption = IIsOption[IGetOpt]

func GetOptionFunc(f func(any) bool, name string, hash uint64) GetOption {
	return &iOptionFunc[IGetOpt]{
		optionFunc: optionFunc{f, name, hash},
	}
}

//
// Set options
//

type SetOption = IOption[ISetOpt]

type IsSetOption = IIsOption[ISetOpt]

func SetOptionFunc(f func(any) bool, name string, hash uint64) SetOption {
	return &iOptionFunc[ISetOpt]{
		optionFunc: optionFunc{f, name, hash},
	}
}

//
// Delete options
//

type DeleteOption = IOption[IDeleteOpt]

type IsDeleteOption = IIsOption[IDeleteOpt]

func DeleteOptionFunc(f func(any) bool, name string, hash uint64) DeleteOption {
	return &iOptionFunc[IDeleteOpt]{
		optionFunc: optionFunc{f, name, hash},
	}
}

//
// Refresh options
//

type RefreshOption = IOption[IRefreshOpt]

type IsRefreshOption = IIsOption[IRefreshOpt]

func RefreshOptionFunc(f func(any) bool, name string, hash uint64) RefreshOption {
	return &iOptionFunc[IRefreshOpt]{
		optionFunc: optionFunc{f, name, hash},
	}
}
