package golanglru

import (
	"context"
	"testing"
	"time"
)

func TestCacheProvider_GetSetDelete(t *testing.T) {
	t.Parallel()

	provider := NewCacheProvider[string](2, time.Second)
	ctx := context.Background()

	if err := provider.Set(ctx, "key", "value", 0); err != nil {
		t.Fatalf("set: %v", err)
	}

	value, ok, err := provider.Get(ctx, "key")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if !ok {
		t.Fatal("expected value to exist")
	}
	if value != "value" {
		t.Fatalf("unexpected value: %q", value)
	}

	if err := provider.Delete(ctx, "key"); err != nil {
		t.Fatalf("delete: %v", err)
	}

	value, ok, err = provider.Get(ctx, "key")
	if err != nil {
		t.Fatalf("get after delete: %v", err)
	}
	if ok {
		t.Fatal("expected value to be deleted")
	}
	if value != "" {
		t.Fatalf("expected zero value after delete, got %q", value)
	}
}

func TestCacheProvider_DefaultTTL(t *testing.T) {
	t.Parallel()

	provider := NewCacheProvider[string](2, 30*time.Millisecond)
	ctx := context.Background()

	if err := provider.Set(ctx, "key", "value", time.Second); err != nil {
		t.Fatalf("set: %v", err)
	}

	if _, ok, err := provider.Get(ctx, "key"); err != nil {
		t.Fatalf("get: %v", err)
	} else if !ok {
		t.Fatal("expected value to exist")
	}

	deadline := time.Now().Add(200 * time.Millisecond)
	for time.Now().Before(deadline) {
		if _, ok, err := provider.Get(ctx, "key"); err != nil {
			t.Fatalf("get after ttl: %v", err)
		} else if !ok {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}

	t.Fatal("expected value to expire")
}
