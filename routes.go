package main

import (
	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine) {

	router.POST("/urls", postURL)

	router.GET("/urls/:id", getUrlStatus)

	router.PATCH("/urls/:id", patchURL)

	router.DELETE("/urls/:id", deleteURL)

	router.POST("/urls/:id/activate", activateURL)

	router.POST("/urls/:id/deactivate", deactivateURL)
}
