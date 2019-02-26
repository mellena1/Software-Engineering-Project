package api

import (
	"log"
	"net/http"
)

// ReportError logs the error and respond with message
func ReportError(err error, msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
	log.Printf("%s, %v\n", msg, err)
}
