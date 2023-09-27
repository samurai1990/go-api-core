package api

import (
	"core_api/database"
	error_code "core_api/errors"
	"core_api/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/fatih/structs"
)

const BEARER string = "token"

type Authorization struct {
	errorCode int
	user      string
	role      string
	token     string
	apiKey    string
}

func accessibleRoles() map[string][]string {
	return map[string][]string{
		"/users": {"is_admin"},
	}
}

func NewAuthentication(c *http.Header) *Authorization {
	token := c.Get("Authorization")
	apiKey := c.Get("API_KEY")
	return &Authorization{
		apiKey: apiKey,
		token:  token,
	}
}

func (auth *Authorization) ApiKeyAuth(apiKey string) error {
	dbHandler := database.NewDBHandler()
	userObj := database.NewUser()
	user, err := dbHandler.GetRecordDatabaseByApiKey(userObj, auth.apiKey)
	if err != nil {
		auth.errorCode = http.StatusUnauthorized
		return errors.New("invalid API_KEY")
	}

	s := structs.New(user)

	var role string
	switch s.Field("IsAdmin").Value().(bool) {
	case true:
		role = "is_admin"
	case false:
		role = "operator"

	}
	auth.user = s.Field("Username").Value().(string)
	auth.role = role
	return nil
}

func (auth *Authorization) TokenAuth(token string) error {
	pieces := strings.SplitN(token, " ", 2)
	if len(pieces) < 2 {
		return error_code.ErrBearerToken
	}
	if pieces[0] != BEARER {
		auth.errorCode = 4
		return error_code.ErrBearerToken
	}
	crypt := utils.NewCryptoGraphic()
	tkObj := utils.NewTokenInfo()
	tkObj.CryptInterface = crypt

	if err := tkObj.CryptInterface.LoadRsaPrivatekey(); err != nil {
		auth.errorCode = 15
		return error_code.ErrInternal
	}
	tkObj.PrivKey = crypt.PrivateKey

	if err := tkObj.TokenDecrypt(pieces[1]); err != nil {
		auth.errorCode = 4
		return error_code.ErrInvalidToken
	}

	if err := tkObj.TokenValid(); err != nil {
		auth.errorCode = 4
		return error_code.ErrInvalidToken
	}
	auth.role = tkObj.Role
	auth.user = tkObj.User
	return nil
}

func (auth *Authorization) CheckPermission(uri string) bool {
	accessibleRoles, ok := accessibleRoles()[uri]
	if ok {
		for _, access := range accessibleRoles {
			if access == auth.role {
				return true
			}
		}
	}
	return false
}
