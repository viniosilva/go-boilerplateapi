package logger

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func Info(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, setTrace(ctx)...)

	slog.InfoContext(ctx, msg, attrsToArgs(args)...)
}

func Debug(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, setTrace(ctx)...)

	slog.DebugContext(ctx, msg, attrsToArgs(args)...)
}

func Warn(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, setTrace(ctx)...)

	slog.WarnContext(ctx, msg, attrsToArgs(args)...)
}

func Error(ctx context.Context, msg string, args ...slog.Attr) {
	args = append(args, setTrace(ctx)...)

	slog.ErrorContext(ctx, msg, attrsToArgs(args)...)
}

func setTrace(ctx context.Context) []slog.Attr {
	span := trace.SpanFromContext(ctx)
	if !span.SpanContext().IsValid() {
		return nil
	}

	spanCtx := trace.SpanFromContext(ctx).SpanContext()

	return []slog.Attr{
		slog.String("trace_id", spanCtx.TraceID().String()),
		slog.String("span_id", spanCtx.SpanID().String()),
	}
}

func attrsToArgs(attrs []slog.Attr) []any {
	args := make([]any, len(attrs))
	for i, a := range attrs {
		args[i] = a
	}

	return args
}
