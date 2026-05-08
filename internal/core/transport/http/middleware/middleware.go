package http_core_middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler
