package validater

import (
	"QCaller/api/validater/custom"
	"QCaller/error"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	valdtr *validator.Validate
)

// IValidater :
type IValidater interface {
	Validate() *error.Error
}

// Init :
func Init() {
	v := validator.New()
	_ = v.RegisterValidation("phoneNumber", custom.PhoneNumber)

	valdtr = v
}

// Of :
func Of(strct interface{}) IValidater {
	return NewGenericValidator(valdtr, strct)
}
