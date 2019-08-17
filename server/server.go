package server

import (
	"QCaller/logger"
	"QCaller/server/middleware"
	"QCaller/server/routes"
	"fmt"
	"net/http"
	"os"
)

// Server : holds the params required for initializing http server
type Server struct {
	port int
}

// NewServer : returns the instance of Server
func NewServer(port int) *Server {
	return &Server{
		port: port,
	}
}

// Start : starts the server listening on the configured port
func (s *Server) Start() {
	router := routes.NewRouter()
	address := fmt.Sprintf("%v:%v", "", s.port)

	if err := http.ListenAndServe(address, mw.RequestCtx(mw.Log(mw.Auth()(router)))); err != nil {
		logger.Get().Fatalf("Unable to start server on port : [ %v ]. Err : %v", s.port, err)
		os.Exit(1)
	}
	logger.Get().Infof("Server started. Listening on : [ %v ]", s.port)
}
