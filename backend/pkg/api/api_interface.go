package api

type API interface {
	CreateRoutes(route ...Route)
	CreatePrefixedRoutes(route ...PrefixedRoute)
	ListenAndServe(address string) error
	Close() error
}
