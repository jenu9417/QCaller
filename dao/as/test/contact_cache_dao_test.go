package test

import (
	"QCaller/core/util"
	cc "QCaller/dao/as"
	"QCaller/db"
	"QCaller/db/as"
	"QCaller/model"
	"fmt"
	"testing"
	"time"

	aerospike "github.com/aerospike/aerospike-client-go"
)

var (
	host      = "127.0.0.1"
	port      = 3000
	namespace = "prefix"
	retention = 86400
)

func TestGetContact(t *testing.T) {
	client, err := db.NewASClient(host, port)
	contact := getDefaultContact()

	cache := cc.NewContactCacheDao(client, namespace, retention)

	res, err := cache.GetContact(contact.Number, contact.Country)
	if err != nil {
		t.Errorf("Failed to create contact. Err : %v", err)
	}

	fmt.Println(res)
}

func TestPutContact(t *testing.T) {
	client, err := db.NewASClient(host, port)
	contact := getDefaultContact()
	contact.Name = "newName"

	cache := cc.NewContactCacheDao(client, namespace, retention)

	err = cache.PutContact(contact)
	if err != nil {
		t.Errorf("Failed to create contact. Err : %v", err)
	}
}

// NewASClient : returns instance of as client
func getASClient(host string, port int) (*aerospike.Client, error) {
	aConf := as.NewConf(host, port)
	return as.InitASClient(aConf)
}

func getDefaultContact() *model.Contact {
	return &model.Contact{
		ID:          util.GetID("564d32", "657657657"),
		Name:        "janax",
		Country:     "india",
		SourceID:    "564d32",
		CountryCode: "+91",
		Number:      "657657657",
		LastUpdated: time.Now().Unix(),
	}
}
