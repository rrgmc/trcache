package trcache

import "context"

type Codec[V any] interface {
	Marshal(ctx context.Context, data V) (any, error)
	Unmarshal(ctx context.Context, data any) (V, error)
}

type Validator[V any] interface {
	ValidateGet(ctx context.Context, data V) error
}
