package utils_test

import (
	"core_api/utils"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {

	// initialize
	viper.Set("API_PATH_PRIVATEKEY", PathPrivate)

	// generate private key
	cypherGen := utils.NewCryptoGraphic()
	assert.Equal(t, cypherGen.GenerateRsaPrivateKey(), nil)
	assert.NotEmpty(t, cypherGen.PrivateKey)

	// export private pem
	assert.Equal(t, cypherGen.ExportRsaPrivateKeyAsPemStr(), nil)
	assert.NotEmpty(t, cypherGen.PrivateKey)

	// generate token
	user_test := "test"
	user_role := true
	tokenStruct := utils.NewTokenInfo()
	tokenStruct.PrivKey = cypherGen.PrivateKey
	token, err := tokenStruct.GenerateToken(user_test, user_role)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, token)

	// dencrypt token
	err = tokenStruct.TokenDecrypt(token)
	assert.Equal(t, err, nil)

	// vaild token
	tokenClaim := utils.NewTokenInfo()
	tokenClaim.Token = tokenStruct.Token
	tokenClaim.PrivKey = cypherGen.PrivateKey
	assert.Equal(t, tokenClaim.TokenValid(), nil)
	assert.Equal(t, tokenClaim.User, user_test)
	assert.Equal(t, tokenClaim.Role, "is_admin")

	// remove test file
	os.Remove(PathConfigFile)
	os.Remove(PathPrivate)

}
