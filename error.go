package trcache

import "errors"

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
