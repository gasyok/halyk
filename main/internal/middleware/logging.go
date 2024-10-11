package middleware

import (
	"net/http"

	"github.com/halyk/main/internal/pkg/logger"
	"go.uber.org/zap"
)

type statusCodeCatcher struct {
	http.ResponseWriter
	responded int
}

func (scc *statusCodeCatcher) WriteHeader(statusCode int) {
	scc.ResponseWriter.WriteHeader(statusCode)
	scc.responded = statusCode
}

func WithLogging() middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scc := &statusCodeCatcher{ResponseWriter: w, responded: http.StatusOK}
			next.ServeHTTP(scc, r)
			logger.Info("Request served",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("code", scc.responded),
			)
		})
	}
}
