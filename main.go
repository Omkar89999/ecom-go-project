package main

import (
	"ecom-go/config"
	"ecom-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.UserProtectedRoutes(r)
	routes.CategoryRoutes(r)
	routes.ProductRoutes(r)

	routes.CartRoute(r)

	r.Run(":8080")

	// sqlDB, err := config.DB.DB()

	// if err != nil {
	// 	log.Fatalf("can not get sqlDB: %v", err)

	// }

	// if err := sqlDB.Ping(); err != nil {
	// 	log.Fatalf("ping failed: %v", err)
	// }

	// log.Println("DB ping successful. You're ready to continue.")
}
