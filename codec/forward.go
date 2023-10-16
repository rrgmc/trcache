package codec

import (
	"context"
	"fmt"

	"github.com/rrgmc/trcache"
)

// ForwardCodec is a Codec that returns the same object passed.
type ForwardCodec[V any] struct {
}

func NewForwardCodec[V any]() trcache.Codec[V] {
	return ForwardCodec[V]{}
}

func (c ForwardCodec[V]) Encode(ctx context.Context, data V) (any, error) {
	return data, nil
}

func (c ForwardCodec[V]) Decode(ctx context.Context, data any) (V, error) {
	switch dt := data.(type) {
	case V:
		return dt, nil
	}
	var empty V
	return empty, &trcache.InvalidValueTypeError{fmt.Sprintf("cannot unmarshall value of type '%T' to type '%T'",
		data, empty)}
}
