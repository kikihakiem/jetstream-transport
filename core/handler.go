package gkit

import (
	"context"
	"log/slog"
)

// FinalizerFunc can be used to perform work at the end of a request.
type FinalizerFunc[Req, Res any] func(ctx context.Context, request Req, response Res, err error)

// ErrorHandler receives a transport error to be processed for diagnostic purposes.
// Usually this means logging the error.
type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

// The ErrorHandlerFunc type is an adapter to allow the use of
// ordinary function as ErrorHandler. If f is a function
// with the appropriate signature, ErrorHandlerFunc(f) is a
// ErrorHandler that calls f.
type ErrorHandlerFunc func(ctx context.Context, err error)

// Handle calls f(ctx, err).
func (f ErrorHandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}

type LogFunc func(context.Context, error)

var LogErrorHandler = func(logFunc LogFunc) ErrorHandlerFunc {
	return func(ctx context.Context, err error) {
		if logFunc == nil {
			slog.ErrorContext(ctx, err.Error())
		} else {
			logFunc(ctx, err)
		}
	}
}