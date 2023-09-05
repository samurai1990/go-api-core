package utils_test

import (
	"core_api/utils"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PathConfigFile string = "app.env"

func TestLoadConfigl(t *testing.T) {

	// initialize env
	env := Initialize()

	// test LoadConfig
	conf := utils.NewConfig()
	conf.LoadConfig(".")

	for key, val := range env {
		assert.Equal(t, os.Getenv(key), val)
	}

	// remove test env file
	err := os.Remove(PathConfigFile)
	assert.Equal(t, err, nil)

	err = os.Remove(PathPrivate)
	assert.Equal(t, err, nil)
}

func Initialize() map[string]string {

	env := make(map[string]string)
	env["API_DB_HOST"] = "192.168.1.1"
	env["API_DB_NAME"] = "db"
	env["API_DB_USER"] = "test"
	env["API_DB_PASSWORD"] = "test"
	env["API_DB_PORT"] = "1000"
	env["API_SECRET_KEY"] = "example_key_1234"
	env["API_BIND_HOST"] = "192.168.1.100"
	env["API_BIND_PORT"] = "8000"
	env["API_PATH_PRIVATEKEY"] = "test.pem"

	// Set the environment variables
	for key, val := range env {
		os.Setenv(key, val)
	}

	// Open the file for writing
	file, err := os.Create(PathConfigFile)
	if err != nil {
		log.Fatal("Error creating test.env file")
	}
	defer file.Close()

	// Write the environment variables to the file
	for _, e := range os.Environ() {
		fmt.Fprintln(file, e)
	}
	return env
}
