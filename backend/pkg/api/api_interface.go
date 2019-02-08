package api

type API interface {
	ListenAndServe(addr string) error
	Close() error
}
