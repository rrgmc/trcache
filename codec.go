package trcache

import (
	"context"
	"errors"
)

// ForwardCodec is a Codec that returns the same object passed.
type ForwardCodec[T any] struct {
}

func NewCacheForwardMarshaller[T any]() Codec[T] {
	return ForwardCodec[T]{}
}

func (c ForwardCodec[T]) Marshal(ctx context.Context, data T) (any, error) {
	return data, nil
}

func (c ForwardCodec[T]) Unmarshal(ctx context.Context, data any) (T, error) {
	switch dt := data.(type) {
	case T:
		return dt, nil
	}
	var empty T
	return empty, errors.New("cannot unmarshall value")
}
