package api

import (
	"awesomeProject/gin-demo/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/register", register)
	r.POST("/login", login)

	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
		UserRouter.POST("/reset_password", resetPassword)
		UserRouter.POST("/find_password", findPassword)
	}

	r.Run(":8080")
}
