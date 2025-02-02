package main

import (
	"log"
	"mock-ses-api/config"
	"mock-ses-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.InitDB()

	routes.SetupRoutes(r, db)

	log.Println("Mock AWS SES API running on :8080")
	r.Run(":8080")
}