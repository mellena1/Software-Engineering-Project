package api

import (
	"net/http"
	"strings"
)

type Helper interface {
}

func getParamsFromRequest(req *http.Request) string {
	uri := req.RequestURI

	if strings.Contains(uri, "/speaker/") {
		return strings.TrimPrefix(uri, "/api/v1/speaker/")
	} else if strings.Contains(uri, "/room/") {
		return strings.TrimPrefix(uri, "/api/v1/room/")
	} else if strings.Contains(uri, "/session/") {
		return strings.TrimPrefix(uri, "/api/v1/session/")
	} else {
		//this uri should not have parameters
		return ""
	}
}
