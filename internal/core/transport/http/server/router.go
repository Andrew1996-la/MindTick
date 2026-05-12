package core_http_server

import (
	"fmt"
	"net/http"
)

type APIVersion string

var (
	APIVersion1 = APIVersion("v1")
	APIVersion2 = APIVersion("v2")
	APIVersion3 = APIVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	APIVersion APIVersion
}

func NewAPIVersionRouter(apiVersion APIVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		APIVersion: apiVersion,
	}
}

func (a *APIVersionRouter) Register(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.method, route.path)

		a.Handle(pattern, route.handler)
	}
}
