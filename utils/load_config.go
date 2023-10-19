package utils

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	DbHost           string `mapstructure:"API_DB_HOST"`
	DbUser           string `mapstructure:"API_DB_USER"`
	DbPassword       string `mapstructure:"API_DB_PASSWORD"`
	DbName           string `mapstructure:"API_DB_NAME"`
	DbPort           int    `mapstructure:"API_DB_PORT"`
	SecretKey        string `mapstructure:"API_SECRET_KEY"`
	BindHost         string `mapstructure:"API_BIND_HOST"`
	BindPort         int    `mapstructure:"API_BIND_PORT"`
	PathPrivatekey   string `mapstructure:"API_PATH_PRIVATEKEY"`
	ApiAdminUsername string `mapstructure:"API_ADMIN_USERNAME"`
	ApiAdminPassword string `mapstructure:"API_ADMIN_PASSWORD"`
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
	if c.BindPort == 0 {
		c.BindPort = 8080
	}
	if c.BindHost == "" {
		c.BindHost = "localhost"
	}
	if _, err := os.Stat(c.PathPrivatekey); errors.Is(err, os.ErrNotExist) {
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
