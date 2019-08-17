package es

import (
	"QCaller/core/util"
	"QCaller/model"
	"encoding/json"
	"fmt"
	"time"

	elastic "gopkg.in/olivere/elastic.v3"
)

const typeNum = "number"

// ContactDao : holds the elasticsearch client
type ContactDao struct {
	client *elastic.Client
}

// NewContactDao : returns instance of ContactDao
func NewContactDao(client *elastic.Client) *ContactDao {
	return &ContactDao{
		client: client,
	}
}

// CreateContact : Create single contact
func (e *ContactDao) CreateContact(contact *model.Contact) (bool, error) {
	index := contact.Country
	typ := typeNum
	id := contact.ID

	contact.LastUpdated = time.Now().Unix()
	res, err := e.client.Index().Index(index).Type(typ).Id(id).BodyJson(contact).Do()

	return res.Created, err
}

// GetContact : Get single contact by id
func (e *ContactDao) GetContact(id, country string) (*model.Contact, error) {
	index := country
	typ := typeNum

	result, err := e.client.Get().Index(index).Type(typ).Id(id).Do()
	if err != nil {
		if eError, ok := err.(*elastic.Error); ok == true {
			if eError.Status == 404 && eError.Details == nil {
				// doc not found in the index
				fmt.Println("Doc Not found")
				return nil, nil
			}
			// index itself not found
			fmt.Println("Index not found : ")
			fmt.Println(eError.Details)
			return nil, err
		}
	}

	contact := new(model.Contact)
	err = json.Unmarshal([]byte(*result.Source), contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

// SearchContact : search for contacts by number and country
func (e *ContactDao) SearchContact(number, country string, size int) ([]*model.Contact, error) {
	index := country
	typ := typeNum
	query := formHeuristicESQuery(number)

	result, err := e.client.Search().Index(index).Type(typ).Query(query).Size(size).Sort("_score", false).Do()
	if err != nil {
		return nil, err
	}

	var contacts []*model.Contact
	for _, hit := range result.Hits.Hits {
		if hit.Source != nil {
			contact := new(model.Contact)
			err = json.Unmarshal([]byte(*hit.Source), contact)
			if err != nil {
				return nil, nil
			}
			contacts = append(contacts, contact)
		}
	}
	return contacts, nil
}

// formHeuristicESQuery : es query to rank the contacts while searching to provide top results
// ToDo : make it extendable to add more custom heuristic logic.
// use a better design like chain of responsiblity or command
func formHeuristicESQuery(number string) elastic.Query {
	numTerm := elastic.NewTermQuery("Number", number)
	spamTerm := elastic.NewTermQuery("Name", "spam").Boost(-1)
	query := elastic.NewBoolQuery().Filter(numTerm).Should(spamTerm)

	return query
}

// UpdateContact : update a single contact
// Name change - id(source&num) will exist. Hence update
// Number change - id(source&num) wont exist. Hence upsert
func (e *ContactDao) UpdateContact(contact *model.Contact) error {
	index := contact.Country
	typ := typeNum
	id := contact.ID

	_, err := e.client.Update().Id(id).Index(index).Type(typ).Doc(contact).DocAsUpsert(true).Do()

	return err
}

// DeleteContact : delete a single contact
func (e *ContactDao) DeleteContact(id, country string) (bool, error) {
	index := country
	typ := typeNum

	result, err := e.client.Delete().Index(index).Type(typ).Id(id).Do()

	return result.Found, err
}

// BulkCreateContact : bulk create contacts
func (e *ContactDao) BulkCreateContact(contacts []model.Contact) (*model.BulkResponse, error) {
	var bulkReq []elastic.BulkableRequest

	for _, contact := range contacts {
		contact.ID = util.GetID(contact.SourceID, contact.Number)
		contact.LastUpdated = time.Now().Unix()
		req := elastic.NewBulkIndexRequest().
			Index(contact.Country).
			Type(typeNum).
			Id(contact.ID).
			Doc(contact)
		bulkReq = append(bulkReq, req)
	}

	return executeBulkRequest(e.client, bulkReq)
}

// BulkUpdateContact : bulk update contacts
// Can this really happen in real life?
func (e *ContactDao) BulkUpdateContact(contacts []model.Contact) (*model.BulkResponse, error) {
	var bulkReq []elastic.BulkableRequest

	for _, contact := range contacts {
		contact.ID = util.GetID(contact.SourceID, contact.Number)
		contact.LastUpdated = time.Now().Unix()
		req := elastic.NewBulkUpdateRequest().
			Index(contact.Country).
			Type(typeNum).
			Id(contact.ID).
			Doc(contact).
			DocAsUpsert(true)
		bulkReq = append(bulkReq, req)
	}

	return executeBulkRequest(e.client, bulkReq)
}

// executeBulkRequest : executes the bulk request and returns response
func executeBulkRequest(client *elastic.Client, bulkReq []elastic.BulkableRequest) (*model.BulkResponse, error) {
	bulkRes, bulkErr := client.Bulk().
		Add(bulkReq...).
		Do()

	var failedResponses []*model.SingleFailedResponse
	if bulkErr != nil || len(bulkRes.Failed()) > 0 {
		for _, failed := range bulkRes.Failed() {
			response := model.NewSingleFailedResponse(failed.Status, failed.Id, failed.Error.Reason)
			failedResponses = append(failedResponses, response)
		}
		return model.NewBulkResponse(false, failedResponses), bulkErr
	}

	return model.NewBulkResponse(true, nil), nil
}
