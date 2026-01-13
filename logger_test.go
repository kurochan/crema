package crema

import (
	"context"
	"log/slog"
	"testing"
)

func TestNoopLogHandler_Basics(t *testing.T) {
	t.Parallel()

	handler := noopLogHandler{}
	if handler.Enabled(context.Background(), slog.LevelInfo) {
		t.Fatal("expected handler to report disabled")
	}

	if err := handler.Handle(context.Background(), slog.Record{}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if handler.WithAttrs([]slog.Attr{}) != handler {
		t.Fatal("expected WithAttrs to return handler")
	}
	if handler.WithGroup("group") != handler {
		t.Fatal("expected WithGroup to return handler")
	}
}
