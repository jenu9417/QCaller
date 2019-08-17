package mw

import (
	"QCaller/config"
	"net/http"

	"github.com/goji/httpauth"
)

// Auth : authentication middleware for basic auth
func Auth() func(http.Handler) http.Handler {
	username := config.GetUsername()
	password := config.GetPassword()

	return httpauth.SimpleBasicAuth(username, password)
}
