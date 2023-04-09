package codec

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/RangelReale/trcache"
)

// GOBCodec is a Codec that marshals from/to [encoding/gob].
type GOBCodec[V any] struct {
	gobCodecOptions
}

type gobCodecOptions struct {
	returnString bool
}

func NewGOBCodec[V any](options ...GOBCodecOption) trcache.Codec[V] {
	ret := GOBCodec[V]{}
	for _, opt := range options {
		opt(&ret.gobCodecOptions)
	}
	return ret
}

func (c GOBCodec[V]) Encode(ctx context.Context, data V) (any, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	if c.returnString {
		return string(buf.Bytes()), nil
	}
	return buf.Bytes(), nil
}

func (c GOBCodec[V]) Decode(ctx context.Context, data any) (V, error) {
	var ret V
	var udata []byte

	switch dt := data.(type) {
	case []byte:
		udata = dt
	case string:
		udata = []byte(dt)
	default:
		return ret, &trcache.InvalidValueTypeError{
			fmt.Sprintf("unknown data type '%T' for GOB unmarshal", data),
		}
	}

	buf := bytes.NewBuffer(udata)
	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&ret); err != nil {
		return ret, err
	}
	return ret, nil
}

type GOBCodecOption func(*gobCodecOptions)

func WithGOBCodecReturnString(returnString bool) GOBCodecOption {
	return func(o *gobCodecOptions) {
		o.returnString = returnString
	}
}
