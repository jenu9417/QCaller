package logger

import (
	"QCaller/config"
	"fmt"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var (
	logr *log.Logger
)

// Init : initalises the logger instance
func Init() {
	logFile := config.GetAppLog()
	if len(logFile) == 0 {
		logFile = "app.log"
	}

	fo, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to open log file : [ %v ].  Err : %v", logFile, err)
	}

	logr = log.New()
	logr.SetOutput(fo)
	logr.SetLevel(log.DebugLevel)
	logr.SetReportCaller(true)
	logr.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}

// Get : returns the current logger instance
func Get() *log.Logger {
	return logr
}
