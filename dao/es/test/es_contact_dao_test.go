package test

import (
	"QCaller/core/util"
	des "QCaller/dao/es"
	"QCaller/db/es"
	"QCaller/model"
	"fmt"
	"testing"
	"time"

	"gopkg.in/olivere/elastic.v3"
)

var esEndPoint = []string{"http://localhost:9200"}

func TestCreateContact(t *testing.T) {
	contact := getDefaultContact()
	dao := des.NewContactDao(getESClient(t))

	res, err := dao.CreateContact(contact)
	if !res || err != nil {
		t.Errorf("Failed to create contact. Err : %v", err)
	}
}

func TestGetContact(t *testing.T) {
	contact := getDefaultContact()
	dao := des.NewContactDao(getESClient(t))

	res, err := dao.GetContact(contact.ID, contact.Country)
	if res == nil || err != nil {
		t.Errorf("Failed to get contact. Err : %v", err)
	}

	fmt.Println(contact)
}

func TestUpdateContact_ByName(t *testing.T) {
	contact := getDefaultContact()
	dao := des.NewContactDao(getESClient(t))
	contact.Name = "newName"

	err := dao.UpdateContact(contact)
	if err != nil {
		t.Errorf("Failed to update contact by name. Err : %v", err)
	}
}

func TestUpdateContact_ByNumber(t *testing.T) {
	contact := getDefaultContact()
	dao := des.NewContactDao(getESClient(t))
	contact.Number = "+919090909090"

	err := dao.UpdateContact(contact)
	if err != nil {
		t.Errorf("Failed to update contact by number. Err : %v", err)
	}
}

func TestDeleteContact(t *testing.T) {
	contact := getDefaultContact()
	dao := des.NewContactDao(getESClient(t))

	res, err := dao.DeleteContact(contact.ID, contact.Country)
	if !res || err != nil {
		t.Errorf("Failed to delete contact. Err : %v", err)
	}
}

func getESClient(t *testing.T) *elastic.Client {
	rConf := es.NewRetryConf(2*time.Second, 10*time.Second)
	eConf := es.NewConf(esEndPoint, rConf)
	var err error
	es, err := es.InitESClient(eConf)
	if err != nil {
		t.Errorf("error initializing es client: " + err.Error())
	}

	return es
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
