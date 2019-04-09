package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func ReportError(err error, message string, httpStatusCode int, writer http.ResponseWriter) {
	writer.WriteHeader(httpStatusCode)
	jsonMessage, _ := json.Marshal(message)
	writer.Write(jsonMessage)
	log.Printf("%s, %v\n", message, err)
}

func getIDFromQueries(request *http.Request) (int64, error) {
	queries := request.URL.Query()
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

func writeIDToClient(writer http.ResponseWriter, id int64) {
	response := map[string]int64{"id": id}
	responseJSON, _ := json.Marshal(response)
	writer.Write(responseJSON)
}
