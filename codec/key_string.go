package codec

import (
	"context"

	"github.com/RangelReale/trcache"
)

// StringKeyCodec is a key codec that converts all values to string.
type StringKeyCodec[K comparable] struct {
}

func NewStringKeyCodec[K comparable]() *StringKeyCodec[K] {
	return &StringKeyCodec[K]{}
}

func (s *StringKeyCodec[K]) Convert(ctx context.Context, key K) (any, error) {
	return trcache.StringValue(key), nil
}
