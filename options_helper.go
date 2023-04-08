package trcache

import (
	"go.uber.org/multierr"
)

// parse options

// ParseOptions parses and applies a list of options, in order of appearance, and returns an error if any of
// the parameters was not accepted.
// If options of [OptionChecker] type exist in the list, the error is reported to them instead of being returned to
// the caller. See [OptionChecker] for details.
func ParseOptions[OA ~[]IOption[TO], TO any](obj any, options ...OA) ParseOptionsResult {
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
	}
	return ParseOptionsResult{
		err:     retErr,
		selfErr: err,
	}
}

// ParseOptionsResult is the result of [ParseOptions]
type ParseOptionsResult struct {
	err, selfErr error
}

// Err returns the options parsing error, if any.
func (r ParseOptionsResult) Err() error {
	return r.err
}

// SelfErr returns the options parsing error, if any, even if it was only sent to an [OptionChecker].
func (r ParseOptionsResult) SelfErr() error {
	return r.selfErr
}

// ConcatOptions concatenates multiple option lists.
func ConcatOptions[S ~[]O, O Option](options ...S) S {
	var ret S
	for _, opt := range options {
		ret = append(ret, opt...)
	}
	return ret
}

// ParseOptionsChecker parses the options using the passed [OptionChecker] to report usage.
func ParseOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO], obj any) ParseOptionsResult {
	return ParseOptions(obj, ConcatOptionsChecker[O, TO](checker, checker.CheckCacheOptList()))
}

// ConcatOptionsChecker concatenates multiple option lists and prepends an [OptionChecker] to it.
func ConcatOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO], options ...[]IOption[TO]) []IOption[TO] {
	return append([]IOption[TO]{checker}, ConcatOptions(options...)...)
}

// ForwardOptionsChecker returns the list of options from checker prepended by the checker itself.
func ForwardOptionsChecker[O IOption[TO], TO any](checker OptionChecker[O, TO]) []IOption[TO] {
	return append([]IOption[TO]{checker}, checker.CheckCacheOptList()...)
}

// OptionChecker is a special option type that is used to postpone checking the usage of options,
// to account for when the same options are used in multiple calls and we want to postpone the checking
// of valid options only at the end.
type OptionChecker[O IOption[TO], TO any] interface {
	IOption[TO]
	CheckCacheOpt(opt Option)
	CheckCacheError() error
	CheckCacheOptList() []IOption[TO]
}

// NewOptionChecker creates a new [OptionChecker] concatenating the list of options to check into it.
func NewOptionChecker[OA ~[]IOption[TO], TO any](options ...OA) OptionChecker[IOption[TO], TO] {
	return &optionCheckerImpl[IOption[TO], TO]{
		check: ConcatOptions(options...),
	}
}

// optionCheckerImpl is an implementation of [OptionChecker].
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

// parseOptionsCheckers returns [OptionChecker] implementations from a list of options.
func parseOptionsCheckers[OA ~[]IOption[TO], TO any](options ...OA) []OptionChecker[IOption[TO], TO] {
	var checkers []OptionChecker[IOption[TO], TO]
	for _, optinstance := range options {
		for _, opt := range optinstance {
			if oc, ok := any(opt).(OptionChecker[IOption[TO], TO]); ok {
				checkers = append(checkers, oc)
			}
		}
	}
	return checkers
}
