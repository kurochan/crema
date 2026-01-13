//go:generate env GOWORK=off go -C tool tool buf generate ../ --template ../buf.gen.yaml
//go:generate env GOWORK=off go -C tool tool buf lint ../ --config ../buf.yaml
package protobuf
