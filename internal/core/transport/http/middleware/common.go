package http_core_middleware

import (
	"context"
	"net/http"

	core_logger "github.com/Andrew1996-la/MindTick/internal/core/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIdHeader = "X-Request-ID"
)

func RequestId() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)

			if requestId == "" {
				requestId = uuid.NewString()
			}

			r.Header.Set(requestIdHeader, requestId)
			w.Header().Set(requestIdHeader, requestId)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIdHeader)

			l := log.With(
				zap.String("request_id", requestId),
				zap.String("URL", r.URL.String()),
			)

			ctx := context.WithValue(
				r.Context(),
				"log",
				l,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
