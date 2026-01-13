# ext/gomemcache

Memcached cache provider for `crema` using `gomemcache`.

## Features

- `MemcacheCacheProvider` for storing cache data in Memcached with TTL handling

## Usage

```go
client := memcache.New("127.0.0.1:11211")
provider := gomemcache.NewMemcacheCacheProvider(client)
```
