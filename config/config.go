package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Config : holds all the app config params
type Config struct {
	Username   string   `json:"username" required:"true"`
	Password   string   `json:"password" required:"true"`
	ServerPort int      `json:"severPort" required:"true"`
	EsURL      []string `json:"esURL" required:"true"`
	Aerospike  struct {
		Host      string `json:"host" required:"true"`
		Port      int    `json:"port" required:"true"`
		Namespace string `json:"namespace" required:"true"`
		Retention int    `json:"retention" required:"true"`
	} `json:"as"`
	AppLog    string `json:"appLog"`
	ServerLog string `json:"serverLog"`
}

var (
	config *Config
)

// Init : initialises and returns config
func Init(configFile string) {
	filePath, _ := filepath.Abs(configFile)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read config file : [ %v ].  Err : %v", configFile, err)
	}

	conf := &Config{}
	err = json.Unmarshal(file, conf)
	if err != nil {
		log.Fatalf("Unable to parse config from file : [ %v ].  Err : %v", configFile, err)
	}

	config = conf
}

// GetUsername : returns username from config
func GetUsername() string {
	return config.Username
}

// GetPassword : returns password from config
func GetPassword() string {
	return config.Password
}

// GetServerPort : returns server port from config
func GetServerPort() int {
	return config.ServerPort
}

// GetESURL : returns es url from config
func GetESURL() []string {
	return config.EsURL
}

// GetASHost : returns as host from config
func GetASHost() string {
	return config.Aerospike.Host
}

// GetASPort : returns as port from config
func GetASPort() int {
	return config.Aerospike.Port
}

// GetASNamespace : returns as namespace from config
func GetASNamespace() string {
	return config.Aerospike.Namespace
}

// GetASRetention : returns as retention from config
func GetASRetention() int {
	return config.Aerospike.Retention
}

// GetAppLog : returns app log name from config
func GetAppLog() string {
	return config.AppLog
}

// GetServerLog : returns server log name from config
func GetServerLog() string {
	return config.ServerLog
}
