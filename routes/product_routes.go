package routes

import (
	"ecom-go/controllers"
	"ecom-go/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {

	product := r.Group("/products")
	product.Use(middleware.AuthMiddleware())

	{
		product.POST("/", controllers.CreateProduct)
		product.GET("/", controllers.GetProducts)
		product.GET("/:id", controllers.GetProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}
}
