package error

import "errors"

var (
	ErrIncorrectParamsFormat = errors.New("Incorrect input parameter format.")
	ErrInsufficientParams    = errors.New("Insufficient parameters.")
)
