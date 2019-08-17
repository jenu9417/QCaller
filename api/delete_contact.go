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

// DeleteContact : handles the request, validates, operates and sends response
func DeleteContact(r *http.Request, ctx types.Context) contract.Response {
	sourceID := r.FormValue("sourceID")
	number := r.FormValue("number")
	country := r.FormValue("country")

	request := &contract.DeleteContactRequest{SourceID: sourceID, Number: number, Country: country}
	response := &contract.DeleteContactResponse{}

	if err := validater.Of(request).Validate(); err != nil {
		logger.Get().Errorf("Validation error. Err : %v", err)
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	if err := core.DeleteContact(ctx, request, response); err != nil {
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	return util.SuccessResponse(ctx, response, http.StatusOK)
}
