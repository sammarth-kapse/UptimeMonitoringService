package main

import (
	"github.com/gin-gonic/gin"
)

// Function routes each request to its respective route-handler. All the route handlers are available in routeHandlers.go
func routes(router *gin.Engine) {

	router.POST("/urls", postURL)

	router.GET("/urls/:id", getUrlData)

	router.PATCH("/urls/:id", patchURL)

	router.DELETE("/urls/:id", deleteURL)

	router.POST("/urls/:id/activate", activateURL)

	router.POST("/urls/:id/deactivate", deactivateURL)
}
