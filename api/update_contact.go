package api

import (
	"QCaller/api/contract"
	"QCaller/api/util"
	"QCaller/api/validater"
	"QCaller/core"
	"QCaller/error"
	"QCaller/logger"
	"QCaller/model"
	"QCaller/types"
	"encoding/json"
	"net/http"
)

// UpdateContact : handles the request, validates, operates and sends response
func UpdateContact(r *http.Request, ctx types.Context) contract.Response {
	contact := &model.Contact{}
	response := &contract.UpdateContactResponse{}

	if err := json.NewDecoder(r.Body).Decode(contact); err != nil {
		logger.Get().Errorf("Unable to parse request body. Err : %v", err)
		er := error.ErrBadRequestInvalidBody(err)
		return util.FailureResponse(ctx, response, er.HTTPCode, er)
	}

	request := &contract.UpdateContactRequest{Contact: contact}

	if err := validater.Of(contact).Validate(); err != nil {
		logger.Get().Errorf("Validation error. Err : %v", err)
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	if err := core.UpdateContact(ctx, request, response); err != nil {
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	return util.SuccessResponse(ctx, response, http.StatusOK)
}
