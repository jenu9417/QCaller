package contract

import (
	"QCaller/model"
)

// GetContactRequest : request for contact GET api
type GetContactRequest struct {
	SourceID string `json:"SourceID" validate:"required"`
	Number   string `json:"Number" validate:"required,phoneNumber"`
	Country  string `json:"Country" validate:"required"`
}

// GetContactResponse : response for contact GET api
type GetContactResponse struct {
	BaseResponse
	Response *model.Contact `json:"Response,omitempty"`
}
