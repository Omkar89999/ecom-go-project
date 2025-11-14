package routes

import (
	"ecom-go/controllers"
	"ecom-go/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}

func UserProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", controllers.Me)
	}
}
