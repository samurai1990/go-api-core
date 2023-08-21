package api

import "github.com/gin-gonic/gin"

func HandlePing(c *gin.Context) (*server, error) {
	data := map[string]any{
		"message": "pong",
	}
	s := &server{
		ctx:  c,
		data: data,
	}
	return s, nil
}
