package database

import (
	"core_api/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
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
	conf := utils.NewConfig()
	conf.LoadConfig(".")

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
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Print("Failed to connect database")
		return err
	}
	H.HDB = db
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	apiKey, err := utils.GenerateApiKey(u.Username, u.Password, u.Email)
	if err != nil {
		return err
	}
	u.ApiKey = apiKey
	return
}
