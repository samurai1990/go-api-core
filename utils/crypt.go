package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type CryptoGraphic struct {
	PrivateKey *rsa.PrivateKey
}

type CryptoGraphicInterface interface {
	GenerateRsaPrivateKey() error
	ExportRsaPrivateKeyAsPemStr() error
	LoadRsaPrivatekey() error
	DoPasswordsMatch(hashedPassword, requestPassword string) bool
	HashPassword(password string) (string, error)
}

func NewCryptoGraphic() *CryptoGraphic {
	return &CryptoGraphic{}
}

func (cr *CryptoGraphic) GenerateRsaPrivateKey() error {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate private key: %s", err)
	}
	cr.PrivateKey = privkey
	return nil
}

func (cr *CryptoGraphic) ExportRsaPrivateKeyAsPemStr() error {
	priKey := x509.MarshalPKCS1PrivateKey(cr.PrivateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY API",
		Bytes: priKey,
	}
	pathPrivateKey := viper.GetString("API_PATH_PRIVATEKEY")
	privatePem, err := os.Create(pathPrivateKey)
	if err != nil {
		return fmt.Errorf("error when create private.pem: %s", err)
	}
	defer privatePem.Close()

	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		return fmt.Errorf("error when encode private pem: %s", err)
	}
	return nil
}

func (cr *CryptoGraphic) LoadRsaPrivatekey() error {
	pathPrivateKey := viper.GetString("API_PATH_PRIVATEKEY")
	file, err := os.ReadFile(pathPrivateKey)
	if err != nil {
		return fmt.Errorf("do not readed Private key file")
	}

	spkiBlock, _ := pem.Decode(file)
	parsePrivate, err := x509.ParsePKCS1PrivateKey(spkiBlock.Bytes)
	if err != nil {
		return fmt.Errorf("do not parsed Private key")
	}

	cr.PrivateKey = parsePrivate
	return nil
}

func (cr *CryptoGraphic) DoPasswordsMatch(hashedPassword, requestPassword string) bool {
	currPasswordHash, _ := cr.HashPassword(requestPassword)

	return currPasswordHash == hashedPassword
}

func (cr *CryptoGraphic) HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	salt := viper.GetString("API_SECRET_KEY")
	passwordBytes = append(passwordBytes, []byte(salt)...)

	if _, err := sha512Hasher.Write(passwordBytes); err != nil {
		return "", err
	}

	hashedPasswordBytes := sha512Hasher.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex, nil
}
