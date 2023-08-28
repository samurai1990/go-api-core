package database

import (
	"core_api/utils"
	"errors"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type User struct {
	*BaseModel
	Username string `gorm:"size:255;not null;unique" binding:"required" json:"username"`
	Password string `gorm:"not null;" binding:"required" json:"password"`
	Email    string `gorm:"size:255;not null;unique" binding:"required" json:"email"`
	ApiKey   string `gorm:"not null;" json:"api_key"`
	IsAdmin  bool   `gorm:"default:false" json:"is_admin"`
}

func NewUser() *User {
	return &User{}
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	crypt := utils.NewCryptoGraphic()
	if hashedPassword, err := crypt.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPassword
	}
	apiKey, err := utils.GenerateApiKey(user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	user.ApiKey = apiKey
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		dest := tx.Statement.Dest
		newPassword := dest.(*User).Password
		crypt := utils.NewCryptoGraphic()
		if crypt.DoPasswordsMatch(user.Password, newPassword) {
			return errors.New("current password in same account is repeated from ex password")
		} else {
			hashedPassword, err := crypt.HashPassword(newPassword)
			if err != nil {
				return err
			} else {
				tx.Statement.SetColumn("Password", hashedPassword)
			}
		}
	}
	return nil
}

func (user *User) BeforeDelete(tx *gorm.DB) error {
	adminUser := viper.GetString("API_ADMIN_USERNAME")
	if user.Username == adminUser {
		return errors.New("persmission deny")
	}
	return nil
}
