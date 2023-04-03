package trcache

import (
	"go.uber.org/multierr"
)

// parse options

func parseOptions[O Option](obj any, options ...[]O) error {
	var checkers []optionChecker

	var err error
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if oc, ok := any(opt).(optionChecker); ok {
				checkers = append(checkers, oc)
			}
			if !opt.ApplyCacheOpt(obj) {
				err = multierr.Append(err, NewOptionNotSupportedError(opt))
			} else {
				for _, chk := range checkers {
					chk.CheckCacheOpt(opt)
				}
			}
		}
	}
	if len(checkers) > 0 {
		return nil
	}
	return err
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

type optionChecker interface {
	Option
	CheckCacheOpt(opt Option)
	isCacheRootOption()
	isCacheGetOption()
	isCacheSetOption()
	isCacheDeleteOption()
	isCacheRefreshOption()
}

type OptionChecker struct {
	optns map[uint64]Option
}

func (o *OptionChecker) ApplyCacheOpt(a any) bool {
	return true
}

func (o *OptionChecker) CacheOptHash() uint64 {
	return 1
}

func (o *OptionChecker) CheckCacheOpt(opt Option) {
	if o.optns == nil {
		o.optns = map[uint64]Option{}
	}

	if _, ok := o.optns[opt.CacheOptHash()]; !ok {
		o.optns[opt.CacheOptHash()] = opt
	}
}

func (o *OptionChecker) isCacheRootOption()    {}
func (o *OptionChecker) isCacheGetOption()     {}
func (o *OptionChecker) isCacheSetOption()     {}
func (o *OptionChecker) isCacheDeleteOption()  {}
func (o *OptionChecker) isCacheRefreshOption() {}

var _ optionChecker = &OptionChecker{}
