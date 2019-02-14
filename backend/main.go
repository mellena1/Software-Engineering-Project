package main

import (
	"log"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api/mysql"

	mysqlDriver "github.com/go-sql-driver/mysql" // mysql driver for database/sql

	_ "github.com/mellena1/Software-Engineering-Project/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Code Camp Counter API
// @version 1.0
// @description The API for the code camp counting program.
func main() {
	var app api.API
	app, err := mysql.NewAPI(mysqlDriver.Config{
		User:      "",
		Passwd:    "",
		Addr:      "",
		DBName:    "",
		ParseTime: true,
	})
	if err != nil {
		panic(err)
	}
	defer app.Close()

	swaggerRoute := api.NewPrefixedRoute("/api/v1/swagger/", httpSwagger.WrapHandler)
	app.CreatePrefixedRoutes(swaggerRoute)

	log.Println("Starting the server...")
	log.Fatal(app.ListenAndServe(":8081"))
}
