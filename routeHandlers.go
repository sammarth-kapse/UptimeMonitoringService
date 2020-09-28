package main

import (
	"UptimeMonitoringService/monitor"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// All the calls in this file to monitor package's functions are available in monitor/services.go

// Adds the request URL and corresponding data to the system.
func postURL(ctx *gin.Context) {
	var postRequest monitor.URLPostRequest

	err := ctx.BindJSON(&postRequest)
	if err != nil {
		log.Fatal(err)
		return
	}

	urlInfo, err := monitor.AddService(postRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id":                urlInfo.ID,
			"url":               urlInfo.URL,
			"crawl_timeout":     urlInfo.CrawlTimeout,
			"frequency":         urlInfo.Frequency,
			"failure_threshold": urlInfo.FailureThreshold,
			"status":            urlInfo.Status,
			"failure_count":     urlInfo.FailureCount,
		})
	}
}

func getURLData(ctx *gin.Context) {

	id := ctx.Param("id")

	if response, isPresent := monitor.GetURLDataByID(id); isPresent {
		ctx.JSON(http.StatusOK, gin.H{
			"id":                response.ID,
			"url":               response.URL,
			"crawl_timeout":     response.CrawlTimeout,
			"frequency":         response.Frequency,
			"failure_threshold": response.FailureThreshold,
			"status":            response.Status,
			"failure_count":     response.FailureCount,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}

}

// Updates the frequency/ status for a given request.
func patchURL(ctx *gin.Context) {

	var patchRequest monitor.URLPatchRequest

	err := ctx.BindJSON(&patchRequest)
	if err != nil {
		fmt.Println(err)
	}

	id := ctx.Param("id")
	if response, isPresent := monitor.UpdateURL(id, patchRequest); isPresent {
		ctx.JSON(http.StatusOK, gin.H{
			"id":                response.ID,
			"url":               response.URL,
			"crawl_timeout":     response.CrawlTimeout,
			"frequency":         response.Frequency,
			"failure_threshold": response.FailureThreshold,
			"status":            response.Status,
			"failure_count":     response.FailureCount,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}

// Removes the URL Data from the system and stops monitoring.
func deleteURL(ctx *gin.Context) {
	id := ctx.Param("id")

	isPresent := monitor.DeleteURLData(id)
	if isPresent {
		ctx.JSON(http.StatusNoContent, gin.H{
			"id":      id,
			"deleted": true,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}

func activateURL(ctx *gin.Context) {

	id := ctx.Param("id")

	if urlAddress, isPresent, isAlreadyActivated := monitor.ActivateURL(id); isAlreadyActivated {
		ctx.JSON(http.StatusConflict, gin.H{
			"id":      id,
			"url":     urlAddress,
			"message": "Already Activated",
		})
	} else if isPresent {
		ctx.JSON(http.StatusOK, gin.H{
			"id":      id,
			"url":     urlAddress,
			"message": "Activated",
		})

	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}

}

func deactivateURL(ctx *gin.Context) {
	id := ctx.Param("id")

	if urlAddress, isPresent, isAlreadyDeactivated := monitor.DeactivateURL(id); isAlreadyDeactivated {
		ctx.JSON(http.StatusConflict, gin.H{
			"id":      id,
			"url":     urlAddress,
			"message": "Already Deactivated",
		})
	} else if isPresent {
		ctx.JSON(http.StatusOK, gin.H{
			"id":      id,
			"url":     urlAddress,
			"message": "Deactivated",
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}
