package utils_test

import (
	"core_api/utils"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGenerateApiKey(t *testing.T) {

	// initialize
	viper.Set("API_SECRET_KEY", "example_key_1234")
	
	user := struct {
		username string
		password string
		email    string
	}{
		username: "test",
		password: "test",
		email:    "test@com.com",
	}

	apikey, err := utils.GenerateApiKey(user.username, user.password, user.email)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, apikey)
}
