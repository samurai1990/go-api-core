package api

import (
	error_code "core_api/errors"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		uriSplit := strings.Split(c.Request.URL.Path, "/")
		uri := "/" + uriSplit[1]

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
				switch {
				case errors.Is(err, error_code.ErrInternal):
					res.StatusCode = auth.errorCode
					res.Error = true
					res.ErrorMessage = err.Error()
					c.AbortWithStatusJSON(http.StatusInternalServerError, res)
					return
				default:
					res.StatusCode = auth.errorCode
					res.Error = true
					res.ErrorMessage = err.Error()
					c.AbortWithStatusJSON(http.StatusForbidden, res)
					return
				}
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

func IsAnyPermission(uri string) bool {
	for _, role := range accessibleRoles()[uri] {
		if role == "any" {
			return true
		}
	}
	return false
}
