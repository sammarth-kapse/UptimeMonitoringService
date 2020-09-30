package monitor

import (
	"UptimeMonitoringService/database"
	"UptimeMonitoringService/httpRequests"
	"fmt"
	"net/http"
	"time"
)

// All the structures and constants are present in utility.go

var repository database.RepositoryController
var httpCalls httpRequests.HttpController

func init() {
	database.SetRepoController(&database.MonitorRepo{})
	httpRequests.SetHTTPController(&httpRequests.MonitorHttp{})
	repository = database.GetRepoController()
	httpCalls = httpRequests.GetHTTPController()
}

// Monitors a URL till it's status is 'active'
func monitor(urlInfo *database.URLData) {
	isFirstCheck := true

	for isURLStatusActive(urlInfo) {
		if !isFirstCheck {
			ticker := time.NewTimer(time.Duration(urlInfo.Frequency) * time.Second)
			<-ticker.C
		}
		go checkURLUptime(urlInfo)
		isFirstCheck = false
	}
}

func stopMonitoring(urlInfo *database.URLData) {

	urlInfo.Status = INACTIVE
	err := repository.DatabaseSave(urlInfo)
	handleError(err)

	time.Sleep(time.Duration(urlInfo.Frequency+2) * time.Second)
}

func checkURLUptime(urlInfo *database.URLData) {

	if !isURLStatusActive(urlInfo) {
		return
	}

	resp, err := httpCalls.MakeHTTPGetRequest(urlInfo.CrawlTimeout, urlInfo.URL)

	if err != nil {
		increaseFailureCount(urlInfo) // Request didn't complete within crawl_timeout.
		return
	}

	if resp.StatusCode != http.StatusOK {
		increaseFailureCount(urlInfo) // Unexpected status-code in response.
	}

	fmt.Println("Checked Url : ", urlInfo.URL, " code: ", resp.StatusCode, "   time: ", time.Now())
}
