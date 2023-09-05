package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"math/rand"

	exc "core_api/errors"
	"github.com/spf13/viper"
)

func GenerateApiKey(texts ...string) (string, error) {
	scretkey := viper.GetString("API_SECRET_KEY")
	block, err := aes.NewCipher([]byte(scretkey))
	if err != nil {
		return "", fmt.Errorf("aes cipher error: %w", exc.ErrInternal)
	}
	var XorText string
	for _, text := range texts {
		for i := 0; i < len(text); i++ {
			XorText += string(text[i] ^ text[rand.Intn(len(text))])
		}
	}
	ciphertext := make([]byte, aes.BlockSize+len(string(XorText)))
	iv := ciphertext[:aes.BlockSize]

	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(XorText))
	cfb.XORKeyStream(cipherText, []byte(XorText))

	sha_512 := sha512.New()
	sha_512.Write(cipherText)
	return base64.StdEncoding.EncodeToString(sha_512.Sum(nil)), nil
}
