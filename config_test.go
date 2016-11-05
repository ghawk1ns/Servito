package servito

import (
    "testing"
    "io/ioutil"
    "os"
    "log"
)

const TEST_CONFIG_PATH = "/tmp/testConfig.txt"

const TEST_DATA = `{
        "debug": false,
        "address": "coolio",
        "port": "foo",
        "readTimeout": 9001,
        "writeTimeout": 42
    }`

func testConfig(t *testing.T) {
    if  config.Debug ||
        config.Address != "coolio" ||
        config.Port != "foo" ||
        config.ReadTimeout != 9001 ||
        config.WriteTimeout != 42 {
        t.Fail()
    }
}

func testErr(e error, t *testing.T) {
    if e != nil {
        log.Println(e)
        t.Fail()
    }
}

func TestLoadConfigFromJSON(t *testing.T) {
    LoadConfigFromJSON(TEST_DATA)
    testConfig(t)
}

func TestLoadConfigFromFile(t *testing.T) {
    bytes := []byte(TEST_DATA)
    err := ioutil.WriteFile(TEST_CONFIG_PATH, bytes, 0644)
    defer os.Remove(TEST_CONFIG_PATH)
    testErr(err,t)
    f, err := os.Open(TEST_CONFIG_PATH)
    testErr(err,t)
    LoadConfigFromFile(f)
    testConfig(t)
}

func TestLoadConfigFromPath(t *testing.T) {
    bytes := []byte(TEST_DATA)
    err := ioutil.WriteFile(TEST_CONFIG_PATH, bytes, 0644)
    defer os.Remove(TEST_CONFIG_PATH)
    testErr(err,t)
    LoadConfigFromPath(TEST_CONFIG_PATH)
    testConfig(t)
}
