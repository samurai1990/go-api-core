package api

import (
	"github.com/gin-gonic/gin"
)

func PathRoutes(engine *gin.Engine) {
	engine.Use(BasicAuthPermission())
	engine.GET("/ping", wrapper(HandlePing))
	engine.POST("/users", wrapper(HandleCreateUsers))
	engine.GET("/users", wrapper(HandleGetUsers))
	engine.GET("/users/:id", wrapper(HandleRetrieveUsers))
	engine.PUT("/users/:id", wrapper(HandleUpdateUsers))
	engine.DELETE("/users/:id", wrapper(HandleDeleteUsers))
	engine.POST("/users/signin/", wrapper(Signin))
	engine.NoRoute(wrapper(HandleNoRoute))
}
