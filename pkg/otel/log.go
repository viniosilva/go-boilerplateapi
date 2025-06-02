package otel

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func OtelLog(ctx context.Context) []any {
	span := trace.SpanFromContext(ctx).SpanContext()

	return []any{
		"trace_id", span.TraceID().String(),
		"span_id", span.SpanID().String(),
	}
}
