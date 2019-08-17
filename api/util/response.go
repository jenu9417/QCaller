package util

import (
	"QCaller/api/contract"
	"QCaller/error"
	"QCaller/types"
)

// SuccessResponse : util method for setting base params to a success response
func SuccessResponse(ctx types.Context, response contract.Response, httpCode int) contract.Response {
	response.SetRequestID(ctx.GetReqID())
	response.SetMethod(ctx.GetMethod())
	response.SetHTTPCode(httpCode)
	response.SetStatus("Success")
	return response
}

// FailureResponse : util method for setting base params to a failure response
func FailureResponse(ctx types.Context, response contract.Response, httpCode int, err *error.Error) contract.Response {
	response.SetRequestID(ctx.GetReqID())
	response.SetMethod(ctx.GetMethod())
	response.SetHTTPCode(httpCode)
	response.SetStatus("Failure")
	response.SetError(err)
	return response
}
