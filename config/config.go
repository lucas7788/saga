package config

import "github.com/ontio/ontology/common/log"

var Version = ""


var (
	DEFAULT_LOG_LEVEL = log.InfoLog
	DEFAULT_REST_PORT = uint(8080)
)

type Config struct {
	RestPort    uint   `json:"rest_port"`
	Version     string `json:"version"`
}


var DefConfig = &Config{
	RestPort: DEFAULT_REST_PORT,
	Version:  "1.0.0",
}