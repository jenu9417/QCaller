package error

import (
	"fmt"
	"net/http"
)

// list of custom errors
var (
	ErrNotFound                  = errorFunc(http.StatusNotFound, 2000, "Requested resource could not be found")
	ErrBadRequestInvalidBody     = errorFunc(http.StatusBadRequest, 2001, "Invalid request body,failed parsing ")
	ErrBadRequestValidationError = errorFunc(http.StatusBadRequest, 2002, "Bad request. Validation error.")
	ErrBulkOpPartialSuccess      = errorFunc(http.StatusPartialContent, 2003, "Bulk operation partial success")
	ErrInternalServerError       = errorFunc(http.StatusInternalServerError, 2004, "Internal server error.")
)

// Error : holds the details of error scenario
type Error struct {
	HTTPCode        int           `json:"-"`
	Code            uint64        `json:"Code"`
	Message         string        `json:"Message"`
	Description     string        `json:"Description"`
	InternalMessage []interface{} `json:"-"`
}

// Error : return string form of the error
func (err Error) Error() string {
	return fmt.Sprintf("[%s]:[%s] - %s []", err.Message, err.Description, err.InternalMessage)
}

// errorFunc : utility func for returning error instance
func errorFunc(httpCode int, code uint64, message string) func(internal ...interface{}) *Error {
	if message == "" {
		message = http.StatusText(httpCode)
	}
	return func(internal ...interface{}) *Error {
		var description string
		var ok bool
		if len(internal) > 0 {
			if description, ok = internal[0].(string); ok {
				internal = internal[1:]
			}
		}
		if description == "" {
			description = message
		}
		return &Error{
			HTTPCode:        httpCode,
			Code:            code,
			Message:         message,
			Description:     description,
			InternalMessage: internal,
		}
	}
}

// GetCode : returns numeric code corresponding to the error
func (err Error) GetCode() uint64 { return err.Code }

// AddMsg : adds internal message to an error
func (err *Error) AddMsg(msg ...interface{}) error {
	if err == nil {
		*err = *ErrInternalServerError(msg...)
		return err
	}
	err.InternalMessage = append(err.InternalMessage, msg...)
	return err
}
