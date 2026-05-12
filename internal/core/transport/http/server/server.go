package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Andrew1996-la/MindTick/internal/core/logger"
	"go.uber.org/zap"
)

type HttpServer struct {
	mux    *http.ServeMux
	config Config
	logger core_logger.Logger
}

func NewHttpServer(
	config Config,
	logger core_logger.Logger,
) *HttpServer {
	return &HttpServer{
		mux:    http.NewServeMux(),
		config: config,
		logger: logger,
	}
}

func (h *HttpServer) RegisterAPIRoutes(routes ...APIVersionRouter) {
	for _, router := range routes {
		prefix := "/api/" + string(router.APIVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}

func (h *HttpServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		h.logger.Warn("http server started", zap.String("Addr", h.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http: %w", err)
		}
	case <-ctx.Done():
		h.logger.Warn("shutdown server...")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server: %w", err)
		}

		h.logger.Warn("HTTP server stopped")
	}

	return nil
}
