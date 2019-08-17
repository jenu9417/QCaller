package api

import (
	"QCaller/api/contract"
	"QCaller/api/util"
	"QCaller/api/validater"
	"QCaller/core"
	"QCaller/error"
	"QCaller/logger"
	"QCaller/types"
	"encoding/json"
	"net/http"
)

// BulkUpdateContact : handles the request, validates, operates and sends response
func BulkUpdateContact(r *http.Request, ctx types.Context) contract.Response {
	request := &contract.BulkUpdateContactRequest{}
	response := &contract.BulkUpdateContactResponse{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		logger.Get().Errorf("Unable to parse request body. Err : %v", err)
		er := error.ErrBadRequestInvalidBody(err)
		return util.FailureResponse(ctx, response, er.HTTPCode, er)
	}

	for _, c := range request.Contacts {
		contact := &c
		if err := validater.Of(contact).Validate(); err != nil {
			logger.Get().Errorf("Validation error. Err : %v", err)
			return util.FailureResponse(ctx, response, err.HTTPCode, err)
		}
	}

	if err := core.BulkUpdateContact(ctx, request, response); err != nil {
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	return util.SuccessResponse(ctx, response, http.StatusCreated)
}
