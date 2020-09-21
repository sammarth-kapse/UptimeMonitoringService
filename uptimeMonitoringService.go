package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()

	// Function in routes.go
	routes(router)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
