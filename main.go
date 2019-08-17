package main

import (
	"QCaller/api/validater"
	"QCaller/config"
	"QCaller/core"
	"QCaller/logger"
	"QCaller/server"
	"log"
	"os"
	"runtime"
)

func main() {

	// panic handling
	defer handlePanic()

	//init config
	config.Init("config.json")

	//init logger
	logger.Init()
	logger.Get().Debug("Logger Initialized")

	//init validater
	validater.Init()
	logger.Get().Debug("Validater Initialized")

	//init core
	core.Init()
	logger.Get().Debug("Core Initialized")

	//init api server
	server := server.NewServer(config.GetServerPort())
	go server.Start()

	select {}
}

// handlePanic : defer func to handle panic and initiate a shutdown
// ToDo : create a signal catcher to avoid abrupt kill
func handlePanic() {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		stackSize := runtime.Stack(buf, false)
		log.Println("Main Stack Trace : ", string(buf[0:stackSize]))
		logger.Get().Infof("QCaller : Panic handler : %v", r)
		logger.Get().Fatalf("Stack Trace : %v", string(buf[0:stackSize]))
		shutDown()
	}
}

// shutDown : logs shutdown and exit
func shutDown() {
	log.Println("Graceful Shutdown")
	logger.Get().Fatalf("Graceful Shutdown")
	os.Exit(1)
}
