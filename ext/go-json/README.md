# ext/go-json

JSON serialization codec for `crema` using `goccy/go-json`.

## Features

- `JSONByteStringSerializationCodec` for encoding/decoding cache objects via goccy/go-json

## Usage

```go
codec := gojson.JSONByteStringSerializationCodec[MyValue]{}
```
