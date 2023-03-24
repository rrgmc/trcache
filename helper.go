package trcache

import "context"

type KeyCodec[K comparable] interface {
	Convert(ctx context.Context, key K) (any, error)
}

type Codec[V any] interface {
	Marshal(ctx context.Context, data V) (any, error)
	Unmarshal(ctx context.Context, data any) (V, error)
}

type Validator[V any] interface {
	ValidateGet(ctx context.Context, data V) error
}
