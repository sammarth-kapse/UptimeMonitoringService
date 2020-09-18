package main

import (
	"UptimeMonitoringService/monitor"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
	}
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

func getUrlStatus(ctx *gin.Context) {

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
		displayMessage := "Invalid ID"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	}

}

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
		displayMessage := "Invalid ID"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	}
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

	if urlAddress, isPresent, isAlreadyActivated := monitor.ActivateURL(id); isAlreadyActivated {
		displayMessage := "URL is already Activated"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	} else if isPresent {
		displayMessage := "Activated URL: " + urlAddress
		ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte(displayMessage))
	} else {
		displayMessage := "Invalid ID"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	}

}

func deactivateURL(ctx *gin.Context) {
	id := ctx.Param("id")

	if urlAddress, isPresent, isAlreadyDeactivated := monitor.DeactivateURL(id); isAlreadyDeactivated {
		displayMessage := "URL: " + urlAddress + " is already Deactivated"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	} else if isPresent {
		displayMessage := "Deactivated URL: " + urlAddress
		ctx.Data(http.StatusOK, "application/json; charset=utf-8", []byte(displayMessage))
	} else {
		displayMessage := "Invalid ID"
		ctx.Data(http.StatusConflict, "application/json; charset=utf-8", []byte(displayMessage))
	}
}
