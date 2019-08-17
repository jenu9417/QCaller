package core

import (
	"QCaller/dao"
	"QCaller/logger"
	"os"
)

// Init : initialises all the core logic
func Init() {

	contactDao, err := dao.NewContactDao()
	if err != nil {
		logger.Get().Fatalf("Error while creating contact dao. Err : %v", err)
		os.Exit(1)
	}

	cacheDao, err := dao.NewContactCacheDao()
	if err != nil {
		logger.Get().Fatalf("Error while creating contact cache dao. Err : %v", err)
		os.Exit(1)
	}

	InitContactCore(contactDao, cacheDao)
}
