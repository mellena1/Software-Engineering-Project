package mysql

import "errors"

// ErrDBNotSet signifies the db being nil
var ErrDBNotSet = errors.New("the database is not set")

// ErrNoRowsFound signifies that a query had no resultant rows
var ErrNoRowsFound = errors.New("the query came up empty")

// ErrTooManyRows signifies the the query returned more results than expected
var ErrTooManyRows = errors.New("the query returned more results than expected")
