package api

import "errors"

// request validation errors
var ErrorRequestBodyEmpty = errors.New("empty request body")
var ErrorRequestFieldMissing = errors.New("request field missing")


// conversion to entity validations
var ErrorConvesionToEntity = errors.New("cannot convert request to entity")