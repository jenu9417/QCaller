package dao

import "QCaller/model"

// IContactCacheDao : interface to cache for fast lookup of contacts
type IContactCacheDao interface {
	GetContact(number string, country string) (*model.Contact, error)
	PutContact(contact *model.Contact) error
}
