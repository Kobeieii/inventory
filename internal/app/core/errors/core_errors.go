package core_errors

import "errors"

var (
	ErrInvalidPrice    = errors.New("price must be positive")
	ErrProductNotFound = errors.New("product not found")
)
