package contract

import "QCaller/model"

// UpdateContactRequest : request for contact PUT api
type UpdateContactRequest struct {
	Contact *model.Contact `json:"Contact" validate:"required"`
}

// UpdateContactResponse : response for contact PUT api
type UpdateContactResponse struct {
	BaseResponse
}
