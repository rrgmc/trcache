package codec

import (
	"context"

	"github.com/RangelReale/trcache"
)

// FuncCodec is a Codec that returns the same object passed.
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

func (c FuncCodec[V]) Marshal(ctx context.Context, data V) (any, error) {
	return c.m(ctx, data)
}

func (c FuncCodec[V]) Unmarshal(ctx context.Context, data any) (V, error) {
	return c.u(ctx, data)
}
