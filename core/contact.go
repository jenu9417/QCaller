package core

import (
	c "QCaller/api/contract"
	"QCaller/core/util"
	"QCaller/dao"
	"QCaller/error"
	"QCaller/logger"
	"QCaller/model"
	"QCaller/types"
	"time"
)

var (
	db    dao.IContactDao
	cache dao.IContactCacheDao
)

// InitContactCore : sets dao and initialises core
func InitContactCore(d dao.IContactDao, c dao.IContactCacheDao) {
	db = d
	cache = c
}

// CreateContact : create a single contact
func CreateContact(ctx types.Context, req *c.CreateContactRequest, res *c.CreateContactResponse) *error.Error {
	req.Contact.ID = util.GetID(req.Contact.SourceID, req.Contact.Number)
	req.Contact.LastUpdated = time.Now().Unix()

	created, err := db.CreateContact(req.Contact)
	if !created || err != nil {
		logger.Get().Errorf("Error while creating the contact. Err : [%v]", err)
		logger.Get().Errorf("Failed contact : [ %v ]", req.Contact)
		return error.ErrInternalServerError("Error while creating contact", err)
	}

	return nil
}

// GetContact : get a single contact corresponding to sourceid and number
func GetContact(ctx types.Context, req *c.GetContactRequest, res *c.GetContactResponse) *error.Error {
	id := util.GetID(req.SourceID, req.Number)
	contact, err := db.GetContact(id, req.Country)
	if err != nil {
		logger.Get().Errorf("Error while getting the contact : [%v]. Err : [%v]", id, err)
		return error.ErrInternalServerError("Error while getting contact", err)
	}

	if contact == nil {
		logger.Get().Errorf("No contact found for : [%v]. Err : [%v]", id, err)
		return error.ErrNotFound("No contact found")
	}

	res.Response = contact

	return nil
}

// SearchContact : search for contacts using number and country
func SearchContact(ctx types.Context, req *c.SearchContactRequest, res *c.SearchContactResponse) *error.Error {
	if req.Immedidate {
		return searchFromCache(ctx, req, res)
	}
	return searchFromDB(ctx, req, res)
}

// searchFromCache : search contacts from cache
func searchFromCache(ctx types.Context, req *c.SearchContactRequest, res *c.SearchContactResponse) *error.Error {
	var contacts []*model.Contact
	contact, err := cache.GetContact(req.Number, req.Country)
	if err != nil {
		logger.Get().Errorf("Error while searching the contact number from cache : [%v]. Err : [%v]", req.Number, err)
		return error.ErrInternalServerError("Error while searching contact", err)
	}

	if contact == nil {
		logger.Get().Errorf("No contact found for the number : [%v] in cache. Err : [%v]", req.Number, err)
		return error.ErrNotFound("No contacts found")
	}

	contacts = append(contacts, contact)
	res.Response = contacts
	return nil
}

// searchFromDB : search contacts from DB
func searchFromDB(ctx types.Context, req *c.SearchContactRequest, res *c.SearchContactResponse) *error.Error {
	contacts, err := db.SearchContact(req.Number, req.Country, req.Size)
	if err != nil {
		logger.Get().Errorf("Error while searching the contact number : [%v]. Err : [%v]", req.Number, err)
		return error.ErrInternalServerError("Error while searching contact", err)
	}

	if contacts == nil || len(contacts) == 0 {
		logger.Get().Errorf("No contacts found for the number : [%v]. Err : [%v]", req.Number, err)
		return error.ErrNotFound("No contacts found")
	}

	res.Response = contacts
	return nil
}

// UpdateContact : update a single contact
func UpdateContact(ctx types.Context, req *c.UpdateContactRequest, res *c.UpdateContactResponse) *error.Error {
	req.Contact.ID = util.GetID(req.Contact.SourceID, req.Contact.Number)
	req.Contact.LastUpdated = time.Now().Unix()

	err := db.UpdateContact(req.Contact)
	if err != nil {
		logger.Get().Errorf("Error while updating the contact. Err : [%v]", err)
		logger.Get().Errorf("Failed contact : [ %v ]", req.Contact)
		return error.ErrInternalServerError("Error while updating contact", err)
	}

	return nil
}

// DeleteContact : delete a single contact
func DeleteContact(ctx types.Context, req *c.DeleteContactRequest, res *c.DeleteContactResponse) *error.Error {
	id := util.GetID(req.SourceID, req.Number)
	deleted, err := db.DeleteContact(id, req.Country)
	if !deleted || err != nil {
		logger.Get().Errorf("Error while deleting the contact : [%v]. Err : [%v]", id, err)
		return error.ErrInternalServerError("Error while deleting contact", err)
	}

	return nil
}

// BulkCreateContact : bulk create contacts
func BulkCreateContact(ctx types.Context, req *c.BulkCreateContactRequest, res *c.BulkCreateContactResponse) *error.Error {
	bulkRes, err := db.BulkCreateContact(req.Contacts)
	if err != nil {
		logger.Get().Errorf("Error while doing bulk create contacts. Err : [%v]", err)
		return error.ErrInternalServerError("Error while doing bulk create contacts", err)
	}

	if !bulkRes.Success {
		logger.Get().Errorf("Bulk create contact operation partially succeeded. Err : [%v]", err)
		return error.ErrBulkOpPartialSuccess("Bulk create contact operation partially succeeded", err)
	}

	res.Response = *bulkRes

	return nil
}

// BulkUpdateContact : bulk update contacts
func BulkUpdateContact(ctx types.Context, req *c.BulkUpdateContactRequest, res *c.BulkUpdateContactResponse) *error.Error {
	bulkRes, err := db.BulkUpdateContact(req.Contacts)
	if err != nil {
		logger.Get().Errorf("Error while doing bulk update contacts. Err : [%v]", err)
		return error.ErrInternalServerError("Error while doing bulk update contacts", err)
	}

	if !bulkRes.Success {
		logger.Get().Errorf("Bulk update contact operation partially succeeded. Err : [%v]", err)
		return error.ErrBulkOpPartialSuccess("Bulk update contact operation partially succeeded", err)
	}

	res.Response = *bulkRes

	return nil
}
