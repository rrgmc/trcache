package trcache

import "go.uber.org/multierr"

// options builder

type optionBuilder[O Option] struct {
	opt []O
}

func (ob *optionBuilder[O]) AppendOptions(opt ...O) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *optionBuilder[O]) doApply(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := opt.ApplyCacheOpt(o); ok {
			found = true
		}
	}
	return found
}

// parse options

func parseOptions[O Option](obj any, options ...[]O) error {
	var err error
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if !opt.ApplyCacheOpt(obj) {
				err = multierr.Append(err, NewOptionNotSupportedError(opt))
			}
		}
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
