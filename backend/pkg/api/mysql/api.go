package mysql

import (
	"database/sql"
	"net/http"

	mysqlDriver "github.com/go-sql-driver/mysql" // mysql driver for database/sql
	"github.com/gorilla/mux"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	mysqlEntities "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/mysql"
)

// API implements a router using gorilla/mux and a db using the go-sql-driver/mysql lib
type API struct {
	router *mux.Router
	db     *sql.DB
}

// NewAPI returns a new API given a mysqlDriver.Config
func NewAPI(mySQLConfig mysqlDriver.Config) (*API, error) {
	db, err := sql.Open("mysql", mySQLConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	apiObj := API{
		router: mux.NewRouter(),
		db:     db,
	}

	api.CreateRoomRoutes(apiObj, mysqlEntities.NewRoomSQL(apiObj.db))
	api.CreateSessionRoutes(apiObj, mysqlEntities.NewSessionSQL(apiObj.db))
	api.CreateSpeakerRoutes(apiObj, mysqlEntities.NewSpeakerSQL(apiObj.db))

	return &apiObj, nil
}

// ListenAndServe basically runs http.ListenAndServe
func (a API) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a.router)
}

// CreateRoute adds a route to the router
func (a API) CreateRoute(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc)
}

// CreateRouteWithMethods adds a route to the router for specific http request methods (GET, POST, etc)
func (a API) CreateRouteWithMethods(path string, handlerFunc func(w http.ResponseWriter, r *http.Request), methods ...string) {
	a.router.HandleFunc(path, handlerFunc).Methods(methods...)
}

// CreatePrefixedRoute adds a route for a prefix to the router
func (a API) CreatePrefixedRoute(path string, handler http.Handler) {
	a.router.PathPrefix(path).Handler(handler)
}

// CreatePrefixedRouteWithMethods adds a route for a prefix to the router for specific http request methods (GET, POST, etc)
func (a API) CreatePrefixedRouteWithMethods(path string, handler http.Handler, methods ...string) {
	a.router.PathPrefix(path).Handler(handler).Methods(methods...)
}

// Close closes the API's db
func (a API) Close() error {
	return a.db.Close()
}
