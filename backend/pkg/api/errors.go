package api

import "errors"

var (
	// ErrQueryNotSet a query in the HTTP request is not set
	ErrQueryNotSet = errors.New("a query in the HTTP request is not set")
	// ErrBadQuery a query in the HTTP request is invalid
	ErrBadQuery = errors.New("a query in the HTTP request is invalid")
	// ErrBadQueryType a query passed is not the right type
	ErrBadQueryType = errors.New("a query is not the right type")
	// ErrInvalidRequest the body request is invalid (body requests)
	ErrInvalidRequest = errors.New("the body request is invalid")
	// ErrInvalidEmail the email is not valid email syntax
	ErrInvalidEmail = errors.New("the email is not valid email syntax")
	// ErrInvalidName the name is not valid syntax
	ErrInvalidName = errors.New("the name is not valid syntax")
)
