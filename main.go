package main

import (
	"accounts/auth"
	"accounts/models"
	"accounts/middlewares"
	
	"github.com/gin-gonic/gin"

)

func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/accounts")

	public.POST("/register", auth.Register)
	public.POST("/login",auth.Login)


	protected := router.Group("/accounts/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user",auth.CurrentUser)


	router.Run(":8080")

}
