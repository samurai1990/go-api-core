package api

import (
	db "core_api/database"
	ser "core_api/serializers"
	"errors"
	"fmt"
	"net/http"

	"core_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserSigninRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateRequest struct {
	*UserSigninRequest
	IsAdmin bool   `json:"is_admin"`
	Email   string `json:"email" binding:"required" `
}

type UserUpdateRequest struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
	IsAdmin  bool      `json:"is_admin"`
	Email    string    `json:"email"`
}

type UserRetrieveRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

func HandleGetUsers(c *gin.Context) (*server, error) {

	userObj := db.NewUser()
	resser := ser.NewUserGetResponse()
	dbHandler := db.NewDBHandler()
	users, err := dbHandler.ListDatabase(userObj)

	serv := &server{
		ctx: c,
	}

	if err != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", err.Error())
	}

	serObj := ser.NewSerializer(users)
	if e := serObj.GetListSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = serObj.SerData
		return serv, nil
	}
}

func Signin(c *gin.Context) (*server, error) {

	var account *UserSigninRequest
	serv := &server{
		ctx: c,
	}

	if err := c.ShouldBind(&account); err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, err
	}
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	user, err := dbHandler.GetRecordDatabaseByName(userObj, account.Username)
	if err != nil {
		serv.ErrorCode = http.StatusUnauthorized
		return serv, err
	}
	userObj = user.(*db.User)
	crypt := utils.NewCryptoGraphic()

	ok := crypt.DoPasswordsMatch(userObj.Password, account.Password)
	if !ok {
		serv.ErrorCode = http.StatusUnauthorized
		return serv, fmt.Errorf("sorry, username / password is not correct")
	}

	hToken := utils.NewTokenInfo()
	token, err := hToken.GenerateToken(userObj.Username, userObj.IsAdmin)
	if err != nil {
		serv.ErrorCode = http.StatusInternalServerError
		return serv, fmt.Errorf("sorry, operation generate token failed")
	}

	signStruct := struct {
		*db.User
		Token string `json:"token"`
	}{
		User:  userObj,
		Token: token,
	}
	resser := ser.NewUserSigninResponse()
	serObj := ser.NewSerializer(signStruct)
	if e := serObj.GetSigninSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = resser
		return serv, nil
	}
}

func HandleCreateUsers(c *gin.Context) (*server, error) {
	serv := &server{
		ctx: c,
	}
	var account *UserCreateRequest
	if err := c.ShouldBind(&account); err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, err
	}
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	dbHandler.Model = account
	user, err := dbHandler.CreateRecordToDatabase(userObj)
	if err != nil {
		serv.ErrorCode = http.StatusInternalServerError
		return serv, err
	}
	resser := ser.NewUserCreateResponse()
	serObj := ser.NewSerializer(user)
	if e := serObj.GetCreateSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = resser
		serv.ErrorCode = http.StatusCreated
		return serv, nil
	}
}

func HandleRetrieveUsers(c *gin.Context) (*server, error) {
	serv := &server{
		ctx: c,
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, errors.New("uuid invalid")
	}
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	user, err := dbHandler.GetRecordDatabaseByID(userObj, id)
	if err != nil {
		serv.ErrorCode = http.StatusNotFound
		return serv, err
	}
	resser := ser.NewUserRetrieveResponse()
	serObj := ser.NewSerializer(user)
	if e := serObj.GetRetrieveSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = resser
		serv.ErrorCode = http.StatusOK
		return serv, nil
	}
}

func HandleUpdateUsers(c *gin.Context) (*server, error) {
	var account *UserUpdateRequest
	serv := &server{
		ctx: c,
	}

	if err := c.ShouldBind(&account); err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, err
	}

	if id, err := uuid.Parse(c.Param("id")); err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, errors.New("uuid invalid")
	} else {
		account.ID = id
	}
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	dbHandler.Model = account
	user, err := dbHandler.UpdateRecordToDatabase(userObj)
	if err != nil {
		serv.ErrorCode = http.StatusNotFound
		return serv, err
	}
	resser := ser.NewUserUpdateResponse()
	serObj := ser.NewSerializer(user)
	if e := serObj.GetUpdateSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = resser
		serv.ErrorCode = http.StatusAccepted
		return serv, nil
	}
}

func HandleDeleteUsers(c *gin.Context) (*server, error) {
	serv := &server{
		ctx: c,
	}
	id := c.Param("id")
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	err := dbHandler.DeleteRecordFromDatabase(userObj, id)
	if err != nil {
		serv.ErrorCode = http.StatusNotFound
		return serv, err
	}
	serv.ErrorCode = http.StatusNoContent
	return serv, nil
}
