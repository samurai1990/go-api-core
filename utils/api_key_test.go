package utils_test

import (
	"core_api/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateApiKey(t *testing.T) {
	user := struct {
		username string
		password string
		email    string
	}{
		username: "test",
		password: "test",
		email:    "test@com.com",
	}
    assert.Equal(t,utils.NewConfig().LoadConfig(".."),nil)
	apikey, err := utils.GenerateApiKey(user.username, user.password, user.email)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, apikey)
}
