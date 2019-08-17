package mw

import (
	"QCaller/server/uuid"
	"QCaller/types"
	"net/http"
)

// RequestCtx : middleware for add request id
func RequestCtx(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := types.NewContext(r.Context(), uuid.New())
		ctx = ctx.SetMethod(r.Method)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
