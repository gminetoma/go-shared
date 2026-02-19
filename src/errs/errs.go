package errs

import "errors"

var (
	ErrNotFound        = errors.New("error.not-found")
	ErrUnauthenticated = errors.New("error.unauthenticated")
)
