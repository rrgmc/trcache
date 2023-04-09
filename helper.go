package trcache

import "context"

// KeyCodec is a codec to convert a cache key to a format supported by the cache.
type KeyCodec[K comparable] interface {
	Convert(ctx context.Context, key K) (any, error)
}

// Codec is a codec to convert a cache value the requested type.
type Codec[V any] interface {
	Encode(ctx context.Context, data V) (any, error)
	Decode(ctx context.Context, data any) (V, error)
}

// Validator allows validating data retrieved from the cache, allowing to check for a previous format for example.
type Validator[V any] interface {
	ValidateGet(ctx context.Context, data V) error
}
