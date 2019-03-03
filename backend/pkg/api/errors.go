package api

import "errors"

var (
	// ErrQueryNotSet a query in the HTTP request is not set
	ErrQueryNotSet = errors.New("a query in the HTTP request is not set")
	// ErrBadQuery a query in the HTTP request is invalid
	ErrBadQuery = errors.New("a query in the HTTP request is invalid")
	// ErrBadQueryType a query passed is not the right type
	ErrBadQueryType = errors.New("a query is not the right type")
)
