package api

import (
	"net/http"
)

// Route holds all needed fields to make a route
type Route struct {
	Path    string
	Handler func(writer http.ResponseWriter, request *http.Request)
	Methods []string
}

// PrefixedRoute holds all needed fields to make a prefixed route
type PrefixedRoute struct {
	Path    string
	Handler http.Handler
	Methods []string
}

// NewRoute returns a new Route struct. methods are zero-to-many
func NewRoute(path string, handler func(writer http.ResponseWriter, request *http.Request), methods ...string) Route {
	return Route{Path: path, Handler: handler, Methods: methods}
}

// NewPrefixedRoute returns a new PrefixedRoute struct. methods are zero-to-many
func NewPrefixedRoute(path string, handler http.Handler, methods ...string) PrefixedRoute {
	return PrefixedRoute{Path: path, Handler: handler, Methods: methods}
}
