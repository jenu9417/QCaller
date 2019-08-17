package mw

import (
	"QCaller/config"
	"QCaller/logger"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// Log : api request logging middleware
func Log(h http.Handler) http.Handler {
	serLog := config.GetServerLog()
	if len(serLog) == 0 {
		serLog = "server.log"
	}

	logFile, err := os.OpenFile(serLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Get().Fatalf("Unable to open server log file : [ %v ]. Err : %v", serLog, err)
		os.Exit(1)
	}

	return handlers.LoggingHandler(logFile, h)
}
