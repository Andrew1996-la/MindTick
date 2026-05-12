package core_http_server

import "net/http"

type Route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) Route {
	return Route{
		method:  method,
		path:    path,
		handler: handler,
	}
} 
