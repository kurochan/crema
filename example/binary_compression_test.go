package example

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/abema/crema"
	"github.com/abema/crema/ext/protobuf"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ExampleNewBinaryCompressionCodec() {
	provider := &byteProvider{items: make(map[string][]byte)}
	protobufCodec, err := protobuf.NewProtobufCodec(&wrapperspb.BytesValue{})
	if err != nil {
		fmt.Println(err)

		return
	}
	codec := crema.NewBinaryCompressionCodec(protobufCodec, 0)

	cache := crema.NewCache(provider, codec)
	payload := bytes.Repeat([]byte("a"), 128)
	value, err := cache.GetOrLoad(context.Background(), "blob", time.Minute, func(ctx context.Context) (*wrapperspb.BytesValue, error) {
		return &wrapperspb.BytesValue{Value: payload}, nil
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	provider.mu.Lock()
	stored := append([]byte(nil), provider.items["blob"]...)
	provider.mu.Unlock()

	fmt.Println(bytes.Equal(value.Value, payload))
	fmt.Println(stored[0] == crema.CompressionTypeIDZlib)
	fmt.Println(len(stored[1:]) < len(payload))
	// Output:
	// true
	// true
	// true
}
