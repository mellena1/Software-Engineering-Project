package api

import "net/http"

// API implements everything needed to run an API
type API interface {
	CreateRoute(path string, handlerFunc func(w http.ResponseWriter, r *http.Request))
	CreateRouteWithMethods(path string, handlerFunc func(w http.ResponseWriter, r *http.Request), methods ...string)
	CreatePrefixedRoute(path string, handler http.Handler)
	CreatePrefixedRouteWithMethods(path string, handler http.Handler, methods ...string)
	ListenAndServe(addr string) error
	Close() error
}
