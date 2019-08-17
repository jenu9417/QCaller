package model

import "fmt"

// Contact : holds contact details
type Contact struct {
	ID          string `json:"-"`
	Name        string `json:"Name" validate:"required"`
	SourceID    string `json:"SourceID" validate:"required"`
	Country     string `json:"Country" validate:"required"`
	CountryCode string `json:"CountryCode" validate:"required"`
	Number      string `json:"Number" validate:"required,phoneNumber"`
	LastUpdated int64  `json:"-"`
}

// String : returns the string form of a contact
func (c *Contact) String() string {
	return fmt.Sprintf("Contact [ ID=%v, Name=%v, SourceID=%v, Country=%v, CountryCode=%v, Number=%v, LastUpdated=%v ]",
		c.ID, c.Name, c.SourceID, c.Country, c.CountryCode, c.Number, c.LastUpdated)
}
