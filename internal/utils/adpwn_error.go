package utils

import "errors"

// Common errors
var (
	ErrUIDRequired     = errors.New("UID is required")
	ErrFieldNotAllowed = errors.New("field is not allowed")
	ErrFieldProtected  = errors.New("field is protected")
)
