package monitor

import (
	"UptimeMonitoringService/database"
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"net/http"
	"os"
	"time"
)

// All the structures and constants are present in utility.go

func init() {
	err := database.DB.AutoMigrate(&URLData{})
	if err != nil {
		os.Exit(1)
	}
	setRepoController(&monitorRepo{})
	setHTTPController(&monitorHttp{})
}

type HttpController interface {
	makeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error)
}

type monitorHttp struct{}

var httpCalls HttpController

func setHTTPController(hType HttpController) {
	httpCalls = hType
}

// Monitors a URL till it's status is 'active'
func monitor(urlInfo *URLData) {

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

func stopMonitoring(urlInfo *URLData) {

	urlInfo.Status = INACTIVE
	err := repository.databaseSave(urlInfo)
	handleError(err)

	time.Sleep(time.Duration(urlInfo.Frequency+2) * time.Second)
}

func checkURLUptime(urlInfo *URLData) {

	if !isURLStatusActive(urlInfo) {
		return
	}

	resp, err := httpCalls.makeHTTPGetRequest(urlInfo.CrawlTimeout, urlInfo.URL)

	if err != nil {
		increaseFailureCount(urlInfo) // Request didn't complete within crawl_timeout.
		return
	}

	if resp.StatusCode != http.StatusOK {
		increaseFailureCount(urlInfo) // Unexpected status-code in response.
	}

	fmt.Println("Checked Url : ", urlInfo.URL, " code: ", resp.StatusCode, "   time: ", time.Now())
}

func (mh *monitorHttp) makeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error) {

	timeout := time.Duration(crawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	return client.Get(url, nil)
}
