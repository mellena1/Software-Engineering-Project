package api

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler func(writer http.ResponseWriter, request *http.Request)
	Methods []string
}

type PrefixedRoute struct {
	Path    string
	Handler http.Handler
	Methods []string
}

func NewRoute(path string, handler func(writer http.ResponseWriter, request *http.Request), methods ...string) Route {
	return Route{Path: path, Handler: handler, Methods: methods}
}

func NewPrefixedRoute(path string, handler http.Handler, methods ...string) PrefixedRoute {
	return PrefixedRoute{Path: path, Handler: handler, Methods: methods}
}
