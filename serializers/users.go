package serializers

import (
	"encoding/json"
	"time"
)

type UserGetResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	IsAdmin   bool      `json:"is_admin"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UserSigninResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
	ApiKey   string `json:"api_key"`
	Token    string `json:"token"`
}

func NewUserGetResponse() *UserGetResponse {
	return &UserGetResponse{}
}

func NewUserSigninResponse() *UserSigninResponse {
	return &UserSigninResponse{}
}

func (ugr *UserGetResponse) ListSerializerResponse(model any) (any, error) {
	var users *[]UserGetResponse
	userJson, _ := json.Marshal(model)
	err := json.Unmarshal(userJson, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (usr *UserSigninResponse) SigninSerializerResponse(model any) error {
	userJson, _ := json.Marshal(model)
	err := json.Unmarshal(userJson, &usr)
	if err != nil {
		return err
	}
	return nil
}
