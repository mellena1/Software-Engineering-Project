package main

import (
	"fmt"
	"log"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api/mysql"
	"github.com/mellena1/Software-Engineering-Project/backend/pkg/config"

	_ "github.com/mellena1/Software-Engineering-Project/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Code Camp Counter API
// @version 1.0
// @description The API for the code camp counting program.
func main() {
	var app api.API
	app, err := mysql.NewAPI(*config.Values.MySQLConfig, config.Values.LogWriter)
	if err != nil {
		panic(err)
	}
	defer app.Close()

	if config.Values.RunSwagger {
		swaggerRoute := api.NewPrefixedRoute("/api/v1/swagger/", httpSwagger.WrapHandler)
		app.CreatePrefixedRoutes(swaggerRoute)
	}

	address := fmt.Sprintf(":%d", config.Values.Port)
	log.Printf("Starting the server at %s...", address)
	log.Fatal(app.ListenAndServe(address))
}
