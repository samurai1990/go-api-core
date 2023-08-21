package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.Path

		if !IsUrlValid(uri) {
			c.Next()
			return
		}

		if IsAnyPermission(uri) {
			c.Next()
			return
		}
		res := NewApiResponse()
		auth := NewAuthentication(&c.Request.Header)

		if auth.apiKey != "" {
			if err := auth.ApiKeyAuth(auth.apiKey); err != nil {
				res.StatusCode = auth.errorCode
				res.Error = true
				res.ErrorMessage = err.Error()
				c.AbortWithStatusJSON(res.StatusCode, res)
				return
			}
		} else if auth.token != "" {
			if err := auth.TokenAuth(auth.token); err != nil {
				res.StatusCode = auth.errorCode
				res.Error = true
				res.ErrorMessage = err.Error()
				c.AbortWithStatusJSON(res.StatusCode, res)
				return
			}
		} else {
			res.StatusCode = auth.errorCode
			res.Error = true
			res.ErrorMessage = "missing authorization header"
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		if auth.CheckPermission(uri) {
			c.Next()
			return
		} else {
			res.StatusCode = http.StatusForbidden
			res.Error = true
			res.ErrorMessage = "Access Denied,You don't have permission to access"
			c.AbortWithStatusJSON(res.StatusCode, res)
		}
	}
}

func IsUrlValid(uri string) bool {
	_, ok := accessibleRoles()[uri]
	if ok {
		return true
	} else {
		return false
	}
}

func IsAnyPermission(uri string) bool {
	for _, role := range accessibleRoles()[uri] {
		if role == "any" {
			return true
		}
	}
	return false
}
