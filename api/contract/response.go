package contract

import "QCaller/error"

// Response : defines the interface all response objects should conform to
type Response interface {
	SetRequestID(string) Response
	SetHTTPCode(int) Response
	SetMethod(string) Response
	SetStatus(string) Response
	SetError(*error.Error) Response
}

// BaseResponse : base api response params
type BaseResponse struct {
	RequestID string `json:"-"`
	HTTPCode  int
	Method    string
	Status    string
	Error     *error.Error `json:"Error,omitempty"`
}

// SetRequestID : set request id to response
func (b *BaseResponse) SetRequestID(id string) Response {
	b.RequestID = id
	return b
}

// SetHTTPCode : set http code to response
func (b *BaseResponse) SetHTTPCode(code int) Response {
	b.HTTPCode = code
	return b
}

// SetMethod : set http method to response
func (b *BaseResponse) SetMethod(method string) Response {
	b.Method = method
	return b
}

// SetStatus : set status to response
func (b *BaseResponse) SetStatus(status string) Response {
	b.Status = status
	return b
}

// SetError : set error details to response
func (b *BaseResponse) SetError(err *error.Error) Response {
	b.Error = err
	return b
}
