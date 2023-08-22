package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	listenAddr string
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
			res.StatusCode = server.ErrorCode
			res.Error = true
			res.ErrorMessage = err.Error()
			c.JSON(res.StatusCode, res)
			return
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
		listenAddr: fmt.Sprintf("%s:%d", listenHost, listenPort),
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

func (s *server) Start() error {
	router := gin.New()
	router.Use(gin.Logger())
	PathRoutes(router)
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})
	if err := router.Run(s.listenAddr); err != nil {
		return err
	} else {
		return nil
	}
}
