package utils

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	DB_HOST         string `mapstructure:"API_DB_HOST"`
	DB_USER         string `mapstructure:"API_DB_USER"`
	DB_PASSWORD     string `mapstructure:"API_DB_PASSWORD"`
	DB_NAME         string `mapstructure:"API_DB_NAME"`
	DB_PORT         int    `mapstructure:"API_DB_PORT"`
	SECRET_KEY      string `mapstructure:"API_SECRET_KEY"`
	BIND_HOST       string `mapstructure:"API_BIND_HOST"`
	BIND_PORT       int    `mapstructure:"API_BIND_PORT"`
	PATH_PRIVATEKEY string `mapstructure:"API_PATH_PRIVATEKEY"`
}

func NewConfig() *config {
	return &config{}
}

func (c *config) LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	if c.BIND_PORT == 0 {
		c.BIND_PORT = 8080
	}
	if c.BIND_HOST == "" {
		c.BIND_HOST = "localhost"
	}
	if _, err := os.Stat(c.PATH_PRIVATEKEY); errors.Is(err, os.ErrNotExist) {
		rsaKey := NewCryptoGraphic()
		if err := rsaKey.GenerateRsaPrivateKey(); err != nil {
			log.Fatal(err)
		}
		if err := rsaKey.ExportRsaPrivateKeyAsPemStr(); err != nil {
			log.Fatal(err)
		}
	} else {

		return nil
	}

	return nil
}
