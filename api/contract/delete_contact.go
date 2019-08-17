package contract

// DeleteContactRequest : request for contact DELETE api
type DeleteContactRequest struct {
	SourceID string `json:"SourceID" validate:"required"`
	Number   string `json:"Number" validate:"required,phoneNumber"`
	Country  string `json:"Country" validate:"required"`
}

// DeleteContactResponse : response for contact DELETE api
type DeleteContactResponse struct {
	BaseResponse
}
