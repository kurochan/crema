package crema

import (
	"context"
	"time"
)

// CacheProvider abstracts storage for encoded cache entries.
// Implementations must be safe for concurrent use by multiple goroutines.
type CacheProvider[S any] interface {
	// Get retrieves a value from the cache by key.
	Get(ctx context.Context, key string) (S, bool, error)
	// Set stores a value in the cache with the specified key.
	Set(ctx context.Context, key string, value S, ttl time.Duration) error
	// Delete removes a value from the cache by key.
	Delete(ctx context.Context, key string) error
}
