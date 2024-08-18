package model

import "errors"

var (
	// ErrNotFound is returned when in database a record is not found.
	ErrNotFound = errors.New("not found")

	// ErrAlreadyExists is returned when in database a record already exists.
	ErrAlreadyExists = errors.New("already exists")

	// ErrWithoutChanges is returned when there are no changes in the request.
	ErrWithoutChanges = errors.New("without changes")
)
