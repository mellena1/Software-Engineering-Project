package mysql

import (
	"database/sql"
	"io"
	"net/http"

	mysqlDriver "github.com/go-sql-driver/mysql" // mysql driver for database/sql
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	mysqlEntities "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/mysql"
)

// API implements a router using gorilla/mux and a db using the go-sql-driver/mysql lib
type API struct {
	router    *mux.Router
	logWriter io.Writer
	db        *sql.DB
}

// NewAPI returns a new API given a mysqlDriver.Config
// passing nil to logWriter means no logging
func NewAPI(mySQLConfig mysqlDriver.Config, logWriter io.Writer) (*API, error) {
	db, err := sql.Open("mysql", mySQLConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	apiObj := API{
		router:    mux.NewRouter(),
		db:        db,
		logWriter: logWriter,
	}

	// Room
	roomDBFacade := mysqlEntities.NewRoomMySQL(apiObj.db)
	roomRoutes := api.CreateRoomRoutes(roomDBFacade)
	apiObj.CreateRoutes(roomRoutes...)

	// Session
	sessionDBFacade := mysqlEntities.NewSessionMySQL(apiObj.db)
	sessionRoutes := api.CreateSessionRoutes(sessionDBFacade)
	apiObj.CreateRoutes(sessionRoutes...)

	// Speaker
	speakerDBFacade := mysqlEntities.NewSpeakerMySQL(apiObj.db)
	speakerRoutes := api.CreateSpeakerRoutes(speakerDBFacade)
	apiObj.CreateRoutes(speakerRoutes...)

	// Timeslot
	timeslotDBFacade := mysqlEntities.NewTimeslotMySQL(apiObj.db)
	timeslotRoutes := api.CreateTimeslotRoutes(timeslotDBFacade)
	apiObj.CreateRoutes(timeslotRoutes...)

	return &apiObj, nil
}

// ListenAndServe basically runs http.ListenAndServe
func (a API) ListenAndServe(addr string) error {
	// Add logging if there is a logWriter defined
	var handler http.Handler = a.router
	if a.logWriter != nil {
		handler = handlers.LoggingHandler(a.logWriter, a.router)
	}

	// Add CORS stuff
	handler = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(handler)

	return http.ListenAndServe(addr, handler)
}

// CreateRoutes adds a route(s) to the router
func (a API) CreateRoutes(route ...api.Route) {
	for _, r := range route {
		if len(r.Methods) > 0 {
			a.router.HandleFunc(r.Path, r.Handler).Methods(r.Methods...)
		} else {
			a.router.HandleFunc(r.Path, r.Handler)
		}
	}
}

// CreatePrefixedRoutes adds route(s) for a prefix to the router
func (a API) CreatePrefixedRoutes(route ...api.PrefixedRoute) {
	for _, r := range route {
		if len(r.Methods) > 0 {
			a.router.PathPrefix(r.Path).Handler(r.Handler).Methods(r.Methods...)
		} else {
			a.router.PathPrefix(r.Path).Handler(r.Handler)
		}
	}
}

// Close closes the API's db
func (a API) Close() error {
	return a.db.Close()
}
