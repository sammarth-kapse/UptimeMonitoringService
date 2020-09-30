package main

import (
	"UptimeMonitoringService/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	database.InitializeDatabase()

	router := gin.Default()

	// Function in routes.go
	routes(router)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
