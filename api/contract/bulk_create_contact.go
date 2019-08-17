package contract

import (
	"QCaller/model"
)

// BulkCreateContactRequest : request for bulk contact POST api
type BulkCreateContactRequest struct {
	Contacts []model.Contact `json:"Contacts" validate:"required"`
}

// BulkCreateContactResponse : response for bulk contact POST api
type BulkCreateContactResponse struct {
	BaseResponse
	Response model.BulkResponse `json:"Response,omitempty"`
}
