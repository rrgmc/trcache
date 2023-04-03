package trcache

import (
	"go.uber.org/multierr"
)

// parse options

func parseOptions[O Option](obj any, options ...[]O) ParseOptionsResult {
	checkers := parseOptionsCheckers(options...)

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
	} else if len(checkers) == 1 {
		retErr = checkers[0].CheckCacheError()
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

// append options

func appendOptions[O Option](options ...[]O) []O {
	var ret []O
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// checker

type optionChecker[O Option] interface {
	Option
	CheckCacheOpt(opt Option)
	CheckCacheError() error
	isCacheRootOption()
	isCacheGetOption()
	isCacheSetOption()
	isCacheDeleteOption()
	isCacheRefreshOption()
}

func NewOptionChecker[O Option](options []O) optionChecker[O] {
	return &optionCheckerImpl[O]{
		Check: options,
	}
}

type optionCheckerImpl[O Option] struct {
	Check []O
	optns map[uint64]Option
}

func (o *optionCheckerImpl[O]) ApplyCacheOpt(a any) bool {
	return true
}

func (o *optionCheckerImpl[O]) CacheOptName() string {
	return "checker"
}

func (o *optionCheckerImpl[O]) CacheOptHash() uint64 {
	return 1
}

func (o *optionCheckerImpl[O]) CheckCacheOpt(opt Option) {
	if o.optns == nil {
		o.optns = map[uint64]Option{}
	}

	if _, ok := o.optns[opt.CacheOptHash()]; !ok {
		o.optns[opt.CacheOptHash()] = opt
	}
}

func (o *optionCheckerImpl[O]) CheckCacheError() error {
	var err error
	for _, opt := range o.Check {
		if _, ok := o.optns[opt.CacheOptHash()]; !ok {
			err = multierr.Append(err, NewOptionNotSupportedError(opt))

		}
	}
	return err
}

func (o *optionCheckerImpl[O]) isCacheRootOption()    {}
func (o *optionCheckerImpl[O]) isCacheGetOption()     {}
func (o *optionCheckerImpl[O]) isCacheSetOption()     {}
func (o *optionCheckerImpl[O]) isCacheDeleteOption()  {}
func (o *optionCheckerImpl[O]) isCacheRefreshOption() {}

// var _ optionCheckerImpl[Option] = &optionCheckerImpl[Option]{}

func parseOptionsCheckers[O Option](options ...[]O) []optionChecker[O] {
	var checkers []optionChecker[O]
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if oc, ok := any(opt).(optionChecker[O]); ok {
				checkers = append(checkers, oc)
			}
		}
	}
	return checkers
}
