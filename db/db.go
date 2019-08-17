package db

import (
	"QCaller/db/as"
	"QCaller/db/es"
	"time"

	"github.com/aerospike/aerospike-client-go"
	elastic "gopkg.in/olivere/elastic.v3"
)

// ToDo : use a factory when more datastores included

// NewESClient : returns instance of es client
func NewESClient(esEndPoint []string) (*elastic.Client, error) {
	rConf := es.NewRetryConf(2*time.Second, 10*time.Second)
	eConf := es.NewConf(esEndPoint, rConf)
	return es.InitESClient(eConf)
}

// NewASClient : returns instance of as client
func NewASClient(host string, port int) (*aerospike.Client, error) {
	aConf := as.NewConf(host, port)
	return as.InitASClient(aConf)
}
