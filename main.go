package main

import (
	"ecom-go/config"
	"log"
)

func main() {
	config.ConnectDatabase()

	sqlDB, err := config.DB.DB()

	if err != nil {
		log.Fatalf("can not get sqlDB: %v", err)

	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	log.Println("DB ping successful. You're ready to continue.")
}
