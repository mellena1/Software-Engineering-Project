package api

// API implements everything needed to run an API
type API interface {
	CreateRoutes(route ...Route)
	CreatePrefixedRoutes(route ...PrefixedRoute)
	ListenAndServe(addr string) error
	Close() error
}
