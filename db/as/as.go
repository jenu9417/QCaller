package as

import (
	as "github.com/aerospike/aerospike-client-go"
)

// InitASClient : initializes and returns AS client
func InitASClient(conf *Conf) (*as.Client, error) {
	return as.NewClient(conf.Host(), conf.Port())
}
