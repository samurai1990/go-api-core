package serializers

import (
	"core_api/utils"
	"time"

	"github.com/google/uuid"
)

type UserCreateResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserGetResponse struct {
	*UserCreateResponse
}

type UserRetrieveResponse struct {
	*UserCreateResponse
}

type UserUpdateResponse struct {
	*UserCreateResponse
}

type UserSigninResponse struct {
	*UserCreateResponse
	ApiKey string `json:"api_key"`
	Token  string `json:"token"`
}

func NewUserGetResponse() *UserGetResponse {
	return &UserGetResponse{}
}

func NewUserRetrieveResponse() *UserRetrieveResponse {
	return &UserRetrieveResponse{}
}

func NewUserSigninResponse() *UserSigninResponse {
	return &UserSigninResponse{}
}

func NewUserCreateResponse() *UserCreateResponse {
	return &UserCreateResponse{}
}

func NewUserUpdateResponse() *UserUpdateResponse {
	return &UserUpdateResponse{}
}

func (ugr *UserGetResponse) ListSerializerResponse(model any) (any, error) {
	var users *[]UserGetResponse
	sw := utils.NewSchemaData(&users)
	if err := sw.SchemaSwap(model); err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (urr *UserRetrieveResponse) RetrieveSerializerResponse(model any) error {
	sw := utils.NewSchemaData(urr)
	if err := sw.SchemaSwap(model); err != nil {
		return err
	} else {
		return nil
	}
}

func (usr *UserSigninResponse) SigninSerializerResponse(model any) error {
	sw := utils.NewSchemaData(usr)
	if err := sw.SchemaSwap(model); err != nil {
		return err
	} else {
		return nil
	}
}

func (ucr *UserCreateResponse) CreateSerializerResponse(model any) error {
	sw := utils.NewSchemaData(ucr)
	if err := sw.SchemaSwap(model); err != nil {
		return err
	} else {
		return nil
	}
}

func (upr *UserUpdateResponse) UpdateSerializerResponse(model any) error {
	sw := utils.NewSchemaData(upr)
	if err := sw.SchemaSwap(model); err != nil {
		return err
	} else {
		return nil
	}
}
