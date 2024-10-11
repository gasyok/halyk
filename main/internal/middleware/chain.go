package middleware

import "net/http"

type middleware = func(http.Handler) http.Handler

// Рапинг мидлеварей по принципу ниже
// Chain returns the chained middleware composed of the given middleware sequence.
//
// The first middleware will be the outer most,
// the last middleware will be the inner most wrapper around the real call.
func Chain(middlewares ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
