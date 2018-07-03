package error

import (
  "errors"
)

var (
  ErrUnrecognizedServiceModel = errors.New("Unrecognized service model.")
  ErrUnrecognizedDBModel      = errors.New("Unrecognized database model.")
  ErrConverterError           = errors.New("Converter did not work correctly.")
)
