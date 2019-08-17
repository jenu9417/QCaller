package as

import (
	"QCaller/logger"
	"QCaller/model"
	"fmt"

	as "github.com/aerospike/aerospike-client-go"
)

var (
	// BinNames : fields to be synced to aerospike
	BinNames = []string{
		"ID",
		"Name",
		"SourceID",
		"Country",
		"CountryCode",
		"Number",
	}
)

// ContactCacheDao : holds the aerospike client
type ContactCacheDao struct {
	client    *as.Client
	namespace string
	retention uint32
}

// NewContactCacheDao : returns instance of ContactCacheDao
func NewContactCacheDao(client *as.Client, namespace string, retention int) *ContactCacheDao {
	return &ContactCacheDao{
		client:    client,
		namespace: namespace,
		retention: uint32(retention),
	}
}

// GetContact : gets the contact with the number from cache
func (c *ContactCacheDao) GetContact(number string, country string) (*model.Contact, error) {
	key, err := as.NewKey(c.namespace, country, number)
	if err != nil {
		logger.Get().Errorf("Error while creating as key for : [ %v ]. Err : %v", number, err)
		return nil, err
	}

	policy := getReadPolicy()
	record, err := c.client.Get(policy, key, BinNames...)
	if err != nil {
		logger.Get().Errorf("Error while getting contact for : [ %v ]. Err : %v", number, err)
		return nil, err
	}

	return convertToContact(record.Bins)
}

// PutContact : put a contact to cache
func (c *ContactCacheDao) PutContact(contact *model.Contact) error {
	bins := convertToBins(contact)
	key, err := as.NewKey(c.namespace, contact.Country, contact.Number)
	if err != nil {
		logger.Get().Errorf("Error while creating as key for : [ %v ]. Err : %v", contact.Number, err)
		return err
	}

	policy := getWritePolicy(0, c.retention)
	err = c.client.PutBins(policy, key, bins...)
	if err != nil {
		logger.Get().Errorf("Error while putting contact to as for : [ %v ]. Err : %v", contact.Number, err)
		return err
	}

	return nil
}

// convertToContact : convert as bins to Contact
func convertToContact(bins as.BinMap) (*model.Contact, error) {
	contact := &model.Contact{}
	var err error
	errStr := "Failed parsing field: "
	for key, value := range bins {
		if value == nil {
			continue //Silently ignore empty fields
		}
		switch key {
		case "ID":
			contact.ID, err = parseStringField(value, key, errStr)
		case "Name":
			contact.Name, err = parseStringField(value, key, errStr)
		case "SourceID":
			contact.SourceID, err = parseStringField(value, key, errStr)
		case "Country":
			contact.Country, err = parseStringField(value, key, errStr)
		case "CountryCode":
			contact.CountryCode, err = parseStringField(value, key, errStr)
		case "Number":
			contact.Number, err = parseStringField(value, key, errStr)
		default:
			logger.Get().Warnf("Unknown key : %v", key)
		}
		if err != nil {
			logger.Get().Errorf("Error while parsing contact from as bins. Err : %v", err)
			return nil, err
		}
	}

	return contact, nil
}

// convertToBins : convert Contact to as bins
func convertToBins(contact *model.Contact) []*as.Bin {
	var asBins []*as.Bin
	asBins = append(asBins, as.NewBin("ID", contact.ID))
	asBins = append(asBins, as.NewBin("Name", contact.Name))
	asBins = append(asBins, as.NewBin("SourceID", contact.SourceID))
	asBins = append(asBins, as.NewBin("Country", contact.Country))
	asBins = append(asBins, as.NewBin("CountryCode", contact.CountryCode))
	asBins = append(asBins, as.NewBin("Number", contact.Number))

	return asBins
}

// getReadPolicy : default as read policy
func getReadPolicy() *as.BasePolicy {
	rp := as.NewPolicy()
	rp.ConsistencyLevel = as.CONSISTENCY_ALL
	return rp
}

// getWritePolicy : default as write policy
func getWritePolicy(generation uint32, expiration uint32) *as.WritePolicy {
	wp := as.NewWritePolicy(generation, expiration)
	wp.RecordExistsAction = as.UPDATE
	wp.ConsistencyLevel = as.CONSISTENCY_ALL
	return wp
}

// parseStringField : util method for parsing string val from as bin
func parseStringField(val interface{}, key string, errStr string) (string, error) {
	casted, ok := val.(string)
	if ok {
		return casted, nil
	}
	return "", fmt.Errorf("%s %s Inside ParseStringField. Type: %T Value: %v", errStr, key, val, val)
}
