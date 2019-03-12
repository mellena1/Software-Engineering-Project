package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// ReportError logs the error and respond with message
func ReportError(err error, msg string, httpStatusCode int, w http.ResponseWriter) {
	w.WriteHeader(httpStatusCode)
	jsonMsg, _ := json.Marshal(msg)
	w.Write(jsonMsg)
	log.Printf("%s, %v\n", msg, err)
}

// getIDFromQueries given an http.Request, return back the id key as an int64
func getIDFromQueries(r *http.Request) (int64, error) {
	queries := r.URL.Query()
	requestedIDArray, ok := queries["id"]
	if !ok || len(requestedIDArray) < 1 {
		return 0, ErrQueryNotSet
	} else if len(requestedIDArray) > 1 {
		return 0, ErrBadQuery
	}
	requestedID, err := strconv.Atoi(requestedIDArray[0])
	if err != nil {
		return 0, ErrBadQueryType
	}
	return int64(requestedID), nil
}

// writeIDToClient writes the id back to the client as JSON
func writeIDToClient(w http.ResponseWriter, id int64) {
	response := map[string]int64{"id": id}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}
