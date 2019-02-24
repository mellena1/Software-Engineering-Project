package api

import (
	"net/http"
	"encoding/json"
)

type Helper interface {
}

type BodyData struct{
	Email string
}

func getParamsFromRequest(req *http.Request) string {
	reader := json.NewDecoder(req.Body)

	var data BodyData
	err := reader.Decode(&data)
	if err != nil {
		return("")
	}

	email := data.Email
	return email;
}
