package trcache

import (
	"go.uber.org/multierr"
)

// parse options

func parseOptions[O Option](obj any, options ...[]O) ParseOptionsResult {
	var checkers []optionChecker[O]
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if oc, ok := any(opt).(optionChecker[O]); ok {
				checkers = append(checkers, oc)
			}
		}
	}

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
	return ParseOptionsResult{
		isCheck: len(checkers) > 0,
		err:     err,
	}
}

type ParseOptionsResult struct {
	isCheck bool
	err     error
}

func (r ParseOptionsResult) Err() error {
	if !r.isCheck {
		return r.err
	}
	return nil
}

func (r ParseOptionsResult) CheckErr() error {
	return r.err
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
	isCacheRootOption()
	isCacheGetOption()
	isCacheSetOption()
	isCacheDeleteOption()
	isCacheRefreshOption()
}

type OptionChecker[O Option] struct {
	Check []O
	optns map[uint64]Option
}

func (o *OptionChecker[O]) ApplyCacheOpt(a any) bool {
	return true
}

func (o *OptionChecker[O]) CacheOptName() string {
	return "checker"
}

func (o *OptionChecker[O]) CacheOptHash() uint64 {
	return 1
}

func (o *OptionChecker[O]) CheckCacheOpt(opt Option) {
	if o.optns == nil {
		o.optns = map[uint64]Option{}
	}

	if _, ok := o.optns[opt.CacheOptHash()]; !ok {
		o.optns[opt.CacheOptHash()] = opt
	}
}

func (o *OptionChecker[O]) CheckError() error {
	var err error
	for _, opt := range o.Check {
		if _, ok := o.optns[opt.CacheOptHash()]; !ok {
			err = multierr.Append(err, NewOptionNotSupportedError(opt))

		}
	}
	return err
}

func (o *OptionChecker[O]) isCacheRootOption()    {}
func (o *OptionChecker[O]) isCacheGetOption()     {}
func (o *OptionChecker[O]) isCacheSetOption()     {}
func (o *OptionChecker[O]) isCacheDeleteOption()  {}
func (o *OptionChecker[O]) isCacheRefreshOption() {}

var _ optionChecker[Option] = &OptionChecker[Option]{}
