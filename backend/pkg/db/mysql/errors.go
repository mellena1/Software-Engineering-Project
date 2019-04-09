package mysql

import "errors"

var ErrDBNotSet = errors.New("the database is not set")
var ErrNoRowsFound = errors.New("the query came up empty")
var ErrTooManyRows = errors.New("the query returned more results than expected")
