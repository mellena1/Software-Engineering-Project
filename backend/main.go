package main

import (
	"log"
	"net/http"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/api"

	"github.com/gorilla/mux"
	_ "github.com/mellena1/Software-Engineering-Project/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Code Camp Counter API
// @version 1.0
// @description The API for the code camp counting program.
func main() {
	router := mux.NewRouter()
	addSpeakerRoutes(router)
	addRoomRoutes(router)
	addSessionRoutes(router)
	router.PathPrefix("/api/v1/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func addSpeakerRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/speaker", api.GetAllSpeakers).Methods("GET")
}

func addRoomRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/room", api.GetAllRooms).Methods("GET")
}

func addSessionRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/session", api.GetAllSessions).Methods("GET")
}
