package trcache

import (
	"context"
	"encoding/json"
	"errors"
)

// ForwardCodec is a Codec that returns the same object passed.
type ForwardCodec[V any] struct {
}

func NewForwardCodec[V any]() Codec[V] {
	return ForwardCodec[V]{}
}

func (c ForwardCodec[V]) Marshal(ctx context.Context, data V) (any, error) {
	return data, nil
}

func (c ForwardCodec[V]) Unmarshal(ctx context.Context, data any) (V, error) {
	switch dt := data.(type) {
	case V:
		return dt, nil
	}
	var empty V
	return empty, errors.New("cannot unmarshall value")
}

// JSONCodec is a Codec that marshals from/to JSON.
type JSONCodec[V any] struct {
}

func NewJSONCodec[V any]() Codec[V] {
	return JSONCodec[V]{}
}

func (c JSONCodec[V]) Marshal(ctx context.Context, data V) (any, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c JSONCodec[V]) Unmarshal(ctx context.Context, data any) (V, error) {
	var ret V
	var udata []byte

	switch dt := data.(type) {
	case []byte:
		udata = dt
	case string:
		udata = []byte(dt)
	default:
		return ret, errors.New("unknown data type for JSON unmarshal")
	}

	if err := json.Unmarshal(udata, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}
