package api

import (
	"github.com/gin-gonic/gin"
)

func PathRoutes(engine *gin.Engine) {
	engine.Use(BasicAuthPermission())
	engine.GET("/ping", wrapper(HandlePing))
	engine.GET("/users", wrapper(HandleGetUsers))
	engine.POST("/users/signin/", wrapper(Signin))
	engine.NoRoute(wrapper(HandleNoRoute))
}
