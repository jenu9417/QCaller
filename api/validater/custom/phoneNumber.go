package custom

import validator "gopkg.in/go-playground/validator.v9"

// PhoneNumber : custom validator for phone number validation
// ToDo : add validation for standard formats such as E.164
var PhoneNumber = func(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) == 10
}
