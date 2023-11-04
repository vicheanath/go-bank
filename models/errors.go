package models

import "errors"

var (
	ErrNotFound = errors.New("Requeued item is not found!")
)
