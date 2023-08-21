package utils

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type TokenInfo struct {
	CryptInterface CryptoGraphicInterface
	Token          string
	PrivKey        *rsa.PrivateKey
	User           string
	Role           string
}

func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

func (t *TokenInfo) GenerateToken(user string, isAdmin bool) (string, error) {
	token_lifespan, err := strconv.Atoi("24")
	if err != nil {
		return "", err
	}

	var role string
	switch isAdmin {
	case true:
		role = "is_admin"
	case false:
		role = "operator"
	}

	token, err := jwt.NewBuilder().
		Claim("username", user).
		Claim("role", role).
		Expiration(time.Now().Add(time.Hour * time.Duration(token_lifespan))).Build()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %s", err)
	}

	rsaCont := NewCryptoGraphic()
	if err := rsaCont.LoadRsaPrivatekey(); err != nil {
		return "", err
	}

	tokenSigned, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, rsaCont.PrivateKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %s", err)
	}
	t.PrivKey = rsaCont.PrivateKey
	tokenEnc, err := t.TokenEncrypt(tokenSigned)
	if err != nil {
		return "", err
	}

	return tokenEnc, nil
}

func (t *TokenInfo) TokenEncrypt(payload []byte) (string, error) {

	encrypted, err := jwe.Encrypt(payload, jwe.WithKey(jwa.RSA1_5, &t.PrivKey.PublicKey), jwe.WithContentEncryption(jwa.A128CBC_HS256))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt payload: %s", err)
	}

	return string(encrypted), nil
}

func (t *TokenInfo) TokenDecrypt(tokenEncrypted string) error {

	decrypted, err := jwe.Decrypt([]byte(tokenEncrypted), jwe.WithKey(jwa.RSA1_5, t.PrivKey))
	if err != nil {
		return fmt.Errorf("failed to decrypt: %s", err)
	}
	t.Token = string(decrypted)
	return nil
}

func (t *TokenInfo) TokenValid() error {
	{
		verifiedToken, err := jwt.Parse([]byte(t.Token), jwt.WithKey(jwa.RS256, t.PrivKey.PublicKey))
		if err != nil {
			return errors.New("token not valid")
		}
		claim := verifiedToken.PrivateClaims()
		t.User = claim["username"].(string)
		t.Role = claim["role"].(string)
	}
	return nil
}
