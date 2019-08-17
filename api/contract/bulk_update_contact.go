package contract

import (
	"QCaller/model"
)

// BulkUpdateContactRequest : request for bulk contact PUT api
type BulkUpdateContactRequest struct {
	Contacts []model.Contact `json:"Contacts" validate:"required"`
}

// BulkUpdateContactResponse : response for bulk contact PUT api
type BulkUpdateContactResponse struct {
	BaseResponse
	Response model.BulkResponse `json:"Response,omitempty"`
}
