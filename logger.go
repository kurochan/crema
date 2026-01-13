package crema

import (
	"context"
	"log/slog"
)

type noopLogHandler struct{}

var _ slog.Handler = noopLogHandler{}

func (noopLogHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (noopLogHandler) Handle(context.Context, slog.Record) error { return nil }
func (h noopLogHandler) WithAttrs([]slog.Attr) slog.Handler      { return h }
func (h noopLogHandler) WithGroup(string) slog.Handler           { return h }
