package trcache

import "go.uber.org/multierr"

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
