package trcache

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

var ErrNotSupported = errors.New("not supported")

type CodecError struct {
	Err error
}

func (e CodecError) Error() string {
	return e.Err.Error()
}

func (e CodecError) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func (e CodecError) As(target any) bool {
	return errors.As(e.Err, target)
}

type ValidationError struct {
	Err error
}

func (e ValidationError) Error() string {
	return e.Err.Error()
}

func (e ValidationError) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func (e ValidationError) As(target any) bool {
	return errors.As(e.Err, target)
}

type ErrInvalidValueType struct {
	Message string
}

func (e *ErrInvalidValueType) Error() string {
	return e.Message
}

type ChainError struct {
	Message string
	Err     error
}

func NewChainError(message string, err error) ChainError {
	return ChainError{message, err}
}

func (e ChainError) Error() string {
	return e.Message
}

func (e ChainError) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func (e ChainError) As(target any) bool {
	return errors.As(e.Err, target)
}

type OptionNotSupportedError[O any] struct {
	Option O
}

func NewOptionNotSupportedError[O any](option O) OptionNotSupportedError[O] {
	return OptionNotSupportedError[O]{}
}

func (e OptionNotSupportedError[O]) Error() string {
	return "option not supported"
}
