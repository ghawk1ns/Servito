package servito

import (
	"encoding/json"
	"log"
	"os"
    "runtime"
    "path"
)

var config *configData

// Config corresponding to config.json go in here
type configData struct {
    Debug        bool   `json:"debug"`
    // the port:address your server will listen at
    Port         string `json:"port"`
    Address      string `json:"address"`
    // timeout for reading/writing in seconds
    ReadTimeout  int64  `json:"readTimeout"`
    WriteTimeout int64  `json:"writeTimeout"`
    // csrf support
    CSRFEnable   bool   `json:"csrfEnable"`
    CSRFKey      string `json:"csrfKey"`
}

func init() {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("Couldn't get path for config.json")
    }
    configPath := path.Dir(filename) + "/config.json"
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	config = &configData{}
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal(err)
	}
}