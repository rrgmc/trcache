package trcache

import (
	"context"
	"encoding/json"
	"fmt"
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
	return empty, fmt.Errorf("cannot unmarshall value of type '%s' to type '%s'",
		getType(data), getType(empty))
}

// JSONCodec is a Codec that marshals from/to JSON.
type JSONCodec[V any] struct {
	jsonCodecOptions
}

type jsonCodecOptions struct {
	returnString bool
}

func NewJSONCodec[V any](options ...JSONCodecOption) Codec[V] {
	ret := JSONCodec[V]{}
	for _, opt := range options {
		opt(&ret.jsonCodecOptions)
	}
	return ret
}

func (c JSONCodec[V]) Marshal(ctx context.Context, data V) (any, error) {
	ret, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if c.returnString {
		return string(ret), nil
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
		return ret, fmt.Errorf("unknown data type '%s' for JSON unmarshal", getType(data))
	}

	if err := json.Unmarshal(udata, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}

type JSONCodecOption func(*jsonCodecOptions)

func WithJSONCodecReturnString(returnString bool) JSONCodecOption {
	return func(o *jsonCodecOptions) {
		o.returnString = returnString
	}
}
