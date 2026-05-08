package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/Andrew1996-la/MindTick/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPHandlerResponse struct {
	logger *core_logger.Logger
	rw     http.ResponseWriter
}

func NewHTTPHandlerResponse(logger *core_logger.Logger, rw http.ResponseWriter) *HTTPHandlerResponse {
	return &HTTPHandlerResponse{
		logger: logger,
		rw:     rw,
	}
}

func (h *HTTPHandlerResponse) PanicResponse(panic any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", panic)

	h.logger.Error(msg, zap.Error(err))

	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"msg": msg,
		"err": err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.logger.Error("write HTTP response: %w", zap.Error(err))
	}
}
