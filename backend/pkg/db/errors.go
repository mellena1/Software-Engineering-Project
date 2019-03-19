package db

import "errors"

var (
	// ErrNothingChanged should be thrown when nothing is changed
	ErrNothingChanged = errors.New("nothing was changed")
)
