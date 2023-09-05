package database

import (
	"fmt"
	"log"
	"os"
	"time"

	exc "core_api/errors"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBHandler struct {
	HDB   *gorm.DB
	Model any
}

func NewDBHandler() *DBHandler {
	return &DBHandler{}
}

func (H *DBHandler) DBConnection() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		viper.GetString("API_DB_HOST"),
		viper.GetInt("API_DB_PORT"),
		viper.GetString("API_DB_USER"),
		viper.GetString("API_DB_PASSWORD"),
		viper.GetString("API_DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Print("Failed to connect database")
		return fmt.Errorf("%w: failed to connect database", exc.ErrDB)
	}
	H.HDB = db
	if err := db.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("%w: migration failed", exc.ErrDB)
	}
	return nil
}
