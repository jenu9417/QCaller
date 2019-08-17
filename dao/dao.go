package dao

import (
	"QCaller/config"
	"QCaller/dao/as"
	"QCaller/dao/es"
	"QCaller/db"
)

// ToDo : can make it a factory if a more dynamic logic needed. eg., in case of multiple datastores.

// NewContactDao : returns contact dao impl
func NewContactDao() (IContactDao, error) {
	return newESContactDao()
}

// NewContactCacheDao : returns contact cache dao impl
func NewContactCacheDao() (IContactCacheDao, error) {
	return newASContactCacheDao()
}

// newESContactDao : returns es impl of contact dao
func newESContactDao() (IContactDao, error) {
	esEndPoint := config.GetESURL()
	esclient, err := db.NewESClient(esEndPoint)
	if err != nil {
		return nil, err
	}

	return es.NewContactDao(esclient), nil
}

// newASContactCacheDao : returns as impl of contact dao
func newASContactCacheDao() (IContactCacheDao, error) {
	asHost := config.GetASHost()
	asPort := config.GetASPort()

	asclient, err := db.NewASClient(asHost, asPort)
	if err != nil {
		return nil, err
	}

	asNamespace := config.GetASNamespace()
	asRetention := config.GetASRetention()

	return as.NewContactCacheDao(asclient, asNamespace, asRetention), nil
}
