package api

import "errors"

var (
	ErrQueryNotSet    = errors.New("a query in the HTTP request is not set")
	ErrBadQuery       = errors.New("a query in the HTTP request is invalid")
	ErrBadQueryType   = errors.New("a query is not the right type")
	ErrInvalidRequest = errors.New("the body request is invalid")
	ErrInvalidEmail   = errors.New("the email is not valid email syntax")
	ErrInvalidName    = errors.New("the name is not valid syntax")
)
