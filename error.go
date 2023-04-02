package trcache

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("not found")     // key was not found
	ErrNotSupported = errors.New("not supported") // operation not supported
)

// CodecError represents an error from a [Codec] or a [KeyCodec].
type CodecError struct {
	Err error
}

func (e CodecError) Error() string {
	return e.Err.Error()
}

func (e CodecError) Unwrap() error {
	return e.Err
}

// ValidationError represents an error from a [Validator].
type ValidationError struct {
	Err error
}

func (e ValidationError) Error() string {
	return e.Err.Error()
}

func (e ValidationError) Unwrap() error {
	return e.Err
}

// InvalidValueTypeError represents an invalid value type error.
type InvalidValueTypeError struct {
	Message string
}

func (e *InvalidValueTypeError) Error() string {
	return e.Message
}

// OptionNotSupportedError represents that an options is not supported by the implementation.
type OptionNotSupportedError struct {
	Message string
	Option  Option
}

func NewOptionNotSupportedError(option Option) OptionNotSupportedError {
	return OptionNotSupportedError{
		Message: fmt.Sprintf("option not supported: %s (%T)", getName(option), option),
		Option:  option,
	}
}

func (e OptionNotSupportedError) Error() string {
	return e.Message
}
