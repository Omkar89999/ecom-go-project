package routes

import (
	"ecom-go/controllers"
	"ecom-go/middleware"

	"github.com/gin-gonic/gin"
)

func CartRoute(r *gin.Engine) {

	cart := r.Group("/cart")

	cart.Use(middleware.AuthMiddleware())
	{
		cart.POST("/add", controllers.AddToCart)
		cart.GET("/", controllers.ViewCart)
		cart.DELETE("/:id", controllers.RemoveFromCart)
	}

}
