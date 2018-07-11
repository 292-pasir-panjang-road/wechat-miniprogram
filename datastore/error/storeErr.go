package error

import (
	"database/sql"
	"errors"
)

var (
	ErrUnrecognizedServiceModel = errors.New("Unrecognized service model.")
	ErrUnrecognizedDBModel      = errors.New("Unrecognized database model.")
	ErrConverterError           = errors.New("Converter did not work correctly.")
	ErrInvalidInputModel        = errors.New("Input model type is not correct")
	ErrDBRetrievalError         = errors.New("Failed to retrieve from db.")
	ErrInvalidQuery             = errors.New("Query not valid.")
	ErrNotSupportedQuery        = errors.New("Query not supported yet.")
	ErrNotFound                 = errors.New("Not found.")
)

func ErrorWrapper(err error) error {
	switch {
	case err == sql.ErrNoRows:
		return ErrNotFound
	default:
		return ErrDBRetrievalError
	}
}
