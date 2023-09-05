package database_test

import (
	"core_api/database"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func initialize() {
	viper.Set("API_DB_HOST", "192.168.10.21")
	viper.Set("API_DB_PORT", 5432)
	viper.Set("API_DB_USER", "user")
	viper.Set("API_DB_PASSWORD", "user")
	viper.Set("API_DB_NAME", "db")
}

func TestRsaFundamental(t *testing.T) {
	initialize()
	h := database.NewDBHandler()
	assert.Equal(t, h.DBConnection(), nil)
	assert.NotEqual(t, h.HDB, nil)
	
}
