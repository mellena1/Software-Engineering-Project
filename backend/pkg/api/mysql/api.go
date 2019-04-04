package mysql

import (
	"database/sql"
	"io"
	"net/http"

	mysqlDriver "github.com/go-sql-driver/mysql"
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
// passing nil to logWriter means no access logs
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

	// Count
	countDBFacade := mysqlEntities.NewCountMySQL(apiObj.db)
	countRoutes := api.CreateCountRoutes(countDBFacade)
	apiObj.CreateRoutes(countRoutes...)

	return &apiObj, nil
}

// ListenAndServe basically runs http.ListenAndServe
func (myAPI API) ListenAndServe(address string) error {
	return http.ListenAndServe(address, myAPI.getHandler())
}

// Gets the handler for this API
func (myAPI API) getHandler() http.Handler {
	// Add logging if there is a logWriter defined
	var handler http.Handler = myAPI.router
	if myAPI.logWriter != nil {
		handler = handlers.LoggingHandler(myAPI.logWriter, myAPI.router)
	}

	// Add CORS stuff
	handler = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(handler)

	return handler
}

// CreateRoutes adds a route(s) to the router
func (myAPI API) CreateRoutes(route ...api.Route) {
	for _, myRoute := range route {
		if len(myRoute.Methods) > 0 {
			myAPI.router.HandleFunc(myRoute.Path, myRoute.Handler).Methods(myRoute.Methods...)
		} else {
			myAPI.router.HandleFunc(myRoute.Path, myRoute.Handler)
		}
	}
}

// CreatePrefixedRoutes adds route(s) for a prefix to the router
func (myAPI API) CreatePrefixedRoutes(route ...api.PrefixedRoute) {
	for _, myRoute := range route {
		if len(myRoute.Methods) > 0 {
			myAPI.router.PathPrefix(myRoute.Path).Handler(myRoute.Handler).Methods(myRoute.Methods...)
		} else {
			myAPI.router.PathPrefix(myRoute.Path).Handler(myRoute.Handler)
		}
	}
}

// Close closes the API's db
func (myAPI API) Close() error {
	return myAPI.db.Close()
}
