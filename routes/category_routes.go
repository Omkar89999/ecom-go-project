package routes

import (
	"ecom-go/controllers"
	"ecom-go/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {

	category := r.Group("/categories")

	category.Use(middleware.AuthMiddleware())
	{
		category.POST("/", controllers.CreateCategory)
		category.GET("/", controllers.GetCategories)
		category.GET("/:id", controllers.GetCategory)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}

}
