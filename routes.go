package main

import (
	"UptimeMonitoringService/Monitor"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func routes(router *gin.Engine) {

	router.POST("/urls", postURL)

	router.GET("/urls/:id", getUrlStatus)

	router.PATCH("/urls/:id", patchURL)

	router.DELETE("/urls/:id", deleteURL)

	router.POST("/urls/:id/activate", activateURL)

	router.POST("/urls/:id/deactivate", deactivateURL)
}

func postURL(ctx *gin.Context) {
	var postRequest Monitor.UrlPostRequest

	err := ctx.BindJSON(&postRequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":                "b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
		"url":               "abc.com",
		"crawl_timeout":     20,
		"frequency":         30,
		"failure_threshold": 50,
		"status":            "active",
		"failure_count":     0,
	})
}

func getUrlStatus(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id":                id,
		"url":               "abc.com",
		"crawl_timeout":     20,
		"frequency":         30,
		"failure_threshold": 50,
		"status":            "active",
		"failure_count":     0,
	})

}

func patchURL(ctx *gin.Context) {
	var patchRequest Monitor.UrlPatchRequest

	id := ctx.Param("id")

	err := ctx.BindJSON(&patchRequest)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":                id,
		"url":               "abc.com",
		"crawl_timeout":     20,
		"frequency":         patchRequest.Frequency,
		"failure_threshold": 50,
		"status":            "active",
		"failure_count":     0,
	})
}

func deleteURL(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id":      id,
		"deleted": true,
	})
}

func activateURL(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id":        id,
		"activated": true,
	})
}

func deactivateURL(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id":          id,
		"deactivated": true,
	})
}
