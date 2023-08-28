package api

import (
	"github.com/gin-gonic/gin"
)

func PathRoutes(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	engine.GET("/health", wrapper(HandleHealth))
	engine.POST("/users/signin/", wrapper(Signin))
	engine.NoRoute(wrapper(HandleNoRoute))

	users := engine.Group("/")
	users.Use(BasicAuthPermission())
	{
		users.POST("/users", wrapper(HandleCreateUsers))
		users.GET("/users", wrapper(HandleGetUsers))
		users.GET("/users/:id", wrapper(HandleRetrieveUsers))
		users.PUT("/users/:id", wrapper(HandleUpdateUsers))
		users.DELETE("/users/:id", wrapper(HandleDeleteUsers))
	}
}
