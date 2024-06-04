package tracex

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Spanx struct {
	s trace.Span
}

func (this Spanx) End(options ...trace.SpanEndOption) {
	if traceIsRun.Load() {
		this.s.End()
	}
}
func (this Spanx) SetAttributes(kv ...attribute.KeyValue) {
	if traceIsRun.Load() {
		this.s.SetAttributes(kv...)
	}
}

func (this Spanx) AddEvent(name string, options ...trace.EventOption) {
	if traceIsRun.Load() {
		this.s.AddEvent(name, options...)
	}
}

func (this Spanx) SetStatus(code codes.Code, description string) {
	if traceIsRun.Load() {
		this.s.SetStatus(code, description)
	}
}

func (this Spanx) RecordError(err error, options ...trace.EventOption) {
	if traceIsRun.Load() {
		this.s.RecordError(err, options...)
	}
}
