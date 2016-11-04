package servito

import (
	"encoding/json"
	"log"
	"os"
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
	file, err := os.Open("config.json")
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