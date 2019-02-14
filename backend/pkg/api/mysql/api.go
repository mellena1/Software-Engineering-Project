package mysql

import (
	"database/sql"
	"net/http"

	mysqlDriver "github.com/go-sql-driver/mysql" // mysql driver for database/sql
	"github.com/gorilla/mux"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	sqlEntities "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/sql"
)

type MySQLApi struct {
	router *mux.Router
	db     *sql.DB
}

func NewMySQLApi(mySQLConfig mysqlDriver.Config) (*MySQLApi, error) {
	db, err := sql.Open("mysql", mySQLConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	apiObj := MySQLApi{
		router: mux.NewRouter(),
		db:     db,
	}

	api.CreateRoomRoutes(apiObj, sqlEntities.NewRoomSQL(apiObj.db))
	api.CreateSessionRoutes(apiObj, sqlEntities.NewSessionSQL(apiObj.db))
	api.CreateSpeakerRoutes(apiObj, sqlEntities.NewSpeakerSQL(apiObj.db))

	return &apiObj, nil
}

func (a MySQLApi) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a.router)
}

func (a MySQLApi) CreateRoute(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc)
}

func (a MySQLApi) CreateRouteWithMethods(path string, handlerFunc func(w http.ResponseWriter, r *http.Request), methods ...string) {
	a.router.HandleFunc(path, handlerFunc).Methods(methods...)
}

func (a MySQLApi) CreatePrefixedRoute(path string, handler http.Handler) {
	a.router.PathPrefix(path).Handler(handler)
}

func (a MySQLApi) CreatePrefixedRouteWithMethods(path string, handler http.Handler, methods ...string) {
	a.router.PathPrefix(path).Handler(handler).Methods(methods...)
}

func (a MySQLApi) Close() error {
	return a.db.Close()
}
