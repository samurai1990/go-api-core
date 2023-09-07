package database_test

import (
	"log"
	"strings"
	"testing"

	"core_api/database"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initialize() {
	viper.Set("API_DB_HOST", "192.168.10.21")
	viper.Set("API_DB_PORT", 5432)
	viper.Set("API_DB_USER", "user")
	viper.Set("API_DB_PASSWORD", "user")
	viper.Set("API_DB_NAME", "db")
}

func CreateTestDB(db *gorm.DB) *gorm.DB {
	hDB := db.Exec("CREATE DATABASE test_db;")
	if hDB.Error != nil {
		if strings.Contains(hDB.Error.Error(), "already exists") {
			log.Println("database \"test_db\" already exists")
		} else {
			log.Fatal("failed to create database")
		}
	}
	return hDB
}

func DropTestDB(db *gorm.DB) {
	if result := db.Exec("DROP DATABASE IF EXISTS test_db;"); result.Error != nil {
		log.Fatal("failed to drop database")
	}
}

func TestUserCRUD(t *testing.T) {
	initialize()
	h := database.NewDBHandler()
	assert.Equal(t, h.DBConnection(), nil)
	assert.NotEqual(t, h.HDB, nil)

	// create test db
	DropTestDB(h.HDB)
	db := CreateTestDB(h.HDB)

	// migrate User
	if err := db.AutoMigrate(&database.User{}); err != nil {
		log.Fatal("migration failed")
	}

	// remove test db
	DropTestDB(db)
}
