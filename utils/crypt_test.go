package utils_test

import (
	"core_api/utils"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var PathPrivate string = "test.pem"

func TestRsaFundamental(t *testing.T) {

	// initialize
	viper.Set("API_PATH_PRIVATEKEY", PathPrivate)

	// test func GenerateRsaPrivateKey
	cypherGen := utils.NewCryptoGraphic()
	assert.Equal(t, cypherGen.GenerateRsaPrivateKey(), nil)
	assert.NotEmpty(t, cypherGen.PrivateKey)

	// test func ExportRsaPrivateKeyAsPemStr
	assert.Equal(t, cypherGen.ExportRsaPrivateKeyAsPemStr(), nil)
	_, err := os.Stat(PathPrivate)
	assert.Equal(t, err, nil)

	// test func LoadRsaPrivatekey
	cypherLoad := utils.NewCryptoGraphic()
	assert.Equal(t, cypherLoad.LoadRsaPrivatekey(), nil)
}

func TestHashPassword(t *testing.T) {

	// initialize
	viper.Set("API_SECRET_KEY", "example_key_1234")

	// test func HashPassword
	password := "test"
	hashCypher := utils.NewCryptoGraphic()
	passwordHashed, err := hashCypher.HashPassword(password)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, passwordHashed)

	//test func DoPasswordsMatch
	result := hashCypher.DoPasswordsMatch(passwordHashed, password)
	assert.Equal(t, result, true)
}
