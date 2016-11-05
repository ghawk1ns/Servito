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

// Sets the default values for config. These values might be overridden when loadConfig() is called
func init() {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("Couldn't get path for config.json")
    }
    configPath := path.Dir(filename) + "/defaultConfig.json"
	file, err := os.Open(configPath)
    defer file.Close()
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

func LoadConfigFromJSON(configJSON string) {
    bytes := []byte(configJSON)
    config = &configData{}
    err := json.Unmarshal(bytes, config)
    if err != nil {
        log.Fatal(err)
    }
    debugLog()
}
func debugLog() {
	if (config.Debug) {
        log.Printf("Config loaded: %v\n", *config)
    }
}

func LoadConfigFromPath(configPath string) {
    file, err := os.Open(configPath)
    if err != nil {
        log.Fatal(err)
    }
    LoadConfigFromFile(file)
}

func LoadConfigFromFile(file *os.File) {
    decoder := json.NewDecoder(file)
    config = &configData{}
    err := decoder.Decode(config)
    if err != nil {
        log.Fatal(err)
    }
    debugLog()
}

func SetPort(port string) {
    config.Port = port
}

func SetAddress(address string) {
    config.Address = address
}

func SetDebug(debug bool) {
    config.Debug = debug
}