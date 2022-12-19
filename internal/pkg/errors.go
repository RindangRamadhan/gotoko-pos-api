package pkg

import "errors"

var (
	ErrFatalQuery   = errors.New("fatal query error")
	ErrDataNotFound = errors.New("data not found")
)
