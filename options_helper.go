package trcache

import (
	"go.uber.org/multierr"
)

// parse options

func parseOptions[O IOption[TO], TO any](obj any, options ...[]IOption[TO]) ParseOptionsResult {
	checkers := parseOptionsCheckers[O, TO](options...)

	var err error
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if !opt.ApplyCacheOpt(obj) {
				err = multierr.Append(err, NewOptionNotSupportedError(opt))
			} else {
				for _, chk := range checkers {
					chk.CheckCacheOpt(opt)
				}
			}
		}
	}

	var retErr error
	if len(checkers) == 0 {
		retErr = err
	}
	return ParseOptionsResult{
		err:     retErr,
		selfErr: err,
	}
}

type ParseOptionsResult struct {
	err, selfErr error
}

func (r ParseOptionsResult) Err() error {
	return r.err
}

func (r ParseOptionsResult) SelfErr() error {
	return r.selfErr
}

func ParseOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO], obj any) ParseOptionsResult {
	return parseOptions[IOption[TO], TO](obj, ConcatOptionsChecker[O, TO](checker, checker.CheckCacheOptList()))
}

func ConcatOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO], options ...[]IOption[TO]) []IOption[TO] {
	return append([]IOption[TO]{checker}, ConcatOptions(options...)...)
}

func ForwardOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO]) []IOption[TO] {
	return append([]IOption[TO]{checker}, checker.CheckCacheOptList()...)
}

// ConcatOptions

func ConcatOptions[S ~[]O, O Option](options ...S) S {
	var ret S
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// checker

type OptionChecker[O IOption[TO], TO any] interface {
	IOption[TO]
	CheckCacheOpt(opt Option)
	CheckCacheError() error
	CheckCacheOptList() []IOption[TO]
}

// func NewOptionChecker[S ~[]O, O Option](options ...S) OptionChecker[O] {
// 	return &optionCheckerImpl[O]{
// 		check: ConcatOptions(options...),
// 	}
// }

// func NewOptionChecker[S ~[]O, O IOption[TO], TO any](options ...S) OptionChecker[O, TO] {
// 	return &optionCheckerImpl[O, TO]{
// 		check: ConcatOptions(options...),
// 	}
// }

func NewOptionChecker[O IOption[TO], TO any](options ...[]IOption[TO]) OptionChecker[O, TO] {
	return &optionCheckerImpl[O, TO]{
		check: ConcatOptions(options...),
	}
}

// func NewOptionChecker[O IOption[TO], TO any](options ...IOption[TO]) OptionChecker[O, TO] {
// 	return &optionCheckerImpl[O, TO]{
// 		check: ConcatOptions[[]IOption[TO], IOption[TO]](options...),
// 	}
// }

type optionCheckerImpl[O IOption[TO], TO any] struct {
	IIsOption[TO]
	check []IOption[TO]
	optns map[uint64]Option
}

func (o *optionCheckerImpl[O, TO]) ApplyCacheOpt(a any) bool {
	return true
}

func (o *optionCheckerImpl[O, TO]) CacheOptName() string {
	return "checker"
}

func (o *optionCheckerImpl[O, TO]) CacheOptHash() uint64 {
	return 1
}

func (o *optionCheckerImpl[O, TO]) CheckCacheOpt(opt Option) {
	if o.optns == nil {
		o.optns = map[uint64]Option{}
	}

	if _, ok := o.optns[opt.CacheOptHash()]; !ok {
		o.optns[opt.CacheOptHash()] = opt
	}
}

func (o *optionCheckerImpl[O, TO]) CheckCacheError() error {
	var err error
	for _, opt := range o.check {
		if _, ok := o.optns[opt.CacheOptHash()]; !ok {
			err = multierr.Append(err, NewOptionNotSupportedError(opt))
		}
	}
	return err
}

func (o *optionCheckerImpl[O, TO]) CheckCacheOptList() []IOption[TO] {
	return o.check
}

func parseOptionsCheckers[O IOption[TO], TO any](options ...[]IOption[TO]) []OptionChecker[O, TO] {
	var checkers []OptionChecker[O, TO]
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if oc, ok := any(opt).(OptionChecker[O, TO]); ok {
				checkers = append(checkers, oc)
			}
		}
	}
	return checkers
}
