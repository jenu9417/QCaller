package api

import (
	"QCaller/api/contract"
	"QCaller/api/util"
	"QCaller/api/validater"
	"QCaller/core"
	"QCaller/logger"
	"QCaller/types"
	"net/http"
)

// GetContact : handles the request, validates, operates and sends response
func GetContact(r *http.Request, ctx types.Context) contract.Response {
	sourceID := r.FormValue("sourceID")
	number := r.FormValue("number")
	country := r.FormValue("country")

	request := &contract.GetContactRequest{SourceID: sourceID, Number: number, Country: country}
	response := &contract.GetContactResponse{}

	if err := validater.Of(request).Validate(); err != nil {
		logger.Get().Errorf("Validation error. Err : %v", err)
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	if err := core.GetContact(ctx, request, response); err != nil {
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	return util.SuccessResponse(ctx, response, http.StatusOK)
}
