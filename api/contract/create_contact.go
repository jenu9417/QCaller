package contract

import (
	"QCaller/model"
)

// CreateContactRequest : request for contact POST api
type CreateContactRequest struct {
	Contact *model.Contact `json:"Contact" validate:"required"`
}

// CreateContactResponse : response for contact POST api
type CreateContactResponse struct {
	BaseResponse
}
