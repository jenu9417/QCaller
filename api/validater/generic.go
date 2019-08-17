package validater

import (
	"QCaller/error"

	validator "gopkg.in/go-playground/validator.v9"
)

// GenericValidator : holds params for validating struct
type GenericValidator struct {
	strct  interface{}
	valdtr *validator.Validate
}

// NewGenericValidator : returns instance of GenericValidator
func NewGenericValidator(v *validator.Validate, s interface{}) *GenericValidator {
	return &GenericValidator{
		valdtr: v,
		strct:  s,
	}
}

// Validate : starts validation and returns error on failure
func (c *GenericValidator) Validate() *error.Error {
	err := c.valdtr.Struct(c.strct)
	if err != nil {
		return error.ErrBadRequestValidationError(err)
	}
	return nil
}
