package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver for database/sql
)

func NewMySQL(dataSource string) (*sql.DB, error) {
	return sql.Open("mysql", dataSource)
}
