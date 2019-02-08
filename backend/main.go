package main

import (
	"log"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api/mysql"

	_ "github.com/mellena1/Software-Engineering-Project/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

var app api.API

// @title Code Camp Counter API
// @version 1.0
// @description The API for the code camp counting program.
func main() {
	app, err := mysql.NewMySQLApi("")
	if err != nil {
		panic(err)
	}
	defer app.Close()

	app.Router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting the server...")
	log.Fatal(app.ListenAndServe(":8081"))
}
