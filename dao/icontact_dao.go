package dao

import (
	"QCaller/model"
)

// IContactDao : interface to dao operation on contact
type IContactDao interface {
	CreateContact(contact *model.Contact) (bool, error)
	GetContact(id string, country string) (*model.Contact, error)
	SearchContact(number string, country string, size int) ([]*model.Contact, error)
	UpdateContact(contact *model.Contact) error
	DeleteContact(id string, country string) (bool, error)
	BulkCreateContact(contacts []model.Contact) (*model.BulkResponse, error)
	BulkUpdateContact(contacts []model.Contact) (*model.BulkResponse, error)
}
