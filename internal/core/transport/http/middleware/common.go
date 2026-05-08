package http_core_middleware

import (
	"context"
	"net/http"

	core_logger "github.com/Andrew1996-la/MindTick/internal/core/logger"
	core_http_response "github.com/Andrew1996-la/MindTick/internal/core/transport/http/response"
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

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := core_logger.FromLogger(ctx)

			httpHandlerResponse := core_http_response.NewHTTPHandlerResponse(logger, w)
			
			defer func() {
				if r := recover(); r != nil {
					httpHandlerResponse.PanicResponse(
						r,
						"during handle HTTP Request got expected panic",
					)
				}
			}()

			next.ServeHTTP(w, r)  
		})
	}
} 
