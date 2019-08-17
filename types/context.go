package types

import (
	"context"
	"net/http"
)

// contextKey : custom type to be used for all keys in go context
// it is adviced to not to use in built types as keys in go context
type contextKey string

const requestIDKey contextKey = "RequestID"
const methodKey contextKey = "Method"

// Context : holds the basic go context with essential functions
type Context struct {
	context.Context
}

// GetNewContext : returns an empty context
func GetNewContext() Context {
	return Context{}
}

// NewContext : creates a context with given requestId
func NewContext(ctx context.Context, requestID string) Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return Context{context.WithValue(ctx, requestIDKey, requestID)}
}

// ContextFrom : create context from http request
func ContextFrom(r *http.Request) Context {
	reqID := r.Context().Value(requestIDKey)
	ctx := NewContext(r.Context(), reqID.(string))

	return ctx.SetMethod(r.Context().Value(methodKey).(string))
}

// GetReqID : returns the request id if set or will return empty string
func (c Context) GetReqID() string {
	if requestID, ok := c.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}

// GetMethod : returns the method if set or will return empty string
func (c Context) GetMethod() string {
	if method, ok := c.Value(methodKey).(string); ok {
		return method
	}
	return ""
}

// SetMethod : creates a new context with the given data set
func (c Context) SetMethod(method string) Context {
	return Context{context.WithValue(c, contextKey(methodKey), method)}
}

// Get : gets a value from the context or nil if it does not exist
func (c Context) Get(key string) interface{} {
	value := c.Value(contextKey(key))
	return value
}

// Set : creates a new context with the given data set
func (c Context) Set(key string, value interface{}) Context {
	return Context{context.WithValue(c, contextKey(key), value)}
}
