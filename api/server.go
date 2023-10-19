package api

import (
	error_code "core_api/errors"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type apiResponse struct {
	*WrappedError
	Data any `json:"data"`
}
type WrappedError struct {
	Error        bool   `json:"error" default:"false"`
	StatusCode   int    `json:"error_code" default:"0"`
	ErrorMessage string `json:"error_message" default:""`
}

type server struct {
	ListenAddr string
	ctx        *gin.Context
	data       any
	ErrorCode  int `default:"0"`
}

func NewApiResponse() *apiResponse {
	return &apiResponse{
		Data:         nil,
		WrappedError: &WrappedError{},
	}
}
func NewWrappedError(err bool, num int, message string) *WrappedError {
	return &WrappedError{
		Error:        err,
		StatusCode:   num,
		ErrorMessage: message,
	}
}

func wrapper(f func(c *gin.Context) (*server, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := NewApiResponse()
		server, err := f(c)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				res.Error = true
				res.StatusCode = 3
				res.ErrorMessage = err.Error()
				c.JSON(http.StatusNotFound, res)
				return
			case errors.Is(err, error_code.ErrUserNotFound):
				res.Error = true
				res.StatusCode = 4
				res.ErrorMessage = err.Error()
				c.JSON(http.StatusNotFound, res)
				return
			case errors.Is(err, error_code.ErrCredentials):
				res.Error = true
				res.StatusCode = 6
				res.ErrorMessage = err.Error()
				c.JSON(http.StatusUnauthorized, res)
				return
			default:
				res.Error = true
				res.StatusCode = server.ErrorCode
				res.ErrorMessage = err.Error()
				c.JSON(res.StatusCode, res)
				return
			}
		} else {
			res.Data = server.data
			c.JSON(server.ErrorCode, res)
		}
	}
}

func NewApiResponseByArgs(err bool, statusCode int, errorMessage string) *apiResponse {
	return &apiResponse{
		Data: nil,
		WrappedError: &WrappedError{
			Error:        err,
			StatusCode:   statusCode,
			ErrorMessage: errorMessage,
		},
	}
}

func NewServer(listenHost string, listenPort int) *server {
	return &server{
		ListenAddr: fmt.Sprintf("%s:%d", listenHost, listenPort),
	}
}

func NewBaseServer() *server {
	return &server{}
}

func (s *server) HandleErrorResponse(err int, errorMessage error) {
	res := NewApiResponseByArgs(true, err, errorMessage.Error())
	s.ctx.JSON(res.StatusCode, res)
}

func HandleNoRoute(c *gin.Context) (*server, error) {
	uri := c.Request.URL.Path
	return &server{
		ctx:       c,
		ErrorCode: 404,
	}, fmt.Errorf("page %s is not found", uri)
}

func (s *server) Setup() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	PathRoutes(router)
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})
	return router
}
