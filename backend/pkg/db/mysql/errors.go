package mysql

import "errors"

// ErrDBNotSet signifies the db being nil
var ErrDBNotSet = errors.New("the database is not set")
