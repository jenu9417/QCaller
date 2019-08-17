package model

// BulkResponse : response for bulk operations
type BulkResponse struct {
	Success         bool
	FailedResponses []*SingleFailedResponse `json:"FailedResponses,omitempty"`
}

// NewBulkResponse : returns instance of BulkResponse
func NewBulkResponse(success bool, failedRecs []*SingleFailedResponse) *BulkResponse {
	return &BulkResponse{
		Success:         success,
		FailedResponses: failedRecs,
	}
}

// SingleFailedResponse : response for single operation in a bulk query
type SingleFailedResponse struct {
	Code int
	ID   string
	Err  string
}

// NewSingleFailedResponse : returns instance of SingleFailedResponse
func NewSingleFailedResponse(code int, id string, err string) *SingleFailedResponse {
	return &SingleFailedResponse{
		Code: code,
		ID:   id,
		Err:  err,
	}
}
