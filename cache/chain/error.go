package chain

import "errors"

var (
	ErrInvalidStrategyResult = errors.New("invalid strategy result")
)

type ChainErrorType int

const (
	ChainErrorTypeError ChainErrorType = iota
	ChainErrorTypeIncomplete
)

type ChainError struct {
	ErrType ChainErrorType
	Message string
	Err     error
}

func NewChainError(errType ChainErrorType, message string, err error) ChainError {
	return ChainError{errType, message, err}
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
