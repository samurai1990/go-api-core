package api

import (
	db "core_api/database"
	ser "core_api/serializers"
	"fmt"
	"net/http"

	"core_api/utils"

	"github.com/gin-gonic/gin"
)

type SigninBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	serObj := ser.NewGetSerializer(users)
	if e := serObj.GetListSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = serObj.SerData
		return serv, nil
	}
}

func Signin(c *gin.Context) (*server, error) {

	var account SigninBody
	serv := &server{
		ctx: c,
	}

	if err := c.ShouldBind(&account); err != nil {
		serv.ErrorCode = http.StatusBadRequest
		return serv, err
	}
	userObj := db.NewUser()
	dbHandler := db.NewDBHandler()
	user, err := dbHandler.GetDatabaseByName(userObj, account.Username)
	if err != nil {
		serv.ErrorCode = http.StatusInternalServerError
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
	token, err := hToken.GenerateToken(userObj.Username,userObj.IsAdmin)
	if err != nil {
		serv.ErrorCode = http.StatusInternalServerError
		return serv, fmt.Errorf("sorry, operation generate token failed")
	}
	apiKey, err := utils.GenerateApiKey(userObj.Username, userObj.Email)
	if err != nil {
		serv.ErrorCode = http.StatusInternalServerError
		return serv, fmt.Errorf("sorry, operation generate api key failed")
	}

	signStruct := struct {
		*db.User
		Token  string `json:"token"`
		ApiKey string `json:"api_key"`
	}{
		User:   userObj,
		Token:  token,
		ApiKey: apiKey,
	}
	resser := ser.NewUserSigninResponse()
	serObj := ser.NewGetSerializer(signStruct)
	if e := serObj.GetSigninSerializer(resser); e != nil {
		serv.ErrorCode = 500
		return serv, fmt.Errorf("internal server error,during error: %s ", e.Error())
	} else {
		serv.data = resser
		return serv, nil
	}
}
