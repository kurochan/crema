//go:generate go generate ./ext/protobuf
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./... --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/go-json --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/golang-lru --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/gomemcache --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/protobuf --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/ristretto --fix
//go:generate go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0 run ./ext/valkey-go --fix
package crema
