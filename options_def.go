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

// Option interface

type IRootOpt int
type IGetOpt int
type ISetOpt int
type IDeleteOpt int
type IRefreshOpt int

type IOption[T any] interface {
	Option
	isCacheOption(T)
}

type IIsOption[T any] struct {
}

func (i IIsOption[T]) isCacheOption(T) {}

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
