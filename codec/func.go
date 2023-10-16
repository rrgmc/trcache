package codec

import (
	"context"

	"github.com/rrgmc/trcache"
)

// FuncCodec is a Codec that uses the passed functions as its methods.
type FuncCodec[V any] struct {
	m func(ctx context.Context, data V) (any, error)
	u func(ctx context.Context, data any) (V, error)
}

func NewFuncCodec[V any](marshal func(ctx context.Context, data V) (any, error),
	unmarshal func(ctx context.Context, data any) (V, error)) trcache.Codec[V] {
	return FuncCodec[V]{
		m: marshal,
		u: unmarshal,
	}
}

func (c FuncCodec[V]) Encode(ctx context.Context, data V) (any, error) {
	return c.m(ctx, data)
}

func (c FuncCodec[V]) Decode(ctx context.Context, data any) (V, error) {
	return c.u(ctx, data)
}
