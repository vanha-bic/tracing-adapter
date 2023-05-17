package ext

import (
	"tracing-adapter"
	"tracing-adapter/log"
)

// LogError sets the error=true tag on the Span and logs err as an "error" event.
func LogError(span opentracing.Span, err error, fields ...log.Field) {
	Error.Set(span, true)
	ef := []log.Field{
		log.Event("error"),
		log.Error(err),
	}
	ef = append(ef, fields...)
	span.LogFields(ef...)
}
