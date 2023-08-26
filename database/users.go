package database

import (
	"core_api/utils"
	"errors"
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/google/uuid"
)

func (user *User) ListRecord() (*[]User, error) {
	id1, _ := uuid.NewRandom()
	id2, _ := uuid.NewRandom()
	users := []User{
		{
			BaseModel: &BaseModel{
				ID: id1,
			},
			Username: "samurai",
			Password: "samurai",
			Email:    "samurai@com.com",
			ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
			IsAdmin:  true,
		},
		{
			BaseModel: &BaseModel{
				ID: id2,
			},
			Username: "bob",
			Password: "bob123456",
			Email:    "bob@com.com",
			ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
			IsAdmin:  false,
		},
	}
	return &users, nil
}

func (user *User) GetRecordByID(id uuid.UUID) (*User, error) {
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: "samurai",
		Email:    "samurai@com.com",
		Password: "samurai",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	return user, nil
}

func (user *User) GetRecordByName(name string) (*User, error) {
	id, _ := uuid.NewRandom()
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: "samurai",
		Password: "858a39db1542daacc92d9bb4fb8b563d35e13833cfc35e4ec2106c2043aa4bfc5a4373350267278dbcb8ee34c214898dfe27ce286b615b74bcb23642a9a067b5",
		Email:    "samurai@com.com",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	return user, nil
}

func (user *User) GetRecordByApiKey(apiKey string) (*User, error) {
	id, _ := uuid.NewRandom()
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: "samurai",
		Password: "858a39db1542daacc92d9bb4fb8b563d35e13833cfc35e4ec2106c2043aa4bfc5a4373350267278dbcb8ee34c214898dfe27ce286b615b74bcb23642a9a067b5",
		Email:    "samurai@com.com",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	return user, nil
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
	if err:=dbHandle.DBConnection();err!=nil{
		return nil , err
	}
	result:=dbHandle.HDB.Create(&user)
	return user, result.Error
}

func (user *User) UpdateRecord(model any) (*User, error) {
	id, _ := uuid.NewRandom()
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: "alic",
		Password: "alic12345",
		Email:    "alice@co.com",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	sw := utils.NewSchemaData(user)
	if err := sw.SchemaSwap(model); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) DeleteRecord(id string) error {
	return errors.New("not found")
}

func (user *User) CreateSuperUser() error {
	adminUser := struct {
		Username string
		Password string
		Email    string
		IsAdmin  bool
	}{
		Username: "admin",
		Password: "admin",
		Email:    "admin@api.co.com",
		IsAdmin:  true,
	}
	_, err := user.CreateRecord(adminUser)
	if err != nil {
		return fmt.Errorf("don`t create Super User with during error: %s", err.Error())
	} else {
		log.Println("super user created!!!")
		return nil
	}
}
