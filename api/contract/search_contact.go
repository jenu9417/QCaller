package contract

import (
	"QCaller/model"
)

// SearchContactRequest : request for contact search GET api
type SearchContactRequest struct {
	Number     string `json:"Number" validate:"required,phoneNumber"`
	Country    string `json:"Country" validate:"required"`
	Immedidate bool   `json:"Immediate"`
	Size       int    `json:"Size" validate:"gte=0,lte=50"`
}

// SearchContactResponse : response for contact search GET api
type SearchContactResponse struct {
	BaseResponse
	Response []*model.Contact `json:"Response,omitempty"`
}
