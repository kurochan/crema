package gojson

import (
	"bytes"

	json "github.com/goccy/go-json"
	"github.com/abema/crema"
)

// JSONByteStringSerializationCodec marshals cache objects as JSON bytes via goccy/go-json.
type JSONByteStringSerializationCodec[V any] struct{}

var _ crema.SerializationCodec[any, []byte] = JSONByteStringSerializationCodec[any]{}

// Encode marshals the cache object into JSON bytes without a trailing newline.
func (j JSONByteStringSerializationCodec[V]) Encode(value crema.CacheObject[V]) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}
	b := buf.Bytes()
	if len(b) > 0 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	return b, nil
}

// Decode unmarshals JSON bytes into a cache object.
func (j JSONByteStringSerializationCodec[V]) Decode(data []byte) (crema.CacheObject[V], error) {
	var out crema.CacheObject[V]
	if err := json.Unmarshal(data, &out); err != nil {
		return crema.CacheObject[V]{}, err
	}
	return out, nil
}
