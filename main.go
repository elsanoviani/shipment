package main

import (
	"log"

	"shipment/database"
	"shipment/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, menggunakan environment variable sistem")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	r := gin.Default()
	routes.UserRoutes(r)

	r.Run(":8080")
}
