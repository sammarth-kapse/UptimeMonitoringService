package main

import (
	"UptimeMonitoringService/database"
	"UptimeMonitoringService/monitor"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	database.InitializeDatabase()
	err := database.DB.AutoMigrate(&monitor.URLData{}) // Makes the table of structure URLData
	if err != nil {
		os.Exit(1)
	}

	router := gin.Default()

	// Function in routes.go
	routes(router)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
