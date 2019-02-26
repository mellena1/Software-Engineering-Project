package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// ReportError logs the error and respond with message
func ReportError(err error, msg string, httpStatusCode int, w http.ResponseWriter) {
	w.WriteHeader(httpStatusCode)
	jsonMsg, _ := json.Marshal(msg)
	w.Write([]byte(jsonMsg))
	log.Printf("%s, %v\n", msg, err)
}
