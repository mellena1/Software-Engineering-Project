package db

import "errors"

// ErrNothingChanged should be thrown when nothing is changed
var ErrNothingChanged = errors.New("no rows were changed")
