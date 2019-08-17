package api

import (
	"QCaller/api/contract"
	"QCaller/api/util"
	"QCaller/api/validater"
	"QCaller/core"
	"QCaller/logger"
	"QCaller/types"
	"net/http"
	"strconv"
)

// SearchContact : handles the request, validates, operates and sends response
func SearchContact(r *http.Request, ctx types.Context) contract.Response {
	number := r.FormValue("number")
	country := r.FormValue("country")
	immediate := false
	if r.FormValue("immediate") == "true" {
		immediate = true
	}

	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		size = 5
	}

	request := &contract.SearchContactRequest{Number: number, Country: country, Immedidate: immediate, Size: size}
	response := &contract.SearchContactResponse{}

	if err := validater.Of(request).Validate(); err != nil {
		logger.Get().Errorf("Validation error. Err : %v", err)
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	if err := core.SearchContact(ctx, request, response); err != nil {
		return util.FailureResponse(ctx, response, err.HTTPCode, err)
	}

	return util.SuccessResponse(ctx, response, http.StatusOK)
}
