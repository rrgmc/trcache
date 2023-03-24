package codec

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/RangelReale/trcache"
)

// JSONCodec is a Codec that marshals from/to JSON.
type JSONCodec[V any] struct {
	jsonCodecOptions
}

type jsonCodecOptions struct {
	returnBytes bool
}

func NewJSONCodec[V any](options ...JSONCodecOption) trcache.Codec[V] {
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
	if !c.returnBytes {
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
		return ret, &trcache.ErrInvalidValueType{
			fmt.Sprintf("unknown data type '%s' for JSON unmarshal", getType(data)),
		}
	}

	if err := json.Unmarshal(udata, &ret); err != nil {
		return ret, err
	}
	return ret, nil
}

type JSONCodecOption func(*jsonCodecOptions)

func WithJSONCodecReturnBytes(returnBytes bool) JSONCodecOption {
	return func(o *jsonCodecOptions) {
		o.returnBytes = returnBytes
	}
}
