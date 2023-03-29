package trcache

import "go.uber.org/multierr"

// options builder

type optionBuilderBase[O any] struct {
	opt   []O
	apply func(O, any) bool
}

func (ob *optionBuilderBase[O]) AppendOptions(opt ...O) {
	ob.opt = append(ob.opt, opt...)
}

func (ob *optionBuilderBase[O]) doApply(o any) bool {
	found := false
	for _, opt := range ob.opt {
		if ok := ob.apply(opt, o); ok {
			found = true
		}
	}
	return found
}

// parse options

func parseOptions[I any, O any](obj I, apply func(O, I) bool, options ...[]O) error {
	var err error
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if !apply(opt, obj) {
				err = multierr.Append(err, NewOptionNotSupportedError(opt))
			}
		}
	}
	return err
}
