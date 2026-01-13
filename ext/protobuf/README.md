# ext/protobuf

Protobuf serialization codec for `crema`.

## Features

- `ProtoCacheCodec` for encoding/decoding cache objects via protobuf
- `ProtoCacheObject` envelope message

## Usage

```go
codec, err := NewProtoCacheCodec(&mypb.MyMessage{})
if err != nil {
    // handle error
}
```

## Generate protobuf code

```sh
go generate ./ext/protobuf
```
