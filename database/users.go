package database

import (
	errorCode "core_api/errors"
	"core_api/utils"
	"errors"
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm/clause"
)

func (user *User) ListRecord() (*[]User, error) {
	users := []User{}
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}
	results := dbHandle.HDB.Limit(10).Offset(10).Find(&users)
	return &users, results.Error
}

func (user *User) GetRecordByID(id uuid.UUID) (*User, error) {
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}
	results := dbHandle.HDB.Find(&user, "id = ?", id)
	return user, results.Error
}

func (user *User) GetRecordByName(name string) (*User, error) {
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}
	result := dbHandle.HDB.First(&user, "username = ?", name)
	return user, result.Error
}

func (user *User) GetRecordByApiKey(apiKey string) (*User, error) {
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}
	result := dbHandle.HDB.First(&user, "api_key = ?", apiKey)
	return user, result.Error
}

func (user *User) CreateRecord(model any) (*User, error) {
	s := structs.New(model)
	username := s.Field("Username").Value().(string)
	password := s.Field("Password").Value().(string)
	email := s.Field("Email").Value().(string)

	id, _ := uuid.NewRandom()
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: username,
		Password: password,
		Email:    email,
		IsAdmin:  s.Field("IsAdmin").Value().(bool),
	}
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}

	if result := dbHandle.HDB.Create(&user); result.Error != nil {
		errStruct := structs.New(result.Error)
		if errStruct.Field("Code").Value().(string) == "23505" {
			return nil, fmt.Errorf("%w : %s is exist", errorCode.ErrDuplicateKey, user.Username)
		}
	}
	return user, nil
}

func (user *User) UpdateRecord(model any) (*User, error) {
	params := NewUser()
	sw := utils.NewSchemaData(&params)
	if err := sw.SchemaSwap(model); err != nil {
		return nil, err
	}

	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return nil, err
	}
	if query := dbHandle.HDB.First(&user, "id = ?", params.ID); query.Error != nil {
		return user, query.Error
	}

	result := dbHandle.HDB.Model(&user).Clauses(clause.Returning{}).Where("id = ?", params.ID).Updates(params)
	return user, result.Error
}

func (user *User) DeleteRecord(id string) error {
	dbHandle := NewDBHandler()
	if err := dbHandle.DBConnection(); err != nil {
		return err
	}
	if query := dbHandle.HDB.First(&user, "id = ?", id); query.Error != nil {
		return errors.New("not found")
	}
	result := dbHandle.HDB.Where("id = ?", id).Delete(user)
	if result.RowsAffected != 1 {
		return errors.New("delete record failed")
	}
	return result.Error
}

func (user *User) CreateSuperUser() error {
	adminUser := viper.GetString("API_ADMIN_USERNAME")
	adminPassword := viper.GetString("API_ADMIN_PASSWORD")
	superUser := struct {
		Username string
		Password string
		Email    string
		IsAdmin  bool
	}{
		Username: adminUser,
		Password: adminPassword,
		Email:    fmt.Sprintf("%s@api.co", adminUser),
		IsAdmin:  true,
	}
	_, err := user.CreateRecord(superUser)
	if err != nil {
		if errors.Is(err, errorCode.ErrDuplicateKey) {
			return fmt.Errorf("superuser: `%s` is exist, during error:%w", superUser.Username, errorCode.ErrDuplicateKey)
		} else {
			return fmt.Errorf("don`t create Super User with during error: %s", err.Error())
		}
	} else {
		log.Println("super user created!!!")
		return nil
	}
}
