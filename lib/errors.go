package lib

import "errors"

var (
	ErrUnableToCreateRecord = errors.New("unable to create record")
	ErrUnableToFindRecord   = errors.New("unable to find record")
)
