package mysql

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver for database/sql
	"github.com/gorilla/mux"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	sqlEntities "github.com/mellena1/Software-Engineering-Project/backend/pkg/db/sql"
)

type MySQLApi struct {
	Router *mux.Router
	db     *sql.DB
}

func NewMySQLApi(dataSource string) (*MySQLApi, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	mysql := MySQLApi{
		Router: mux.NewRouter(),
		db:     db,
	}

	api.AddAllRoomRoutesToRouter(mysql.Router, sqlEntities.NewRoomSQL(mysql.db))
	api.AddAllSessionRoutesToRouter(mysql.Router, sqlEntities.NewSessionSQL(mysql.db))
	api.AddAllSpeakerRoutesToRouter(mysql.Router, sqlEntities.NewSpeakerSQL(mysql.db))

	return &mysql, nil
}

func (a MySQLApi) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a.Router)
}

func (a *MySQLApi) Close() error {
	return a.db.Close()
}
